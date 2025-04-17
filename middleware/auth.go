package middleware

import (
	"blogging-platform/backend-service/common/helpers"
	"blogging-platform/backend-service/common/packages"
	"blogging-platform/backend-service/common/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || strings.HasPrefix("Bearer ", authHeader) {
			helpers.CustomErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &packages.CustomJWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return packages.JwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			helpers.CustomErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Set("username", structs.AuthUser{
			Username: claims.Username,
		})
		c.Next()
	}
}
