package virtualmachines
// This gives information this virtual disk type. - VMFS - virtual machine flat disks. - VVOL - virtual volume - PRDM - physical raw disk mapping - VRDM - virtual raw disk mapping
type VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type int

const (
    VMFS_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type = iota
    VVOL_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
    PRDM_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
    VRDM_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
)

func (i VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type) String() string {
    return []string{"VMFS", "VVOL", "PRDM", "VRDM"}[i]
}
func ParseVirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type(v string) (any, error) {
    result := VMFS_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
    switch v {
        case "VMFS":
            result = VMFS_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
        case "VVOL":
            result = VVOL_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
        case "PRDM":
            result = PRDM_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
        case "VRDM":
            result = VRDM_VIRTUALMACHINESGETRESPONSE_ITEMS_VIRTUALDISKS_APPINFO_VMWARE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type(values []VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_virtualDisks_appInfo_vmware_type) isMultiValue() bool {
    return false
}
