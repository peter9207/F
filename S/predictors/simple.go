package predictors

import (
	"github.com/peter9207/F/S/average"
)

type Predictor struct {
	Score      int64
	windowSize int64
}

func SimpleRolling(bucketSize int64) (p *Predictor) {
	return &Predictor{windowSize: bucketSize}
}

func (p *Predictor) Predict(data []float64) (b bool) {

	averages := average.Rolling(data, p.windowSize)

	b = averages[0] < averages[len(averages)-1]
	return
}
