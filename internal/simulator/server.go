// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

// TODO: (API) Replace fake data with real data when possible
//
//go:embed fixtures/servers/create/post.json
var serverPost string

//go:embed fixtures/servers/create/get.json
var serverGet string

// TODO: (API) Replace fake data with real data when possible
//
//go:embed fixtures/servers/create/task.json
var serverTask string

func simulateServerCreate() {
	systemID := "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	serverID := "697e8cbf-df7e-570c-a3c7-912d4ce8375a"
	taskID := "b1b1b1b1-b1b1-b1b1-b1b1-b1b1b1b1b1b1"

	gock.New("http://localhost").
		Post("/private-cloud-business/v1beta1/systems/"+
			systemID+"/add-hypervisor-servers").
		MatchType("json").
		BodyString(serverPost).
		Reply(202).
		SetHeader("Location", "/data-services/v1beta1/async-operations/"+taskID)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(serverTask)

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
