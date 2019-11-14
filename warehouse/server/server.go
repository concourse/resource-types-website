package server

import (
	"encoding/json"
	"github.com/concourse/dutyfree/infoData"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type Server struct {
	Port int
	Exited chan bool
}

func (s *Server) Start() {

	r := mux.NewRouter()
	r.HandleFunc("/info", InfoHandler)

	s.Exited = make(chan bool)

	go func() {
		err := http.ListenAndServe(":9090", r)
		if err != nil {
			fmt.Println("ERROR: ", err)
			s.Exited <- true
		}
	}()
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data := infoData.Info{
		Data: "this is dutyfree",
	}
	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
