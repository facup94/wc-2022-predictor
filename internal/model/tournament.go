package model

type Tournament struct {
	Owner        Player
	Participants []Player
	Predictions  []Prediction
}
