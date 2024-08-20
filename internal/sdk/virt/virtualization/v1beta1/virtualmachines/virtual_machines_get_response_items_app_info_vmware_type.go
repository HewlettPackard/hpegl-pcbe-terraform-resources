package virtualmachines
// Type of the virtual machine. - VMFS - virtual machine which is created from one or more VMFS datastores. - VVOL - virtual machine which is created from a VVOL datastores. - NFS  - virtual machine which is created from a NFS datastores. - VSAN - virtual machine which is created from a VSAN datastores.
type VirtualMachinesGetResponse_items_appInfo_vmware_type int

const (
    VMFS_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE VirtualMachinesGetResponse_items_appInfo_vmware_type = iota
    VVOL_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
    NFS_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
    VSAN_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
)

func (i VirtualMachinesGetResponse_items_appInfo_vmware_type) String() string {
    return []string{"VMFS", "VVOL", "NFS", "VSAN"}[i]
}
func ParseVirtualMachinesGetResponse_items_appInfo_vmware_type(v string) (any, error) {
    result := VMFS_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
    switch v {
        case "VMFS":
            result = VMFS_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
        case "VVOL":
            result = VVOL_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
        case "NFS":
            result = NFS_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
        case "VSAN":
            result = VSAN_VIRTUALMACHINESGETRESPONSE_ITEMS_APPINFO_VMWARE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_appInfo_vmware_type(values []VirtualMachinesGetResponse_items_appInfo_vmware_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_appInfo_vmware_type) isMultiValue() bool {
    return false
}
