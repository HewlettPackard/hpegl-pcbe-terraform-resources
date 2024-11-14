package item
// Type of the Hypervisor Cluster
type GetResponse_hypervisorClusters_hypervisorClusterType int

const (
    ESX_CLUSTER_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE GetResponse_hypervisorClusters_hypervisorClusterType = iota
    HPE_VIRTUALIZATION_CLUSTER_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
)

func (i GetResponse_hypervisorClusters_hypervisorClusterType) String() string {
    return []string{"ESX_CLUSTER", "HPE_VIRTUALIZATION_CLUSTER"}[i]
}
func ParseGetResponse_hypervisorClusters_hypervisorClusterType(v string) (any, error) {
    result := ESX_CLUSTER_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
    switch v {
        case "ESX_CLUSTER":
            result = ESX_CLUSTER_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
        case "HPE_VIRTUALIZATION_CLUSTER":
            result = HPE_VIRTUALIZATION_CLUSTER_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_hypervisorClusters_hypervisorClusterType(values []GetResponse_hypervisorClusters_hypervisorClusterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_hypervisorClusters_hypervisorClusterType) isMultiValue() bool {
    return false
}
