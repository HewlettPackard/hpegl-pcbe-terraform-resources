// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

//go:embed fixtures/datastores/create/post-request.json
var dsPostRequest string

//go:embed fixtures/datastores/create/hypervisorclustersfiltered.json
var hcFilter string

//go:embed fixtures/datastores/create/async1.json
var dsAsync1 string

//go:embed fixtures/datastores/create/async2.json
var dsAsync2 string

//go:embed fixtures/datastores/create/async3.json
var dsAsync3 string

//go:embed fixtures/datastores/create/async4.json
var dsAsync4 string

//go:embed fixtures/datastores/create/async5.json
var dsAsync5 string

//go:embed fixtures/datastores/create/get.json
var dsGet string

func datastoreCreate() {
	taskID := "be55685c-f84f-4ad5-a3d1-2d7692ed47b1"
	datastoreName := "mclaren-ds19"
	hypervisorClusterID := "298a299e-78f5-5acb-86ce-4e9fdc290ab7" // nolint goconst
	clusterName := "5305-CL"
	systemID := "126fd201-9e6e-5e31-9ffb-a766265b1fd3" // nolint goconst

	gock.New("http://localhost").
		Get("/virtualization/v1beta1/hypervisor-clusters").
		MatchParam("filter", "hciClusterUuid eq "+hypervisorClusterID+
			" and name eq "+clusterName).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(hcFilter)

	gock.New("http://localhost").
		Post("/virtualization/v1beta1/datastore").
		MatchType("json").
		BodyString(dsPostRequest).
		Reply(202).
		SetHeader("Location", "/data-services/v1beta1/async-operations/"+taskID)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsAsync1)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsAsync2)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsAsync3)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsAsync4)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsAsync5)

	gock.New("http://localhost").
		Get("/virtualization/v1beta1/datastores").
		MatchParam("filter", "hciClusterUuid eq "+systemID+
			" and name eq "+datastoreName+
			" and clusterInfo.id eq "+hypervisorClusterID).
		MatchHeader("Authorization", "Bearer abcdefghijklmnopqrstuvwxyz-0123456789").
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsGet)

	gock.New("http://localhost").
		Get("/virtualization/v1beta1/datastores").
		MatchParam("filter", "hciClusterUuid eq "+systemID+
			" and name eq "+datastoreName+
			" and clusterInfo.id eq "+hypervisorClusterID).
		MatchHeader("Authorization", "Bearer missing-datastore").
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(`{"items":null,"count":0,"offset":0,"total":0}`)
}
