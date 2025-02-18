package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/garcios/asset-trak-portfolio/graphql/generated"
	"github.com/garcios/asset-trak-portfolio/graphql/middlewares"
	"github.com/garcios/asset-trak-portfolio/graphql/resolvers"
	"github.com/gin-gonic/gin"

	"log"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize Keycloak OIDC Provider
	//middlewares.InitKeycloak()

	r := gin.Default()

	r.Use(middlewares.Services())

	// TODO: properly implement protected and unprotected queries
	//protected := r.Group("/protected")
	//protected.Use(middlewares.AuthMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 GraphQL Gateway running on :8080")

}

func graphqlHandler() gin.HandlerFunc {
	h := handler.New(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &resolvers.Resolver{},
		}))

	// Server setup:
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
