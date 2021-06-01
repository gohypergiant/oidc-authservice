package main

const (
	ScopeOpenID         = "openid"
	ScopeSDKDevelopment = "sdk_development"
	ScopeSDKProduction  = "sdk_production"
	ScopeService        = "service"
)

var ValidScopes = []string{ScopeOpenID, ScopeSDKDevelopment, ScopeSDKProduction, ScopeService}
