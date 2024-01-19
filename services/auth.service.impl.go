package services

import (
	"ecommerce/gmr/interfaces"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	AuthDataLayer interfaces.AuthDataLayer
}

const jwtSecretKey = "your-testing-secret-key"

func NewAuthService(adl interfaces.AuthDataLayer) interfaces.AuthServiceLayer {
	return &AuthService{
		AuthDataLayer: adl,
	}
}

func (as *AuthService) Login(username string, password string) (string, error) {

	user, err := as.AuthDataLayer.VerifyCredentials(username, password)

	if err != nil {
		fmt.Println("sharudl2")
		return "", err
	}

	fmt.Println(user)
	token, err := as.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (as *AuthService) GenerateToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userName,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (as *AuthService) ValidateToken(token string) (string, error) {
	fmt.Print("sh")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", jwt.ErrSignatureInvalid
	}

	return userID, nil
}
