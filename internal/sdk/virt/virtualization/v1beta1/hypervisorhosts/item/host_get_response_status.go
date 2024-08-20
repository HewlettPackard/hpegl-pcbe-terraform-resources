package item
// The current status of the hypervisor host. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type HostGetResponse_status int

const (
    OK_HOSTGETRESPONSE_STATUS HostGetResponse_status = iota
    ERROR_HOSTGETRESPONSE_STATUS
    WARNING_HOSTGETRESPONSE_STATUS
)

func (i HostGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseHostGetResponse_status(v string) (any, error) {
    result := OK_HOSTGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_HOSTGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_HOSTGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_HOSTGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostGetResponse_status(values []HostGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostGetResponse_status) isMultiValue() bool {
    return false
}
