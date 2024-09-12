// (C) Copyright 2024 Hewlett Packard Enterprise Development LP
//go:build simulation

package system

import (
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/simulator"
)

func init() {
	simulator.System()
}
