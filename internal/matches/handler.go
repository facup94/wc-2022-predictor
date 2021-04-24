package matches

import "net/http"

type Handler struct {
	Service
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {

}
