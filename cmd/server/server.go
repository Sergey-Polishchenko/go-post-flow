package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/generated"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/repository"
)

const defaultPort = "8080"

func main() {
	// TODO: make custom config
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// TODO: repository selection
	storage := repository.LoadStorage(true)

	execSchema := generated.NewExecutableSchema(
		generated.Config{
			Resolvers: resolvers.NewResolver(storage),
		},
	)

	srv := handler.New(execSchema)

	srv.AddTransport(
		&transport.Websocket{
			Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			KeepAlivePingInterval: 10 * time.Second,
		},
	)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 Server ready at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
