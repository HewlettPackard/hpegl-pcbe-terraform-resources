package hypervisorhosts
// Application type of this asset
type HypervisorHostsGetResponse_items_appType int

const (
    VMWARE_HYPERVISORHOSTSGETRESPONSE_ITEMS_APPTYPE HypervisorHostsGetResponse_items_appType = iota
    HPE_VM_HYPERVISORHOSTSGETRESPONSE_ITEMS_APPTYPE
)

func (i HypervisorHostsGetResponse_items_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseHypervisorHostsGetResponse_items_appType(v string) (any, error) {
    result := VMWARE_HYPERVISORHOSTSGETRESPONSE_ITEMS_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_HYPERVISORHOSTSGETRESPONSE_ITEMS_APPTYPE
        case "HPE_VM":
            result = HPE_VM_HYPERVISORHOSTSGETRESPONSE_ITEMS_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorHostsGetResponse_items_appType(values []HypervisorHostsGetResponse_items_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorHostsGetResponse_items_appType) isMultiValue() bool {
    return false
}
