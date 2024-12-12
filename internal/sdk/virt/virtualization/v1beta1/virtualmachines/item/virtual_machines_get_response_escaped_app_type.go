package item
// Type of the application to which the VM belongs.
type VirtualMachinesGetResponse_appType int

const (
    VMWARE_VIRTUALMACHINESGETRESPONSE_APPTYPE VirtualMachinesGetResponse_appType = iota
    HPE_VM_VIRTUALMACHINESGETRESPONSE_APPTYPE
)

func (i VirtualMachinesGetResponse_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseVirtualMachinesGetResponse_appType(v string) (any, error) {
    result := VMWARE_VIRTUALMACHINESGETRESPONSE_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_VIRTUALMACHINESGETRESPONSE_APPTYPE
        case "HPE_VM":
            result = HPE_VM_VIRTUALMACHINESGETRESPONSE_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_appType(values []VirtualMachinesGetResponse_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_appType) isMultiValue() bool {
    return false
}
