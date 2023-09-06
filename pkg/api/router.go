package api

import (
	httpserver "github.com/AmirMirzayi/clean-go-project/internal/server/http"

	"net/http"
)

func InitRouting(server *httpserver.HttpServer) {
	helloWorld(server)
}

func helloWorld(server *httpserver.HttpServer) {
	server.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test"))
	})
}
