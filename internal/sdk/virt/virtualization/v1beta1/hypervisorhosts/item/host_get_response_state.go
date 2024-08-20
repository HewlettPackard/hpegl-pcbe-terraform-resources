package item
// The current state of the hypervisor host object
type HostGetResponse_state int

const (
    OK_HOSTGETRESPONSE_STATE HostGetResponse_state = iota
    ERROR_HOSTGETRESPONSE_STATE
    REFRESHING_HOSTGETRESPONSE_STATE
)

func (i HostGetResponse_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseHostGetResponse_state(v string) (any, error) {
    result := OK_HOSTGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_HOSTGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_HOSTGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_HOSTGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostGetResponse_state(values []HostGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostGetResponse_state) isMultiValue() bool {
    return false
}
