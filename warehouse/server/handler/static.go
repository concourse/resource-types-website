package handler

import "net/http"

type staticHandler struct {
	Path string
}

func (s staticHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}
