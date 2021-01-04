package main

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

type RollingWindowCrossingEvent struct {
	Date   int64   `json:"timestamp"`
	Value  float64 `json:"value"`
	Window int     `json:"size"`
	Meta   string  `json:"meta"`
}

func produceEvent(e RollingWindowCrossingEvent) (err error) {

	return
}

func connectES() (client *elasticsearch.Client, err error) {
	client, err = elasticsearch.NewDefaultClient()
	return

}
