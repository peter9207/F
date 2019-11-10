package main

import (
	"bufio"
	"encoding/csv"
	"github.com/peter9207/F/S/predictors"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strconv"
)

type Stock struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   int64
}

func readExport(filename string) (stocks []Stock, err error) {

	f, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(f))
	first := true
	for {
		var line []string
		line, err = reader.Read()
		if err == io.EOF {
			return stocks, nil
		} else if err != nil {
			log.Fatal(err)
			return
		}
		if first {
			first = false
			continue
		}

		stocks = append(stocks, Stock{
			Date:     line[0],
			Open:     parseFloat(line[1]),
			High:     parseFloat(line[2]),
			Low:      parseFloat(line[3]),
			Close:    parseFloat(line[4]),
			AdjClose: parseFloat(line[5]),
			Volume:   parseInt(line[6]),
		})
	}

}

func parseFloat(s string) (f float64) {
	// log.Printf("trying to parse %s", s)
	var err error
	if f, err = strconv.ParseFloat(s, 64); err != nil {
		panic(err)
	}
	return
}

func parseInt(s string) (i int64) {
	var err error
	if i, err = strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	}
	return
}

var simpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "a simple rolling average calculation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		csvFile := args[0]
		stocks, err := readExport(csvFile)

		if err != nil {
			log.Fatal(err)
			return
		}

		data := []float64{}
		for _, v := range stocks {
			data = append(data, v.Close)
		}

		p := predictors.Simple()
		result := p.Predict(data)

		log.Printf("result: %v", result)
		return
	},
}

var rootCmd = &cobra.Command{
	Use:   "S",
	Short: "an attempt to try various means of analysing stocks data",
}

func main() {

	rootCmd.AddCommand(simpleCmd)
	rootCmd.Execute()
}
