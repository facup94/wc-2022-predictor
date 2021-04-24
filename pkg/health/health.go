package health

import (
	"io"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "pong")
}
