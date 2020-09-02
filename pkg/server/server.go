package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaitlynp1206/imageRepo/pkg/image"
	"github.com/kaitlynp1206/imageRepo/pkg/user"
)

type Server struct {
	mux        *mux.Router
	imgManager *image.ImagesManager
	usrManager *user.UsersManager
}

func NewServer(ctx context.Context) *Server {
	db, _ := sql.Open(DriverName, DataSourceName)
	return &Server{
		mux:        mux.NewRouter().StrictSlash(true),
		imgManager: image.NewImagesManager(ctx, db),
		usrManager: user.NewUsersManager(db),
	}
}

func (s *Server) Start() {
	s.mux.HandleFunc("/", HomeHandler)
	s.mux.HandleFunc("/image", s.imgManager.ImageHandler)
	s.mux.HandleFunc("/user", s.usrManager.UserHandler)
	http.Handle("/", s.mux)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	s.mux.ServeHTTP(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page of Image Repo")
}
