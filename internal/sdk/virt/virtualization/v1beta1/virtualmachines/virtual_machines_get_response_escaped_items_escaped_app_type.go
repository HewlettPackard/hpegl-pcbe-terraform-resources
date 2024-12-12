package virtualmachines
// Type of the application to which the VM belongs.
type VirtualMachinesGetResponse_items_appType int

const (
    VMWARE_VIRTUALMACHINESGETRESPONSE_ITEMS_APPTYPE VirtualMachinesGetResponse_items_appType = iota
    HPE_VM_VIRTUALMACHINESGETRESPONSE_ITEMS_APPTYPE
)

func (i VirtualMachinesGetResponse_items_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseVirtualMachinesGetResponse_items_appType(v string) (any, error) {
    result := VMWARE_VIRTUALMACHINESGETRESPONSE_ITEMS_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_VIRTUALMACHINESGETRESPONSE_ITEMS_APPTYPE
        case "HPE_VM":
            result = HPE_VM_VIRTUALMACHINESGETRESPONSE_ITEMS_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_appType(values []VirtualMachinesGetResponse_items_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_appType) isMultiValue() bool {
    return false
}
