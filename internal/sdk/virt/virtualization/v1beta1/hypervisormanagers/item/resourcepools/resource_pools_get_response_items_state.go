package resourcepools
// The current state of the hypervisor resource pool.
type ResourcePoolsGetResponse_items_state int

const (
    OK_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE ResourcePoolsGetResponse_items_state = iota
    ERROR_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE
    REFRESHING_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE
)

func (i ResourcePoolsGetResponse_items_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseResourcePoolsGetResponse_items_state(v string) (any, error) {
    result := OK_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_RESOURCEPOOLSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeResourcePoolsGetResponse_items_state(values []ResourcePoolsGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ResourcePoolsGetResponse_items_state) isMultiValue() bool {
    return false
}
