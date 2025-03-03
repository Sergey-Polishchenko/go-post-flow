package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/userapp"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/pkg/logging"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/storage/redisrepo"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/generated"
	"github.com/Sergey-Polishchenko/go-post-flow/internal/transport/graph/resolvers"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{Addr: ":6379"})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis client error: %s", err)
	}

	_ = rdb.FlushDB(ctx).Err()

	repoFactory := redisrepo.New(rdb)
	userRepo := repoFactory.NewUserRepo()

	logger := logging.NewZapLogger()

	userApp := userapp.New(userRepo, logger)

	execSchema := generated.NewExecutableSchema(
		generated.Config{
			Resolvers: resolvers.NewResolver(userApp),
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

	log.Printf("🚀 Server ready at http://localhost:%d/", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
