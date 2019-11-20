package predictors

type DecoratedPredictor struct {
	predictors []*Predictor
}

func (p *DecoratedPredictor) Predict(data []float64) bool {

	return false
}
