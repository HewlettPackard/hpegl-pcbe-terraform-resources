// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

// TODO: (API) Replace fake data with real data when possible
//
//go:embed fixtures/systems/getByName.json
var systemsByName string

func simulateSystemGetByName() {
	systemName := "array-5305-grp"

	gock.New("http://localhost").
		Get("/private-cloud-business/v1beta1/systems").
		MatchParam("filter", "name eq "+systemName).
		MatchHeader("Authorization", "Bearer abcdefghijklmnopqrstuvwxyz-0123456789").
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(systemsByName)

	gock.New("http://localhost").
		Get("/private-cloud-business/v1beta1/systems").
		MatchParam("filter", "name eq "+systemName).
		MatchHeader("Authorization", "Bearer expired-token").
		Reply(401).
		SetHeader("Content-Type", "text/plain").
		BodyString("Jwt is not in the form of Header.Payload.Signature " +
			"with two dots and 3 sections")
}

func System() {
	simulateSystemGetByName()
}
