package packages

import (
	"blogging-platform/backend-service/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey = []byte("secret-key")

type CustomJWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// func (claims *CustomJWTClaims) Validate() error {
// 	return claims.RegisteredClaims.Valid()
// }

func GenerateJWT(user model.User) (string, error) {
	// expiration time
	// expirationTime := time.Now().Add(24 * time.Hour)

	// normal flexible claim
	// claims token
	// claims := &jwt.MapClaims{
	// 	"Subject": map[string]interface{}{
	// 		"username": user.Id,
	// 	},
	// 	"ExpiresAt": jwt.NewNumericDate(expirationTime),
	// }

	// production mode
	claims := CustomJWTClaims{
		Username: user.Id.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	tokenString, err := token.SignedString(JwtSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
