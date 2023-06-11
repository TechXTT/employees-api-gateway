package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/TechXTT/employees-api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) *AuthMiddlewareConfig {
	return &AuthMiddlewareConfig{
		svc: svc,
	}
}

// pass in the mux router
type contextKey string

const userIDKey contextKey = "UserID"

func (c *AuthMiddlewareConfig) AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.Split(auth, "Bearer ")[1]

		res, err := c.svc.Client.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: token})
		if err != nil || res.Status != http.StatusOK {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//set UserID in context
		ctx := context.WithValue(r.Context(), userIDKey, res.UserId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
