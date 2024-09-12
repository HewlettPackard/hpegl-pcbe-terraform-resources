---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hpegl_pc_datastore Resource - hpegl"
subcategory: ""
description: |-
  
---

# hpegl_pc_datastore (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `capacity_in_bytes` (Number) Size of the datastore in bytes.
- `cluster_info` (Attributes) (see [below for nested schema](#nestedatt--cluster_info))
- `datastore_type` (String) Supported datastore types are VMFS or vVOL
- `hci_cluster_uuid` (String) UUID string uniquely identifying the HCI cluster.
- `name` (String) The data store name

### Read-Only

- `datacenters_info` (Attributes List) List of datacenters to which the datastore is presented to. (see [below for nested schema](#nestedatt--datacenters_info))
- `id` (String) UUID string uniquely identifying the datastore

<a id="nestedatt--cluster_info"></a>
### Nested Schema for `cluster_info`

Required:

- `name` (String) Name of the cluster as reported by the hypervisor manager.

Read-Only:

- `id` (String) UUID string uniquely identifying the hypervisor cluster.


<a id="nestedatt--datacenters_info"></a>
### Nested Schema for `datacenters_info`

Read-Only:

- `id` (String) UUID string uniquely identifier of the datacenter.
- `name` (String) VMware provided name for the datacenter.