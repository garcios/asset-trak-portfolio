package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var verifier *oidc.IDTokenVerifier

func InitKeycloak() {
	godotenv.Load()
	providerURL := os.Getenv("KEYCLOAK_ISSUER_URL") // e.g., http://localhost:8080/realms/graphql-app
	provider, err := oidc.NewProvider(context.Background(), providerURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Keycloak: %v", err))
	}

	verifier = provider.Verifier(&oidc.Config{ClientID: os.Getenv("KEYCLOAK_CLIENT_ID")})
}

// Keycloak JWT Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		tokenStr := strings.Split(authHeader, "Bearer ")[1]
		token, err := verifier.Verify(context.Background(), tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		var claims map[string]interface{}
		if err := token.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot parse claims"})
			return
		}

		// Attach user info to context
		c.Set("user", claims["preferred_username"])
		c.Set("roles", claims["realm_access"].(map[string]interface{})["roles"])

		c.Next()
	}
}
