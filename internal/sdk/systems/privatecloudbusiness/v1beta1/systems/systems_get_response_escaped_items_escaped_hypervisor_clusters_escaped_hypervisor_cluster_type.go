package systems
// Type of the Hypervisor Cluster
type SystemsGetResponse_items_hypervisorClusters_hypervisorClusterType int

const (
    ESX_CLUSTER_SYSTEMSGETRESPONSE_ITEMS_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE SystemsGetResponse_items_hypervisorClusters_hypervisorClusterType = iota
    HPE_VIRTUALIZATION_CLUSTER_SYSTEMSGETRESPONSE_ITEMS_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
)

func (i SystemsGetResponse_items_hypervisorClusters_hypervisorClusterType) String() string {
    return []string{"ESX_CLUSTER", "HPE_VIRTUALIZATION_CLUSTER"}[i]
}
func ParseSystemsGetResponse_items_hypervisorClusters_hypervisorClusterType(v string) (any, error) {
    result := ESX_CLUSTER_SYSTEMSGETRESPONSE_ITEMS_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
    switch v {
        case "ESX_CLUSTER":
            result = ESX_CLUSTER_SYSTEMSGETRESPONSE_ITEMS_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
        case "HPE_VIRTUALIZATION_CLUSTER":
            result = HPE_VIRTUALIZATION_CLUSTER_SYSTEMSGETRESPONSE_ITEMS_HYPERVISORCLUSTERS_HYPERVISORCLUSTERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_hypervisorClusters_hypervisorClusterType(values []SystemsGetResponse_items_hypervisorClusters_hypervisorClusterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_hypervisorClusters_hypervisorClusterType) isMultiValue() bool {
    return false
}
