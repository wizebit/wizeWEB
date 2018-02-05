package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
	"net/http"
	"log"
)

type AuthClaims struct {
	ExpiredAt int
	CustomerRole string
}

//RSA KEYS AND INITIALISATION

const (
	privKeyPath = "keys/rs256-4096-private.rsa"
	pubKeyPath  = "keys/rs256-4096-public.pem"
)

var VerifyKey, SignKey []byte

func initKeys() {
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}
//
////AUTH TOKEN CREATE
//func CreateTokenMiddleware(w http.ResponseWriter, r *http.Request, role string) {
//	initKeys()
//
//	//create a rsa 256 token
//	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
//		"exp":          time.Now().Add(time.Minute * 20).Unix(),
//		"customerRole": role,
//	})
//
//	//set claims
//	tokenString, err := token.SignedString(SignKey)
//
//	if err != nil {
//		fmt.Println(SignKey)
//		RespondWithError(w, http.StatusInternalServerError, "Error signing token: "+err.Error())
//		return
//	}
//
//	//create a token instance using the token string
//	RespondWithJSON(w, http.StatusOK, map[string]string{"token": tokenString})
//}
//
////AUTH TOKEN VALIDATION
//
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//initKeys()

	//token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
	//	func(token *jwt.Token) (interface{}, error) {
	//		return VerifyKey, nil
	//	},
	//)

	ParseTokenFromSignedTokenString(r.Header.Get("Authorization"))

	//fmt.Println(err, token.Claims.(jwt.MapClaims)["exp"])

	//if err == nil {
	//if token.Valid {
	//	next(w, r)
	//} else {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	fmt.Fprint(w, "Token is not valid")
	//}

	//} else {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	fmt.Fprint(w, "Unauthorised access to this resource")
	//}

	//ParseTokenFromSignedTokenString(token.Raw)
}


//AUTH TOKEN CREATE

func CreateSignedTokenString(role string) (string, error) {
	privateKey, err := ioutil.ReadFile("keys/rs256-4096-private.rsa")
	if err != nil {
		return "", fmt.Errorf("error reading private key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key: %v\n", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":          time.Now().Add(time.Minute * 20).Unix(),
		"customerRole": role,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v\n", err)
	}

	return tokenString, nil
}

//AUTH TOKEN VALIDATION

func ParseTokenFromSignedTokenString(tokenString string) (*jwt.Token, error) {
	publicKey, err := ioutil.ReadFile("keys/rs256-4096-public.pem")
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
	//token:=parsedToken.Claims["exp"]



	//fmt.Println(parsedToken.Claims)
	fmt.Println((parsedToken.Claims.(jwt.MapClaims)["exp"]).(string)+"~"+ (parsedToken.Claims.(jwt.MapClaims)["customerRole"]).(string))



	return parsedToken, nil
}