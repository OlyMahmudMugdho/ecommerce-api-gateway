package server

import (
	"log"
	"net/http"
)

type Server struct {
	Port   string
	Router *http.ServeMux
}

func NewServer(port string) *Server {
	return &Server{
		Port:   ":" + port,
		Router: http.NewServeMux(),
	}
}

func (s *Server) Run() {
	log.Printf("server is running on port %v \n", s.Port)
	err := http.ListenAndServe(s.Port, s.Router)

	if err != nil {
		log.Fatal("server crashed")
	}
}
