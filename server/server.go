package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	s.Router.HandleFunc("GET /", s.ProxyAuth())
	log.Printf("server is running on port %v \n", s.Port)
	err := http.ListenAndServe(s.Port, s.Router)

	if err != nil {
		log.Fatal("server crashed")
	}
}

func (s *Server) ProxyAuth() http.HandlerFunc {
	fmt.Println("ok")
	authUrl, _ := url.Parse("http://localhost:8082/")
	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(authUrl)
			pr.Out.Host = pr.In.Host
		},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}
