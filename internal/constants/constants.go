// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

package constants

const (
	ProviderType         = "hpegl"
	ProviderBlock        = "pc"
	NameFilter           = "name eq "
	HciClusterUUIDFilter = "hciClusterUuid eq "
	AndFilter            = " and "

	TaskHypervisorCluster = "hypervisor-cluster" // task's "associatedResources" string
	TaskHypervisorServer  = "server"             // task's "associatedResources" string
	TaskDatastore         = "datastore"          // task's "associatedResources" string
)
