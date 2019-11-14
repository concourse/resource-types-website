package publichandler

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

type Handler struct {
	Path string
}

func (s Handler) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	box := packr.NewBox(s.Path)
	CacheNearlyForever(http.StripPrefix("/public/", http.FileServer(box))).ServeHTTP(r, w)
}
