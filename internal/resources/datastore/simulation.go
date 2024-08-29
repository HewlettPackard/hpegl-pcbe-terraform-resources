// (C) Copyright 2024 Hewlett Packard Enterprise Development LP
//go:build simulation

package datastore

import (
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/simulator"
)

func init() {
	simulator.Datastore()
}
