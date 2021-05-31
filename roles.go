package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Roles = []string

func getRolesByEmail(svcUrl string, email string) *Roles {
	svcByteString := []byte(svcUrl)
	url := fmt.Sprintf("%s/%s", bytes.TrimRight("/"), email)

	req, err := http.NewRequest("GET", url)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "<idToken>"))
	req.Header.Set("Content-Type", "application/json")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	roles := new(Roles)
	json.NewDecoder(r.Body).Decode(roles)

	return roles
}
