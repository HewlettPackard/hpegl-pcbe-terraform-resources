package item
// The current state of the hypervisor resource pool.
type PoolGetResponse_state int

const (
    OK_POOLGETRESPONSE_STATE PoolGetResponse_state = iota
    ERROR_POOLGETRESPONSE_STATE
    REFRESHING_POOLGETRESPONSE_STATE
)

func (i PoolGetResponse_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParsePoolGetResponse_state(v string) (any, error) {
    result := OK_POOLGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_POOLGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_POOLGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_POOLGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePoolGetResponse_state(values []PoolGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PoolGetResponse_state) isMultiValue() bool {
    return false
}
