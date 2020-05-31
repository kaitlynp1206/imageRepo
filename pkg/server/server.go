package server

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

type Server struct {
	mux *mux.Router
}

func NewServer() *Server {
	return &Server{
		mux : mux.NewRouter(),
	}
}

func (s *Server) Start() {
	s.mux.HandleFunc("/home", HandleExample)
	http.Handle("/", s.mux)
	http.ListenAndServe(":8080", nil)
}

func HandleExample(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}