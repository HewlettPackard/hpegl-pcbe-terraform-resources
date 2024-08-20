package hypervisorhosts
// The current status of the hypervisor host. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type HypervisorHostsGetResponse_items_status int

const (
    OK_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS HypervisorHostsGetResponse_items_status = iota
    ERROR_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS
    WARNING_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS
)

func (i HypervisorHostsGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseHypervisorHostsGetResponse_items_status(v string) (any, error) {
    result := OK_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_HYPERVISORHOSTSGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorHostsGetResponse_items_status(values []HypervisorHostsGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorHostsGetResponse_items_status) isMultiValue() bool {
    return false
}
