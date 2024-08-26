package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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

	s.Router.HandleFunc("/", s.ProxyAuth)

	log.Printf("server is running on port %v \n", s.Port)

	err := http.ListenAndServe(s.Port, s.Router)

	if err != nil {
		log.Fatal("server crashed")
	}
}

func (s *Server) ProxyAuth(w http.ResponseWriter, r *http.Request) {
	var host string

	service, _ := strings.CutPrefix(r.URL.Path, "/")

	switch service {
	case "auth":
		host = "http://localhost:8082/"
	case "cart":
		host = "http://localhost:8083/"
	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error":   true,
			"message": "requested service is not available",
		})
		return
	}

	authUrl, _ := url.Parse(host)

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(authUrl)
			pr.Out.Host = pr.In.Host
		},
	}

	proxy.ServeHTTP(w, r)

	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	proxy.ServeHTTP(w, r)
	// })
}
