package item
// The current status of the hypervisor manager resource.
type HypervisorGetResponse_status int

const (
    OK_HYPERVISORGETRESPONSE_STATUS HypervisorGetResponse_status = iota
    ERROR_HYPERVISORGETRESPONSE_STATUS
    WARNING_HYPERVISORGETRESPONSE_STATUS
)

func (i HypervisorGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseHypervisorGetResponse_status(v string) (any, error) {
    result := OK_HYPERVISORGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_HYPERVISORGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_HYPERVISORGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_HYPERVISORGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorGetResponse_status(values []HypervisorGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorGetResponse_status) isMultiValue() bool {
    return false
}
