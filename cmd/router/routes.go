package router

import (
	"github.com/facup94/worldcup-predictor/internal/matches"
	"github.com/facup94/worldcup-predictor/internal/players"
	"github.com/facup94/worldcup-predictor/internal/predictions"
	"github.com/facup94/worldcup-predictor/internal/tournaments"
	"github.com/facup94/worldcup-predictor/pkg/health"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(
	r chi.Router,
	//session sessionHandler,
	tournament tournaments.Handler,
	player players.Handler,
	match matches.Handler,
	prediction predictions.Handler,
) {
	// session
	//r.Post("/signup", session.SignUp)
	//r.Post("/login", session.LogIn)

	// players
	r.Get("/players/{id}", player.Get)

	// tournaments
	r.Get("/tournaments/{id}", tournament.GetByID)
	r.Get("/tournaments/{id}/scores", tournament.GetByID)
	r.Get("/players/{playerID}/tournaments", tournament.GetByPlayerID)

	// matches
	r.Get("/matches/{id}", match.Get)

	// predictions
	r.Get("/predictions/{id}/", prediction.GetByID)
	r.Post("/predictions", prediction.Create)
	r.Put("/predictions", prediction.Update)
	r.Get("/tournaments/{tournamentID}/players/{playerID}/predictions", prediction.GetByTournamentPlayer)
	r.Get("/tournaments/{tournamentID}/matches/{matchID}/predictions", prediction.GetByTournamentMatch)

	// extra
	r.Get("/ping", health.Ping)
}
