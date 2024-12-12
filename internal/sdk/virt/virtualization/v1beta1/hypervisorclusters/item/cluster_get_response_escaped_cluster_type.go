package item
// The type of the hypervisor cluster.
type ClusterGetResponse_clusterType int

const (
    ESX_CLUSTER_CLUSTERGETRESPONSE_CLUSTERTYPE ClusterGetResponse_clusterType = iota
    HPE_VM_CLUSTER_CLUSTERGETRESPONSE_CLUSTERTYPE
)

func (i ClusterGetResponse_clusterType) String() string {
    return []string{"ESX_CLUSTER", "HPE_VM_CLUSTER"}[i]
}
func ParseClusterGetResponse_clusterType(v string) (any, error) {
    result := ESX_CLUSTER_CLUSTERGETRESPONSE_CLUSTERTYPE
    switch v {
        case "ESX_CLUSTER":
            result = ESX_CLUSTER_CLUSTERGETRESPONSE_CLUSTERTYPE
        case "HPE_VM_CLUSTER":
            result = HPE_VM_CLUSTER_CLUSTERGETRESPONSE_CLUSTERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeClusterGetResponse_clusterType(values []ClusterGetResponse_clusterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClusterGetResponse_clusterType) isMultiValue() bool {
    return false
}
