package item
// The type of the hypervisor host.
type HostGetResponse_hostType int

const (
    ESXI_HOSTGETRESPONSE_HOSTTYPE HostGetResponse_hostType = iota
    HPE_VM_HOSTGETRESPONSE_HOSTTYPE
)

func (i HostGetResponse_hostType) String() string {
    return []string{"ESXI", "HPE_VM"}[i]
}
func ParseHostGetResponse_hostType(v string) (any, error) {
    result := ESXI_HOSTGETRESPONSE_HOSTTYPE
    switch v {
        case "ESXI":
            result = ESXI_HOSTGETRESPONSE_HOSTTYPE
        case "HPE_VM":
            result = HPE_VM_HOSTGETRESPONSE_HOSTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostGetResponse_hostType(values []HostGetResponse_hostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostGetResponse_hostType) isMultiValue() bool {
    return false
}
