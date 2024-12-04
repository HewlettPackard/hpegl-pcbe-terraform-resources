// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

//go:embed fixtures/servers/delete/post.json
var serverDeletePost string

//go:embed fixtures/servers/delete/task.json
var serverDeleteTask string

func simulateServerDelete() {
	systemID := "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	taskID := "ccccc106-e22d-4c06-8de7-7550a646b441"

	gock.New("http://localhost").
		Post("/private-cloud-business/v1beta1/systems/"+
			systemID+"/remove-hypervisor-servers").
		MatchType("json").
		BodyString(serverDeletePost).
		Reply(202).
		SetHeader("Location", "/data-services/v1beta1/async-operations/"+taskID)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(serverDeleteTask)
}
