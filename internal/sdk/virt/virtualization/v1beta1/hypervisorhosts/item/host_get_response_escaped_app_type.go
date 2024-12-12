package item
// Application type of this asset
type HostGetResponse_appType int

const (
    VMWARE_HOSTGETRESPONSE_APPTYPE HostGetResponse_appType = iota
    HPE_VM_HOSTGETRESPONSE_APPTYPE
)

func (i HostGetResponse_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseHostGetResponse_appType(v string) (any, error) {
    result := VMWARE_HOSTGETRESPONSE_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_HOSTGETRESPONSE_APPTYPE
        case "HPE_VM":
            result = HPE_VM_HOSTGETRESPONSE_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostGetResponse_appType(values []HostGetResponse_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostGetResponse_appType) isMultiValue() bool {
    return false
}
