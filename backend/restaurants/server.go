package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required!")
	}

	broker := &Broker{
		config: config,
	}

	return broker, nil
}

func (n *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)
	log.Println("Starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Println("Error starting server:", err)
	} else {
		log.Fatalf("Server stopped!")
	}
}
