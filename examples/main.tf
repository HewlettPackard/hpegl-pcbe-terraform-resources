#(C) Copyright 2024 Hewlett Packard Enterprise Development LP

data "hpegl_pc_system" "my_system" {
  name = "array-5305-grp"
}

resource "hpegl_pc_datastore" "my_datastore" {
        name = "mclaren-ds19"
        datastore_type = "VMFS"
        capacity_in_bytes = 17179869184
        hci_cluster_uuid = data.hpegl_pc_system.my_system.id
        cluster_info = {
                "name": "5305-CL"
        }
}

resource "hpegl_pc_hypervisor_cluster" "my_hypervisor_cluster" {
  name = "mclaren01"
  hci_cluster_uuid = data.hpegl_pc_system.my_system.id
  app_info = {
    vmware = {
      datacenter_info = {
        name = "5305-DC"
      }
    }
  }
}

resource "hpegl_pc_server" "my_server" {
        name = "16.182.105.217"
        system_id = data.hpegl_pc_system.my_system.id
}
