package predictors

import (
	"github.com/peter9207/F/S/average"
)

type Predictor struct {
	Score int64
}

func Simple() (p *Predictor) {
	return &Predictor{}
}

func (p *Predictor) Predict(data []float64) (b bool) {

	averages := average.Rolling(data, 7)

	b = averages[0] < averages[len(averages)]
	return
}
