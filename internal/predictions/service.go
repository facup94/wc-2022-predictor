package predictions

import "github.com/facup94/worldcup-predictor/internal/model"

type Service struct {
	Repo repository
}

type repository interface {
}

func (s Service) Get(id string) (model.Prediction, error) {
	return model.Prediction{}, nil
}
