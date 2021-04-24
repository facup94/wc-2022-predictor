package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/facup94/worldcup-predictor/cmd/router"
	"github.com/facup94/worldcup-predictor/internal/matches"
	"github.com/facup94/worldcup-predictor/internal/players"
	"github.com/facup94/worldcup-predictor/internal/predictions"
	"github.com/facup94/worldcup-predictor/internal/repository/mysql"
	"github.com/facup94/worldcup-predictor/internal/tournaments"
	"github.com/facup94/worldcup-predictor/pkg/dependencies"
	"github.com/facup94/worldcup-predictor/pkg/environment"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // seed the random generator. Comment this line for deterministic behaviour

	// (external) dependencies
	depManager, err := dependencies.GetDependencyManager(environment.GetEnvironment())
	if err != nil {
		log.Fatal("Error starting up server", err)
	}

	// Repositories
	db := depManager.DB()
	repository := mysql.Repository{DB: db}

	// services
	matchesService := matches.Service{Repo: repository}
	playersService := players.Service{Repo: repository}
	predictionsService := predictions.Service{Repo: repository}
	tournamentsService := tournaments.Service{Repo: repository}

	// Handlers
	matchesHandler := matches.Handler{Service: matchesService}
	playersHandler := players.Handler{Service: playersService}
	predictionsHandler := predictions.Handler{Service: predictionsService}
	tournamentsHandler := tournaments.Handler{Service: tournamentsService}

	// Routing
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	router.SetupRoutes(r, tournamentsHandler, playersHandler, matchesHandler, predictionsHandler)

	// Start server
	port := os.Getenv("HOST_PORT")
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
