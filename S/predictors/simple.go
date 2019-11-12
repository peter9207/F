package predictors

import (
	"github.com/peter9207/F/S/average"
)

type Predictor interface {
	Predict(data []float64) bool
}

type SimplePredictor struct {
	Score      int64
	windowSize int64
}

func SimpleRolling(bucketSize int64) (p *SimplePredictor) {
	return &SimplePredictor{windowSize: bucketSize}
}

func (p *SimplePredictor) Predict(data []float64) (b bool) {

	averages := average.Rolling(data, p.windowSize)

	b = averages[0] < averages[len(averages)-1]
	return
}
