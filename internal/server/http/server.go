package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AmirMirzayi/clean-go-project/internal/config"
)

type HttpServer struct {
	server http.Server
	mux    *http.ServeMux
}

func NewServer(cfg config.HttpServer) *HttpServer {
	return &HttpServer{
		server: http.Server{
			Addr: fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		},
		mux: http.NewServeMux(),
	}
}

func (s *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func (s *HttpServer) Start() {
	log.Println("http server started on address: ", s.server.Addr)
	s.server.Handler = s.mux
	s.server.ListenAndServe()
}
