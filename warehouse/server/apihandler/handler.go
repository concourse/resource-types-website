package apihandler

import (
	"net/http"

	"github.com/concourse/dutyfree/persistence"
)

type Handler struct {
	p persistence.Persistence
}

func NewApiHandler(p persistence.Persistence) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/resources", resourcesHandler{p: p})
	return mux
}
