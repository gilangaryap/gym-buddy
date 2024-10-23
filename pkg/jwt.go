package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	Id    		string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Role  		string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWT(id, phone_number string, role string) *claims {
	return &claims{
		Id:    id,
		PhoneNumber: phone_number,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "TICKITZ",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
}

func (c *claims) GenerateToken() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

func VerifyToken(token string) (*claims, error) {
	secret := os.Getenv("JWT_SECRET")
	data, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claimData := data.Claims.(*claims)
	return claimData, nil
}