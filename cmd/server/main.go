package main

import (
	"github.com/angrodrigo10/rate-limiter/config"
	"github.com/angrodrigo10/rate-limiter/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	server.StartServer(cfg)
}
