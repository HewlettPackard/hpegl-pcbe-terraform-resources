// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"flag"
	"log"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var version = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false,
		"set to true to run the provider with support for debuggers like delve",
	)
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "github.com/HewlettPackard/hpegl-pcbe-terraform-resources",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
