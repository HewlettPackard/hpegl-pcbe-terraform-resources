// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

// TODO: (API) Replace fake data with real data when possible
//
//go:embed fixtures/secrets/getByName.json
var secretByName string

func simulateSecretGetByName() {
	secretName := "mysecret1"

	gock.New("http://localhost").
		Get("/data-services/v1beta1/secrets").
		MatchParam("filter", "name eq "+"'"+secretName+"'").
		MatchHeader("Authorization", "Bearer abcdefghijklmnopqrstuvwxyz-0123456789").
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(secretByName)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/secrets").
		MatchParam("filter", "name eq "+"'"+secretName+"'").
		MatchHeader("Authorization", "Bearer expired-token").
		Reply(401).
		SetHeader("Content-Type", "text/plain").
		BodyString("Jwt is not in the form of Header.Payload.Signature " +
			"with two dots and 3 sections")
}

func Secret() {
	simulateSecretGetByName()
}
