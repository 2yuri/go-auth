package services

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTServiceInterface interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTService struct {
	secretKey string
	issure    string
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: "secret",
		issure:    "ecocentauro-sistemas",
	}
}

type returnClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *JWTService) GenerateToken(name string, mail string) string {
	claims := &returnClaim{
		name,
		mail,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Print(token)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Print(err)
		return ""
	}

	return t
}

func (Service *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid Token: %v", encodedToken)
		}
		return []byte(Service.secretKey), nil
	})
}
