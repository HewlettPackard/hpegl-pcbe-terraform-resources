package item
// Application type of this asset
type HypervisorGetResponse_appType int

const (
    VMWARE_HYPERVISORGETRESPONSE_APPTYPE HypervisorGetResponse_appType = iota
    HPE_VM_HYPERVISORGETRESPONSE_APPTYPE
)

func (i HypervisorGetResponse_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseHypervisorGetResponse_appType(v string) (any, error) {
    result := VMWARE_HYPERVISORGETRESPONSE_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_HYPERVISORGETRESPONSE_APPTYPE
        case "HPE_VM":
            result = HPE_VM_HYPERVISORGETRESPONSE_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorGetResponse_appType(values []HypervisorGetResponse_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorGetResponse_appType) isMultiValue() bool {
    return false
}
