package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/garcios/asset-trak-portfolio/graphql/generated"
	"github.com/garcios/asset-trak-portfolio/graphql/resolvers"

	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &resolvers.Resolver{},
		}))

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Println("🚀 GraphQL Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
