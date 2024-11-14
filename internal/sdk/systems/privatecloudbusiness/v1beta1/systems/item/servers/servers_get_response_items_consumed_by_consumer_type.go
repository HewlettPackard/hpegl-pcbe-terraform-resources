package servers
// Consumer type.
type ServersGetResponse_items_consumedBy_consumerType int

const (
    HYPERVISOR_CLUSTER_VMWARE_ESXI_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE ServersGetResponse_items_consumedBy_consumerType = iota
    HYPERVISOR_CLUSTER_HPE_VIRTUALIZATION_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
    K8S_CLUSTER_AI_WORKLOAD_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
    K8S_CLUSTER_AI_CONTROL_PLANE_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
    K8S_CLUSTER_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
)

func (i ServersGetResponse_items_consumedBy_consumerType) String() string {
    return []string{"HYPERVISOR_CLUSTER_VMWARE_ESXI", "HYPERVISOR_CLUSTER_HPE_VIRTUALIZATION", "K8S_CLUSTER_AI_WORKLOAD", "K8S_CLUSTER_AI_CONTROL_PLANE", "K8S_CLUSTER"}[i]
}
func ParseServersGetResponse_items_consumedBy_consumerType(v string) (any, error) {
    result := HYPERVISOR_CLUSTER_VMWARE_ESXI_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
    switch v {
        case "HYPERVISOR_CLUSTER_VMWARE_ESXI":
            result = HYPERVISOR_CLUSTER_VMWARE_ESXI_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
        case "HYPERVISOR_CLUSTER_HPE_VIRTUALIZATION":
            result = HYPERVISOR_CLUSTER_HPE_VIRTUALIZATION_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
        case "K8S_CLUSTER_AI_WORKLOAD":
            result = K8S_CLUSTER_AI_WORKLOAD_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
        case "K8S_CLUSTER_AI_CONTROL_PLANE":
            result = K8S_CLUSTER_AI_CONTROL_PLANE_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
        case "K8S_CLUSTER":
            result = K8S_CLUSTER_SERVERSGETRESPONSE_ITEMS_CONSUMEDBY_CONSUMERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_consumedBy_consumerType(values []ServersGetResponse_items_consumedBy_consumerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_consumedBy_consumerType) isMultiValue() bool {
    return false
}
