package matches

import "github.com/facup94/worldcup-predictor/internal/model"

type Service struct {
	Repo repository
}

type repository interface {
}

func (s Service) GetByID(id string) (model.Match, error) {
	return model.Match{}, nil
}
