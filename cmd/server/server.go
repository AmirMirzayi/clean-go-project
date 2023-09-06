package main

import (
	"github.com/AmirMirzayi/clean-go-project/internal/config"
	"github.com/AmirMirzayi/clean-go-project/internal/server/http"
	"github.com/AmirMirzayi/clean-go-project/pkg/api"
)

func main() {
	cfg := config.LoadConfig()
	server := http.NewServer(cfg.HttpServer)
	api.InitRouting(server)
	server.Start()
}
