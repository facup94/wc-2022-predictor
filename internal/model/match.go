package model

import "time"

type Match struct {
	TeamA      Team
	TeamB      Team
	Time       time.Time
	TeamAScore int
	TeamBScore int
}
