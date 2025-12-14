package middleware

import (
	"coba_dulu/internal/handler"
	"coba_dulu/pkg"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT(jwtService pkg.JWTService, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			handler.ErrorResponse(c, "No token found", http.StatusUnauthorized, "Request needs a token")
			return
		}
		
		if !strings.HasPrefix(authHeader, "Bearer ") {
			handler.ErrorResponse(c, "Invalid token format", http.StatusUnauthorized, "Token must be in Bearer format")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			log.Println(err)
			handler.ErrorResponse(c, "Invalid token", http.StatusUnauthorized, err.Error())
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])

			// Check role if required
			if requiredRole != "" {
				if claims["role"] != requiredRole {
					handler.ErrorResponse(c, "Forbidden", http.StatusForbidden, "You are not authorized to access this resource")
					return
				}
			}
			c.Next()
		} else {
			handler.ErrorResponse(c, "Invalid token", http.StatusUnauthorized, "Token is not valid")
		}
	}
}
