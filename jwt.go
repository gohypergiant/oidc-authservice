package main

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func hasNonInteractiveScope(claims jwt.MapClaims) {

	scopesClaim := claims["scopes"]

	if scopesClaim != nil {

		scopes := interfaceSliceToStringSlice(scopesClaim.([]interface{}))

		for _, scope := range scopes {

			if strings.HasPrefix(scope, "sdk") {
				return true
			}
		}
	}

	return false
}

type jwtExchange struct {
	oauth2Config *oauth2.Config
}

func (j *jwtExchange) sign(externalClaims *map[string]interface{}) (string, *map[string]interface{}) {

	// Create jwt
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	if externalClaims != nil {
		for k, claim := range *externalClaims {
			claims[k] = claim
		}
	}

	// NOTE: these should be set by the caller, retrieved from the valid oidc token
	// claims["iss"] = issuerID
	// claims["aud"] = audience
	// claims["jti"] = tokenId
	// claims["sub"] = userId

	now := time.Now()
	claims["iat"] = now.Unix()

	if hasNonInteractiveScope(claims) {
		claims["exp"] = now.AddDate(5, 0, 0).Unix() // 5 years for sdk tokens
	} else {
		claims["exp"] = now.AddDate(0, 0, 1).Unix() // 1 day for standard tokens
	}

	// Sign the token string
	return token.SignedString(j.oauth2Config.ClientSecret), &claims
}

func (j *jwtExchange) verify(token string) (*jwt.MapClaims, error) {

	claims := new(jwt.MapClaims)

	_, err = jwt.ParseWithClaims(token, claims, func(_token *jwt.Token) (interface{}, error) {
		return j.oauth2Config.ClientSecret, nil
	})

	return claims, err
}
