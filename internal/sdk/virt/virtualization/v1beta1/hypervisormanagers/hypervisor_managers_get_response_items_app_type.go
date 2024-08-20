package hypervisormanagers
// Application type of this asset
type HypervisorManagersGetResponse_items_appType int

const (
    VMWARE_HYPERVISORMANAGERSGETRESPONSE_ITEMS_APPTYPE HypervisorManagersGetResponse_items_appType = iota
    HPE_VM_HYPERVISORMANAGERSGETRESPONSE_ITEMS_APPTYPE
)

func (i HypervisorManagersGetResponse_items_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseHypervisorManagersGetResponse_items_appType(v string) (any, error) {
    result := VMWARE_HYPERVISORMANAGERSGETRESPONSE_ITEMS_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_HYPERVISORMANAGERSGETRESPONSE_ITEMS_APPTYPE
        case "HPE_VM":
            result = HPE_VM_HYPERVISORMANAGERSGETRESPONSE_ITEMS_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorManagersGetResponse_items_appType(values []HypervisorManagersGetResponse_items_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorManagersGetResponse_items_appType) isMultiValue() bool {
    return false
}
