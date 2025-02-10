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
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/dataloaders"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/generated"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/resolvers"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/repository"
)

func main() {
	flags := config.ParseFlags()

	env, err := config.GetConfig(flags.InMemory)
	if err != nil {
		log.Fatal(err)
	}

	var connStr string
	if env.DB != nil {
		connStr = env.DB.ConnStr()
	}

	storage, err := repository.LoadStorage(flags.InMemory, connStr)
	if err != nil {
		log.Fatal(err)
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
	http.Handle("/query", dataloaders.Middleware(storage)(srv))

	log.Printf("🚀 Server ready at http://localhost:%s/", env.Port)
	log.Fatal(http.ListenAndServe(":"+env.Port, nil))
}
