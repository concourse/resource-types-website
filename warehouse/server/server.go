package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/concourse/dutyfree/fetcher"
	"github.com/concourse/dutyfree/persistence"
)

type Server struct {
	Port                     int
	Exited                   chan bool
	PublicFilesFetcher       fetcher.Fetcher
	ResourceTypesFileFetcher fetcher.Fetcher
	Kill                     chan bool
	srv                      *http.Server
}

func (s *Server) Close() error {
	return s.srv.Shutdown(context.TODO())
}
func (s *Server) Start() {

	indexHndlr, err := NewIndexHandler(s.PublicFilesFetcher)
	if err != nil {
		//TODO: don't panic
		panic("server error: " + err.Error())
	}

	fs := &persistence.Filesystem{
		Fetcher: s.ResourceTypesFileFetcher,
	}
	err = fs.LoadResources()
	if err != nil {
		//TODO: don't panic
		panic(err)
	}
	warehouseMux := http.NewServeMux()
	warehouseMux.Handle("/api/v1/", NewApiHandler(fs))

	warehouseMux.Handle("/public/", NewPublicHandler(s.PublicFilesFetcher))
	warehouseMux.Handle("/", indexHndlr)

	go func() {
		s.srv = &http.Server{
			Handler:      warehouseMux,
			Addr:         net.JoinHostPort("localhost", strconv.Itoa(s.Port)),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("ERROR: ", err)
		}
		s.Exited <- true
	}()
}
