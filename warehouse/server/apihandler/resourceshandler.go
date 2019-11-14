package apihandler

import (
	"encoding/json"
	"github.com/concourse/dutyfree/persistence"
	"net/http"
)

type resourcesHandler struct {
	p persistence.Persistence
}

func (rh resourcesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resources := rh.p.GetAllResources()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resources)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
