package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createSignedJwt(key string, customClaims *map[string]interface{}) (string, error) {

	// Create jwt
	token := jwt.New(jwt.SigningMethod)

	now := time.Now()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	// claims["iss"] = issuerID
	// claims["aud"] = audience
	// claims["jti"] = tokenId
	// claims["sub"] = subject/userId
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Minute * time.Duration(app.AccessTokenExpires)).Unix()

	if custom != nil {
		for k, claim := range *custom {
			claims[k] = claim
		}
	}

	// Sign the token string
	return token.SignedString(key)
}

func verifyJwt(key string, token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, &claims, func(_token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	return claims, err
}
