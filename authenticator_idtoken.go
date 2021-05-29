package main

import (
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

type idTokenAuthenticator struct {
	header       string // header name where id token is stored
	caBundle     []byte
	provider     *oidc.Provider
	oauth2Config *oauth2.Config
	userIDClaim  string // retrieve the userid if the claim exists
	groupsClaim  string
}

func (s *idTokenAuthenticator) AuthenticateRequest(r *http.Request) (*authenticator.Response, bool, error) {
	logger := loggerForRequest(r)

	// get id-token from header
	bearer := getBearerToken(r.Header.Get(s.header))
	if len(bearer) == 0 {
		return nil, false, nil
	}

	ctx := setTLSContext(r.Context(), s.caBundle)

	var claims map[string]interface{}

	// Check first for a valid exchanged id token
	exchange := &jwtExchange{oauth2Config: s.oauth2Config}
	exchangeClaims, err := exchange.verify(bearer)

	if err == nil {
		for k, v := range *exchangeClaims {
			claims[k] = v
		}
	}

	// Verifying received ID token
	if err != nil {
		verifier := s.provider.Verifier(&oidc.Config{ClientID: s.oauth2Config.ClientID})
		token, err := verifier.Verify(ctx, bearer)
		if err != nil {
			logger.Errorf("id-token verification failed: %v", err)
			return nil, false, nil
		}

		if err = token.Claims(&claims); err != nil {
			logger.Errorf("retrieving user claims failed: %v", err)
			return nil, false, nil
		}
	}

	if claims[s.userIDClaim] == nil {
		// No USERID_CLAIM, pass this authenticator
		logger.Error("USERID_CLAIM doesn't exist in the id token")
		return nil, false, nil
	}

	groups := []string{}
	groupsClaim := claims[s.groupsClaim]
	if groupsClaim != nil {
		groups = interfaceSliceToStringSlice(groupsClaim.([]interface{}))
	}

	// TODO: unpack roles here too?

	resp := &authenticator.Response{
		User: &user.DefaultInfo{
			Name:   claims[s.userIDClaim].(string),
			Groups: groups,
		},
	}

	return resp, true, nil
}
