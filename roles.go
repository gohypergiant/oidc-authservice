package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Roles = []string

func getRolesByEmail(svcUrl string, serviceToken string, idTokenHeader string, email string) *Roles {
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

	roles := new(Roles)
	json.NewDecoder(r.Body).Decode(roles)

	return roles
}
