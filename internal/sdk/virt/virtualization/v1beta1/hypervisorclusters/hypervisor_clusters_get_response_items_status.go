package hypervisorclusters
// The current status of the hypervisor cluster. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type HypervisorClustersGetResponse_items_status int

const (
    OK_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS HypervisorClustersGetResponse_items_status = iota
    ERROR_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS
    WARNING_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS
)

func (i HypervisorClustersGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseHypervisorClustersGetResponse_items_status(v string) (any, error) {
    result := OK_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorClustersGetResponse_items_status(values []HypervisorClustersGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorClustersGetResponse_items_status) isMultiValue() bool {
    return false
}
