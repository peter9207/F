package main

type RollingWindowCrossingEvent struct {
	Date   int64   `json:"timestamp"`
	Value  float64 `json:"value"`
	Window int     `json:"size"`
	Meta   string  `json:"meta"`
}
