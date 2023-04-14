package main

import (
	"context"
	"os"
	"os/signal"
	"net/http"

	"github.com/fujiwara/ridge"
	"github.com/Papillon6814/xemar/service"
)

const (
	defaultPort = "8080"
)

func serve(ctx context.Context) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()
	ridge.RunWithContext(ctx, ":"+port, "/", mux)
}

func main() {
	ctx := context.Background()

	xenCli, err := NewXenClient()

	sctx, _ := signal.NotifyContext(ctx, os.Interrupt, os.Kill)

	serve(sctx)
}
