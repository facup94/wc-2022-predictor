package mysql

import "github.com/facup94/worldcup-predictor/internal/model"

func (repo Repository) GetTournamentByID(id string) (model.Tournament, error) {
	return model.Tournament{}, nil
}

func (repo Repository) GetTournamentsByPlayerID(id string) ([]model.Tournament, error) {
	return nil, nil
}
