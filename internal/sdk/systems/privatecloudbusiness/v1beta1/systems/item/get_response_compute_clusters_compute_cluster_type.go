package item
// Type of compute cluster.
type GetResponse_computeClusters_computeClusterType int

const (
    HYPERVISOR_VMWARE_ESXI_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE GetResponse_computeClusters_computeClusterType = iota
    HYPERVISOR_HPE_VIRTUALIZATION_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
    K8S_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
    K8S_AI_WORKLOAD_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
    K8S_AI_CONTROL_PLANE_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
)

func (i GetResponse_computeClusters_computeClusterType) String() string {
    return []string{"HYPERVISOR_VMWARE_ESXI", "HYPERVISOR_HPE_VIRTUALIZATION", "K8S", "K8S_AI_WORKLOAD", "K8S_AI_CONTROL_PLANE"}[i]
}
func ParseGetResponse_computeClusters_computeClusterType(v string) (any, error) {
    result := HYPERVISOR_VMWARE_ESXI_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
    switch v {
        case "HYPERVISOR_VMWARE_ESXI":
            result = HYPERVISOR_VMWARE_ESXI_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
        case "HYPERVISOR_HPE_VIRTUALIZATION":
            result = HYPERVISOR_HPE_VIRTUALIZATION_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
        case "K8S":
            result = K8S_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
        case "K8S_AI_WORKLOAD":
            result = K8S_AI_WORKLOAD_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
        case "K8S_AI_CONTROL_PLANE":
            result = K8S_AI_CONTROL_PLANE_GETRESPONSE_COMPUTECLUSTERS_COMPUTECLUSTERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_computeClusters_computeClusterType(values []GetResponse_computeClusters_computeClusterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_computeClusters_computeClusterType) isMultiValue() bool {
    return false
}
