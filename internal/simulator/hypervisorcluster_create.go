// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

//go:embed fixtures/hypervisorclusters/create/asyncfake.json
var clusterAsync string

//go:embed fixtures/hypervisorclusters/create/getfake.json
var clusterGet string

func hypervisorClusterCreate() {
	// TODO: (API) Replace manually create stubs when possible

	taskID := "99992745-532b-404c-860c-20a6571b92c3"
	systemID := "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
	clusterID := "298a299e-78f5-5acb-86ce-4e9fdc290ab7"

	gock.New("http://localhost").
		Post("/private-cloud-business/v1beta1/systems/"+systemID+"/add-hypervisor-cluster").
		Reply(202).
		SetHeader("Location", "/data-services/v1beta1/async-operations/"+taskID)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(clusterAsync)

	gock.New("http://localhost").
		Get("/virtualization/v1beta1/hypervisor-clusters/"+clusterID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(clusterGet)
}

func HypervisorCluster() {
	hypervisorClusterCreate()
}
