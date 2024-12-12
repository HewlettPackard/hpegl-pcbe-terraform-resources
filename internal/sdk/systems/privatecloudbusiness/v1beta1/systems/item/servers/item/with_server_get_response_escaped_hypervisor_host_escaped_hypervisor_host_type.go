package item
// Type of the hypervisor host.
type WithServerGetResponse_hypervisorHost_hypervisorHostType int

const (
    ESXI_WITHSERVERGETRESPONSE_HYPERVISORHOST_HYPERVISORHOSTTYPE WithServerGetResponse_hypervisorHost_hypervisorHostType = iota
    HPE_VIRTUALIZATION_WITHSERVERGETRESPONSE_HYPERVISORHOST_HYPERVISORHOSTTYPE
)

func (i WithServerGetResponse_hypervisorHost_hypervisorHostType) String() string {
    return []string{"ESXI", "HPE_VIRTUALIZATION"}[i]
}
func ParseWithServerGetResponse_hypervisorHost_hypervisorHostType(v string) (any, error) {
    result := ESXI_WITHSERVERGETRESPONSE_HYPERVISORHOST_HYPERVISORHOSTTYPE
    switch v {
        case "ESXI":
            result = ESXI_WITHSERVERGETRESPONSE_HYPERVISORHOST_HYPERVISORHOSTTYPE
        case "HPE_VIRTUALIZATION":
            result = HPE_VIRTUALIZATION_WITHSERVERGETRESPONSE_HYPERVISORHOST_HYPERVISORHOSTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_hypervisorHost_hypervisorHostType(values []WithServerGetResponse_hypervisorHost_hypervisorHostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_hypervisorHost_hypervisorHostType) isMultiValue() bool {
    return false
}
