package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Roles = []string

type Identity struct {
	ID    string `json:"id"`
	Roles Roles  `json:"roles"`
}

func getIdentityByEmail(svcUrl string, serviceToken string, idTokenHeader string, email string) *Identity {
	svcByteString := []byte(svcUrl)
	url := fmt.Sprintf("%s/%s", bytes.TrimRight(svcByteString, "/"), email)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set(idTokenHeader, fmt.Sprintf("Bearer %s", serviceToken))
	req.Header.Set("Content-Type", "application/json")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer r.Body.Close()

	identity := new(Identity)
	json.NewDecoder(r.Body).Decode(identity)

	return identity
}
