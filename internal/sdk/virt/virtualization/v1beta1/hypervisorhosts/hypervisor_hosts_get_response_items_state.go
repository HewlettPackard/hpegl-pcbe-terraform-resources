package hypervisorhosts
// The current state of the hypervisor host object
type HypervisorHostsGetResponse_items_state int

const (
    OK_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE HypervisorHostsGetResponse_items_state = iota
    ERROR_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE
    REFRESHING_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE
)

func (i HypervisorHostsGetResponse_items_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseHypervisorHostsGetResponse_items_state(v string) (any, error) {
    result := OK_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorHostsGetResponse_items_state(values []HypervisorHostsGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorHostsGetResponse_items_state) isMultiValue() bool {
    return false
}
