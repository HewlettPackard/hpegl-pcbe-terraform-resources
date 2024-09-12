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
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(systemsByName)
}

func System() {
	simulateSystemGetByName()
}
