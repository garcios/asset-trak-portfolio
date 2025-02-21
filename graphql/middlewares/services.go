package middlewares

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/graphql/services"
	"github.com/gin-gonic/gin"
)

var serviceKey = &contextKey{name: "services"}

type contextKey struct {
	name string
}

func Services() gin.HandlerFunc {
	return func(c *gin.Context) {
		svcs := services.Services{
			PortfolioService: services.NewPortfolioService(),
		}

		// Enhance the context with the services
		c.Request = c.Request.WithContext(
			context.WithValue(
				c.Request.Context(),
				serviceKey,
				&svcs))
		c.Next()
	}
}

// GetServices is used to retrieve the service directory from the context
func GetServices(ctx context.Context) *services.Services {
	return ctx.Value(serviceKey).(*services.Services)
}
