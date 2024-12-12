package servers
// Type of the hypervisor host.
type ServersGetResponse_items_hypervisorHost_hypervisorHostType int

const (
    ESXI_SERVERSGETRESPONSE_ITEMS_HYPERVISORHOST_HYPERVISORHOSTTYPE ServersGetResponse_items_hypervisorHost_hypervisorHostType = iota
    HPE_VIRTUALIZATION_SERVERSGETRESPONSE_ITEMS_HYPERVISORHOST_HYPERVISORHOSTTYPE
)

func (i ServersGetResponse_items_hypervisorHost_hypervisorHostType) String() string {
    return []string{"ESXI", "HPE_VIRTUALIZATION"}[i]
}
func ParseServersGetResponse_items_hypervisorHost_hypervisorHostType(v string) (any, error) {
    result := ESXI_SERVERSGETRESPONSE_ITEMS_HYPERVISORHOST_HYPERVISORHOSTTYPE
    switch v {
        case "ESXI":
            result = ESXI_SERVERSGETRESPONSE_ITEMS_HYPERVISORHOST_HYPERVISORHOSTTYPE
        case "HPE_VIRTUALIZATION":
            result = HPE_VIRTUALIZATION_SERVERSGETRESPONSE_ITEMS_HYPERVISORHOST_HYPERVISORHOSTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_hypervisorHost_hypervisorHostType(values []ServersGetResponse_items_hypervisorHost_hypervisorHostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_hypervisorHost_hypervisorHostType) isMultiValue() bool {
    return false
}
