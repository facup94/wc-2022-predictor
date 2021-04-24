package model

type Prediction struct {
	Match      Match
	Player     Player
	Tournament Tournament
	TeamAScore int
	TeamBScore int
}
