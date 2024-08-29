// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

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

//go:embed fixtures/datastores/create/filter.json
var dsFilter string

//go:embed fixtures/datastores/create/get.json
var dsGet string

func datastoreCreate() {
	taskID := "be55685c-f84f-4ad5-a3d1-2d7692ed47b1"
	datastoreID := "698de955-87b5-5fe6-b683-78c3948beede"
	datastoreName := "mclaren-ds19"

	gock.New("http://localhost").
		Post("/virtualization/v1beta1/datastore").
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
		MatchParam("filter", "name eq "+datastoreName).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsFilter)

	gock.New("http://localhost").
		Get("/virtualization/v1beta1/datastores/"+datastoreID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsGet)
}
