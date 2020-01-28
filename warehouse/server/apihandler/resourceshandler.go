package apihandler

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/dutyfree/persistence"
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
	//TODO: write errors when exist
}
