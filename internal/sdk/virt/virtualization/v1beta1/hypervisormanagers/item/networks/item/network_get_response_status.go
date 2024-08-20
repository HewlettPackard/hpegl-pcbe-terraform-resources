package item
// The current status of the hypervisor network. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type NetworkGetResponse_status int

const (
    OK_NETWORKGETRESPONSE_STATUS NetworkGetResponse_status = iota
    ERROR_NETWORKGETRESPONSE_STATUS
    WARNING_NETWORKGETRESPONSE_STATUS
)

func (i NetworkGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseNetworkGetResponse_status(v string) (any, error) {
    result := OK_NETWORKGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_NETWORKGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_NETWORKGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_NETWORKGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworkGetResponse_status(values []NetworkGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworkGetResponse_status) isMultiValue() bool {
    return false
}
