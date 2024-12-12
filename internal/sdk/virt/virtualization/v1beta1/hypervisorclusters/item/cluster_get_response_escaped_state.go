package item
// The current state of the hypervisor cluster object.
type ClusterGetResponse_state int

const (
    OK_CLUSTERGETRESPONSE_STATE ClusterGetResponse_state = iota
    ERROR_CLUSTERGETRESPONSE_STATE
    REFRESHING_CLUSTERGETRESPONSE_STATE
)

func (i ClusterGetResponse_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseClusterGetResponse_state(v string) (any, error) {
    result := OK_CLUSTERGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_CLUSTERGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_CLUSTERGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_CLUSTERGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeClusterGetResponse_state(values []ClusterGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClusterGetResponse_state) isMultiValue() bool {
    return false
}
