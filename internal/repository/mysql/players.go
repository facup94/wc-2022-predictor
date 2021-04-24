package mysql

import "github.com/facup94/worldcup-predictor/internal/model"

func (repo Repository) GetPlayerByID(id string) (model.Player, error) {
	return model.Player{}, nil
}
