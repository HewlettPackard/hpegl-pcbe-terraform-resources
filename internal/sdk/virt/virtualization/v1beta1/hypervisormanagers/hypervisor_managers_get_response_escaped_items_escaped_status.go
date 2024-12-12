package hypervisormanagers
// The current status of the hypervisor manager resource.
type HypervisorManagersGetResponse_items_status int

const (
    OK_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS HypervisorManagersGetResponse_items_status = iota
    ERROR_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS
    WARNING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS
)

func (i HypervisorManagersGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseHypervisorManagersGetResponse_items_status(v string) (any, error) {
    result := OK_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorManagersGetResponse_items_status(values []HypervisorManagersGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorManagersGetResponse_items_status) isMultiValue() bool {
    return false
}
