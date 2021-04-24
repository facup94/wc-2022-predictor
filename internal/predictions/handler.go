package predictions

import "net/http"

type Handler struct {
	Service
}

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetByTournamentPlayer(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetByTournamentMatch(w http.ResponseWriter, r *http.Request) {

}
