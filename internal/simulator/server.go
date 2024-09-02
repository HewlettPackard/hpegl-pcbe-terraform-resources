// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

// TODO: (API) Replace fake data with real data when possible
//
//go:embed fixtures/servers/create/get.json
var serverGet string

func simulateServerCreate() {
	systemID := "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	serverID := "697e8cbf-df7e-570c-a3c7-912d4ce8375a"

	gock.New("http://localhost").
		Get("/private-cloud-business/v1beta1/systems/"+
			systemID+"/servers/"+serverID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(serverGet)
}

func Server() {
	simulateServerCreate()
}
