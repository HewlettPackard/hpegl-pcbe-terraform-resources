// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

//go:embed fixtures/hypervisorclusters/delete/async1.json
var deleteTaskOne string

//go:embed fixtures/hypervisorclusters/delete/async2.json
var deleteTaskTwo string

func hypervisorClusterDelete() {
	taskID := "a73607cc-ae57-4efe-b573-615813e6c77d"
	systemID := "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	clusterID := "298a299e-78f5-5acb-86ce-4e9fdc290ab7"

	gock.New("http://localhost").
		Post("/private-cloud-business/v1beta1/systems/"+systemID+
			"/remove-hypervisor-clusters").
		MatchType("json").
		JSON(map[string][]string{"hypervisor_cluster_ids": {clusterID}}).
		Reply(202).
		SetHeader("Location", "/data-services/v1beta1/async-operations/"+taskID)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(deleteTaskOne)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(deleteTaskTwo)
}
