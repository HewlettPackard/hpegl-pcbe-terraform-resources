package item
// The current status of the hypervisor resource pool. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type PoolGetResponse_status int

const (
    OK_POOLGETRESPONSE_STATUS PoolGetResponse_status = iota
    ERROR_POOLGETRESPONSE_STATUS
    WARNING_POOLGETRESPONSE_STATUS
)

func (i PoolGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParsePoolGetResponse_status(v string) (any, error) {
    result := OK_POOLGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_POOLGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_POOLGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_POOLGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePoolGetResponse_status(values []PoolGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PoolGetResponse_status) isMultiValue() bool {
    return false
}
