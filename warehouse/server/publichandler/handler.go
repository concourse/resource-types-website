package publichandler

import (
	"net/http"

	"github.com/concourse/dutyfree/fetcher"
)

type Handler struct {
	Fetcher fetcher.Fetcher
}

func (s Handler) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	CacheNearlyForever(http.StripPrefix("/public/", http.FileServer(s.Fetcher))).ServeHTTP(r, w)
}
