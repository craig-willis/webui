package main

import (
	"log"
	"net/http"

	"earthcube.org/Project418/webui/search"

	"github.com/gorilla/mux"
)

// MyServer struct for mux router
type MyServer struct {
	r *mux.Router
}

func main() {
	searchroute := mux.NewRouter()
	searchroute.HandleFunc("/", search.HoldingPage) // temporary handler for this URL
	// searchroute.HandleFunc("/", search.DoSearch)  // the REAL handler for this URL
	http.Handle("/", searchroute)

	imageRouter := mux.NewRouter()
	imageRouter.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/images/", &MyServer{imageRouter})

	cssRouter := mux.NewRouter()
	cssRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/css/", &MyServer{cssRouter})

	htmlRouter := mux.NewRouter()
	htmlRouter.PathPrefix("/html/").Handler(http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/html/", &MyServer{htmlRouter})

	log.Printf("About to listen on 9900. Go to http://127.0.0.1:9900/")

	err := http.ListenAndServe(":9900", nil)
	// http 2.0 http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Let the Gorilla work
	s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
