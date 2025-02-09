package main

import (
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/config"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/generated"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/repository"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("cant load env: %s", err)
	}
	port := cfg.Port

	// TODO: repository selection
	storage, err := repository.LoadStorage(true, "")
	if err != nil {
		log.Fatal("cant load storage")
	}

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
