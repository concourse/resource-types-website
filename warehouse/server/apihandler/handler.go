package apihandler

import (
	"github.com/concourse/dutyfree/persistence"
	"net/http"
)

type Handler struct {
	p persistence.Persistence
}

func NewApiHandler(p persistence.Persistence) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/resources", resourcesHandler{p: p})
	return mux
}
