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


resource "hpegl_pc_server" "test" {
        system_id   = "126fd201-9e6e-5e31-9ffb-a766265b1fd3"
        esx_root_credential_id  = "cccfcad1-85b7-4162-b16e-f7cadc2c46b5"
        ilo_admin_credential_id = "dddfcad1-85b7-4162-b16e-f7cadc2c46b5"
        hypervisor_host = {
                hypervisor_cluster_id = "acd4daea-e5e3-5f35-8be3-ce4a4b6d946c"
                hypervisor_host_ip = "16.182.105.217"
        }
        server_network = [
                {
                        data_ip_infos = [
                                {
                                        ip_address = "16.182.105.217"
                                }
                        ]
                }
        ]
}
