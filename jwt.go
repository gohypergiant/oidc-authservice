package main

import (
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func mapClaimsToInterface(mapClaims *jwt.MapClaims) *map[string]interface{} {

	claims := map[string]interface{}{}

	for k, v := range *mapClaims {
		claims[k] = v
	}

	return &claims
}

func hasNonInteractiveScope(mapClaims *jwt.MapClaims) bool {

	claims := *mapClaims

	scopesClaim := claims["scopes"]

	if scopesClaim != nil {

		scopes := scopesClaim.([]string)

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

func (j *jwtExchange) sign(externalClaims *map[string]interface{}, scopes *[]string) (string, *map[string]interface{}) {

	// Create jwt
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	if externalClaims != nil {
		for k, claim := range *externalClaims {
			claims[k] = claim
		}
	}

	if scopes != nil {
		claims["scopes"] = *scopes
	}

	// NOTE: these should be set by the caller, retrieved from the valid oidc token
	// claims["iss"] = issuerID
	// claims["aud"] = audience
	// claims["jti"] = tokenId
	// claims["sub"] = userId

	now := time.Now()
	claims["iat"] = now.Unix()

	// Generate an appropriate expiry for this token
	if hasNonInteractiveScope(&claims) {
		claims["exp"] = now.AddDate(5, 0, 0).Unix() // 5 years for non interactive (sdk) tokens
	} else {
		claims["exp"] = now.AddDate(0, 0, 1).Unix() // 1 day for standard tokens
	}

	signed, err := token.SignedString([]byte(j.oauth2Config.ClientSecret))
	if err != nil {
		log.Println(err.Error())
		return "", nil
	}

	// Sign the token string
	return signed, mapClaimsToInterface(&claims)
}

func (j *jwtExchange) verify(token string) (*map[string]interface{}, error) {

	claims := new(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(token, claims, func(_token *jwt.Token) (interface{}, error) {
		return []byte(j.oauth2Config.ClientSecret), nil
	})

	return mapClaimsToInterface(claims), err
}
