package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/OlyMahmudMugdho/ecommerce-api-gateway/configs"
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

	serviceName := strings.Split(r.URL.Path, "/")[1]
	fmt.Println(serviceName)
	sConfig := configs.NewServiceConfig()
	host := sConfig.GetHost(serviceName)
	fmt.Println(host)
	if host == "" {
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
			fmt.Println(pr.Out.URL)
		},
	}

	proxy.ServeHTTP(w, r)
}
