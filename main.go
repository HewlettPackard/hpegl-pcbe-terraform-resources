// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false,
		"set to true to run the provider with support for debuggers like delve",
	)
	flag.Parse()

	_ = providerserver.ServeOpts{
		Address: "github.com/HewlettPackard/hpegl-pcbe-terraform-resources",
		Debug:   debug,
	}
}
