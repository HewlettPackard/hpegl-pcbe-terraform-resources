package item
// The current status of the hypervisor cluster. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type ClusterGetResponse_status int

const (
    OK_CLUSTERGETRESPONSE_STATUS ClusterGetResponse_status = iota
    ERROR_CLUSTERGETRESPONSE_STATUS
    WARNING_CLUSTERGETRESPONSE_STATUS
)

func (i ClusterGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseClusterGetResponse_status(v string) (any, error) {
    result := OK_CLUSTERGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_CLUSTERGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_CLUSTERGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_CLUSTERGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeClusterGetResponse_status(values []ClusterGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClusterGetResponse_status) isMultiValue() bool {
    return false
}
