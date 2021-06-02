module github.com/gohypergiant/oidc-authservice

go 1.16

require (
	github.com/bmartel/boltstore v1.0.1
	github.com/boltdb/bolt v1.3.1
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/emirpasic/gods v1.12.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/tevino/abool v1.2.0
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	gonum.org/v1/gonum v0.9.1
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/apiserver v0.21.1
	k8s.io/client-go v0.21.1
	sigs.k8s.io/controller-runtime v0.8.3
)
