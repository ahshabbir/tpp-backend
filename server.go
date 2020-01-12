package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type server struct {
	port         string
	workdir      string
	logWriter    io.Writer
	chiefHandler *chi.Mux
}

func newServer(p string, lw io.Writer) *server {
	return &server{
		port:         p,
		logWriter:    lw,
		chiefHandler: chi.NewRouter(),
	}
}

func (s *server) loadRoutes() {
	s.chiefHandler.HandleFunc(
		"/helloworld",
		handleString("yay this works :)"),
	)
}

func handleString(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}
}

func (s *server) start() error {
	s.loadRoutes()
	return http.ListenAndServe(":"+s.port, s.chiefHandler)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	s := newServer(port, os.Stdout)
	log.Fatal(s.start())
}
