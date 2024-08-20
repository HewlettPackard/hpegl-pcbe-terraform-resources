package hypervisorclusters
// The type of the hypervisor cluster.
type HypervisorClustersGetResponse_items_clusterType int

const (
    ESX_CLUSTER_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_CLUSTERTYPE HypervisorClustersGetResponse_items_clusterType = iota
    HPE_VM_CLUSTER_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_CLUSTERTYPE
)

func (i HypervisorClustersGetResponse_items_clusterType) String() string {
    return []string{"ESX_CLUSTER", "HPE_VM_CLUSTER"}[i]
}
func ParseHypervisorClustersGetResponse_items_clusterType(v string) (any, error) {
    result := ESX_CLUSTER_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_CLUSTERTYPE
    switch v {
        case "ESX_CLUSTER":
            result = ESX_CLUSTER_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_CLUSTERTYPE
        case "HPE_VM_CLUSTER":
            result = HPE_VM_CLUSTER_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_CLUSTERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorClustersGetResponse_items_clusterType(values []HypervisorClustersGetResponse_items_clusterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorClustersGetResponse_items_clusterType) isMultiValue() bool {
    return false
}
