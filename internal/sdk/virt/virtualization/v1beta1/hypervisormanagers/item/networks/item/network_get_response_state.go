package item
// The current state of the hypervisor network.
type NetworkGetResponse_state int

const (
    OK_NETWORKGETRESPONSE_STATE NetworkGetResponse_state = iota
    ERROR_NETWORKGETRESPONSE_STATE
    REFRESHING_NETWORKGETRESPONSE_STATE
)

func (i NetworkGetResponse_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseNetworkGetResponse_state(v string) (any, error) {
    result := OK_NETWORKGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_NETWORKGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_NETWORKGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_NETWORKGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworkGetResponse_state(values []NetworkGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworkGetResponse_state) isMultiValue() bool {
    return false
}
