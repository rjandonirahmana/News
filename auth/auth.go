package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type auth struct {
	secret      []byte
	secretAdmin []byte
}

type Auth interface {
	CreateToken(userID *string) (*string, error)
	ValidateToken(encodedToken *string) (*string, error)
	TokenAdmin(id *string) (*string, error)
	ValidateAdmin(encodedToken *string) (*string, error)
}

func NewAuth(secret string, secretAdmin string) *auth {
	return &auth{secret: []byte(secret), secretAdmin: []byte(secretAdmin)}
}

func (a *auth) CreateToken(userID *string) (*string, error) {
	claim := jwt.MapClaims{
		"user_id": *userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"active":  true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err := token.SignedString(a.secret)
	if err != nil {
		return nil, err
	}

	return &tokenstring, nil
}

func (a *auth) ValidateToken(encodedToken *string) (*string, error) {
	token, err := jwt.Parse(*encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID ERROR")
		}

		return []byte(a.secret), nil

	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, errors.New("Unauthorized")
	}

	userID := claims["user_id"].(string)
	return &userID, nil
}

func (a *auth) TokenAdmin(id *string) (*string, error) {
	claims := jwt.MapClaims{
		"admin_id": *id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenstr, err := token.SignedString(a.secretAdmin)

	if err != nil {
		return nil, err
	}

	return &tokenstr, nil

}

func (a *auth) ValidateAdmin(encodedToken *string) (*string, error) {
	token, err := jwt.Parse(*encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID ERROR")
		}

		return []byte(a.secretAdmin), nil

	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, errors.New("Unauthorized")
	}

	userID := claims["admin_id"].(string)
	return &userID, nil
}
