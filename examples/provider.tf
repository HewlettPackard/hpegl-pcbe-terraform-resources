#(C) Copyright 2024 Hewlett Packard Enterprise Development LP

terraform {
  required_providers {
    hpegl = {
      source = "github.com/HewlettPackard/hpegl-pcbe-terraform-resources"
    }
  }
}

provider "hpegl" {
  pc {
    host            = "http://localhost:8080"
    token           = "abcdefghijklmnopqrstuvwxyz-0123456789"

    http_dump       = true
    # poll_interval in seconds
    poll_interval   = 10.0
    max_polls       = 10
  }
}
