package server

import (
	"github.com/concourse/dutyfree/persistence"
	"github.com/concourse/dutyfree/server/apihandler"
	"net/http"

	"github.com/concourse/dutyfree/server/indexhandler"
	"github.com/concourse/dutyfree/server/publichandler"
)

func NewPublicHandler(path string) http.Handler {
	return publichandler.Handler{Path: path}
}

func NewIndexHandler(path string) (http.Handler, error) {
	return indexhandler.NewHandler(path)
}

func NewApiHandler(p persistence.Persistence) http.Handler {
	return apihandler.NewApiHandler(p)
}
