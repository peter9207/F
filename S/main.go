package main

import (
	"bufio"
	"encoding/csv"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strconv"
)

func parseFloat(s string) (f float64) {
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

type Stock struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   int64
}

var rootCmd = &cobra.Command{
	Use:   "S",
	Short: "my feeble attempt to figure out what stocks to buy",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			cmd.Help()
			return
		}

		csvFile := args[0]

		f, _ := os.Open(csvFile)

		data := []Stock{}

		reader := csv.NewReader(bufio.NewReader(f))
		for {
			line, err := reader.Read()
			if err == io.EOF {
				return
			} else if err != nil {
				log.Fatal(err)
				return
			}

			data = append(data, Stock{
				Date:     line[0],
				Open:     parseFloat(line[1]),
				High:     parseFloat(line[2]),
				Low:      parseFloat(line[3]),
				Close:    parseFloat(line[4]),
				AdjClose: parseFloat(line[5]),
				Volume:   parseInt(line[6]),
			})

			log.Printf("row read: %v", line)
		}

	},
}

func main() {
	rootCmd.Execute()
}
