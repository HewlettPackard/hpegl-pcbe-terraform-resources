package resourcepools
// The current status of the hypervisor resource pool. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type ResourcePoolsGetResponse_items_status int

const (
    OK_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS ResourcePoolsGetResponse_items_status = iota
    ERROR_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS
    WARNING_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS
)

func (i ResourcePoolsGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseResourcePoolsGetResponse_items_status(v string) (any, error) {
    result := OK_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_RESOURCEPOOLSGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeResourcePoolsGetResponse_items_status(values []ResourcePoolsGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ResourcePoolsGetResponse_items_status) isMultiValue() bool {
    return false
}
