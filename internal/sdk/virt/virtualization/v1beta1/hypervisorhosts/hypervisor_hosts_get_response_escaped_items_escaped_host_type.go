package hypervisorhosts
// The type of the hypervisor host.
type HypervisorHostsGetResponse_items_hostType int

const (
    ESXI_HYPERVISORHOSTSGETRESPONSE_ITEMS_HOSTTYPE HypervisorHostsGetResponse_items_hostType = iota
    HPE_VM_HYPERVISORHOSTSGETRESPONSE_ITEMS_HOSTTYPE
)

func (i HypervisorHostsGetResponse_items_hostType) String() string {
    return []string{"ESXI", "HPE_VM"}[i]
}
func ParseHypervisorHostsGetResponse_items_hostType(v string) (any, error) {
    result := ESXI_HYPERVISORHOSTSGETRESPONSE_ITEMS_HOSTTYPE
    switch v {
        case "ESXI":
            result = ESXI_HYPERVISORHOSTSGETRESPONSE_ITEMS_HOSTTYPE
        case "HPE_VM":
            result = HPE_VM_HYPERVISORHOSTSGETRESPONSE_ITEMS_HOSTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorHostsGetResponse_items_hostType(values []HypervisorHostsGetResponse_items_hostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorHostsGetResponse_items_hostType) isMultiValue() bool {
    return false
}
