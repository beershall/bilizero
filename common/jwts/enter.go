package jwts

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayLoad struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

func GenToken(user JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claim := CustomClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expires))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(accessSecret))
}
