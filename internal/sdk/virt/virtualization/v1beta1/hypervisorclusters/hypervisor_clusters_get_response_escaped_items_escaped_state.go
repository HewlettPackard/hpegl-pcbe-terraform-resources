package hypervisorclusters
// The current state of the hypervisor cluster object.
type HypervisorClustersGetResponse_items_state int

const (
    OK_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE HypervisorClustersGetResponse_items_state = iota
    ERROR_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE
    REFRESHING_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE
)

func (i HypervisorClustersGetResponse_items_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseHypervisorClustersGetResponse_items_state(v string) (any, error) {
    result := OK_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorClustersGetResponse_items_state(values []HypervisorClustersGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorClustersGetResponse_items_state) isMultiValue() bool {
    return false
}
