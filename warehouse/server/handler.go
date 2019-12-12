package server

import (
	"net/http"

	"github.com/concourse/dutyfree/server/publichandler"
)

func NewPublicHandler(path string) http.Handler {
	return publichandler.Handler{Path: path}
}
