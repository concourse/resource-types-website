package server

import (
	"net/http"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/persistence"
	"github.com/concourse/dutyfree/server/apihandler"
	"github.com/concourse/dutyfree/server/indexhandler"
	"github.com/concourse/dutyfree/server/publichandler"
)

func NewPublicHandler(fetchr *fetcher.Fetcher) http.Handler {
	return publichandler.Handler{Fetcher: *fetchr}
}

func NewIndexHandler(fetchr fetcher.Fetcher) (http.Handler, error) {
	return indexhandler.NewHandler(fetchr)
}

func NewApiHandler(p persistence.Persistence) http.Handler {
	return apihandler.NewApiHandler(p)
}
