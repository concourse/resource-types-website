package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	Port   int
	Exited chan bool
}

func (s *Server) Start() {
	warehouseMux := http.NewServeMux()

	//warehouseMux.Handle("/api/v1", newApiHandler)
	warehouseMux.Handle("/", NewPublicHandler("../../../web/public/"))
	go func() {
		srv := &http.Server{
			Handler:      warehouseMux,
			Addr:         net.JoinHostPort("localhost", strconv.Itoa(s.Port)),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		fmt.Println("I am about to start")
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println("ERROR: ", err)
			s.Exited <- true
		}
	}()
}

//func (s *Server) Start() error{
//
//	r := mux.NewRouter()
//	r.HandleFunc("/info", InfoHandler)
//
//	staticHandler := staticServer{staticPath: "../web/public", indexPath: "index.html"}
//	r.PathPrefix("/").Handler(staticHandler)
//
//	s.Exited = make(chan bool)
//
//}
//
//// staticServer implements the http.Handler interface, so we can use it
//// to respond to HTTP requests. The path to the static directory and
//// path to the index file within that static directory are used to
//// serve the SPA in the given static directory.
//type staticServer struct {
//	staticPath string
//	indexPath  string
//}
//
//// ServeHTTP inspects the URL path to locate a file within the static dir
//// on the SPA handler. If a file is found, it will be served. If not, the
//// file located at the index path on the SPA handler will be served. This
//// is suitable behavior for serving an SPA (single page application).
//func (h staticServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	// get the absolute path to prevent directory traversal
//	path, err := filepath.Abs(r.URL.Path)
//	if err != nil {
//		// if we failed to get the absolute path respond with a 400 bad request
//		// and stop
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// prepend the path with the path to the static directory
//	path = filepath.Join(h.staticPath, path)
//
//	// check whether a file exists at the given path
//	_, err = os.Stat(path)
//	if os.IsNotExist(err) {
//		// file does not exist, serve index.html
//		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
//		return
//	} else if err != nil {
//		// if we got an error (that wasn't that the file doesn't exist) stating the
//		// file, return a 500 internal server error and stop
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// otherwise, use http.FileServer to serve the static dir
//	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
//}
//
//func InfoHandler(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusOK)
//	data := infoData.Info{
//		Data: "this is dutyfree",
//	}
//	response, err := json.Marshal(data)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	_, err = w.Write(response)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//}
