package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type AuthClaims struct {
	ExpiredAt    int
	CustomerRole string
}

//RSA KEYS AND INITIALISATION

const (
	privKey = "keys/rs256-4096-private.rsa"
	pubKey  = "keys/rs256-4096-public.pem"
)

//AUTH TOKEN CREATE

func CreateSignedTokenString(id int) (string, int, error) {
	//read private key
	privateKey, err := ioutil.ReadFile(privKey)
	if err != nil {
		return "", 0, fmt.Errorf("error reading private key file: %v\n", err)
	}
	//parse token
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", 0, fmt.Errorf("error parsing RSA private key: %v\n", err)
	}
	//token claims
	expiresIn := 1200

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":          time.Now().Add(time.Second * 1200).Unix(),
		"customerId": id,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", 0, fmt.Errorf("error signing token: %v\n", err)
	}

	return tokenString, expiresIn, nil
}

//AUTH TOKEN VALIDATION

func ParseTokenFromSignedTokenString(tokenString string) (*jwt.Token, error) {
	publicKey, err := ioutil.ReadFile(pubKey)
	if err != nil {
		return nil, fmt.Errorf("error reading public key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA public key: %v\n", err)
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	return parsedToken, nil
}
