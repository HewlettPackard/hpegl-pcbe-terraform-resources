// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package simulator

import (
	_ "embed"

	"github.com/h2non/gock"
)

//go:embed fixtures/datastores/delete/delete.json
var dsDelete string

//go:embed fixtures/datastores/delete/async1.json
var dsDeleteAsync1 string

//go:embed fixtures/datastores/delete/async2.json
var dsDeleteAsync2 string

//go:embed fixtures/datastores/delete/async3.json
var dsDeleteAsync3 string

//go:embed fixtures/datastores/delete/async4.json
var dsDeleteAsync4 string

func datastoreDelete() {
	taskID := "5a7b6b4d-8327-48ec-a5ca-98215a24bd51"
	datastoreID := "698de955-87b5-5fe6-b683-78c3948beede"

	gock.New("http://localhost").
		Delete("/virtualization/v1beta1/datastores/"+datastoreID).
		Reply(202).
		SetHeader("Location", "/data-services/v1beta1/async-operations/"+taskID).
		SetHeader("Content-Type", "application/json").
		BodyString(dsDelete)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsDeleteAsync1)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsDeleteAsync2)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsDeleteAsync3)

	gock.New("http://localhost").
		Get("/data-services/v1beta1/async-operations/"+taskID).
		Reply(200).
		SetHeader("Content-Type", "application/json").
		BodyString(dsDeleteAsync4)
}
