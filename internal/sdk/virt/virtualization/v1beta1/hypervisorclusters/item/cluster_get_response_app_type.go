package item
// Application type of this asset
type ClusterGetResponse_appType int

const (
    VMWARE_CLUSTERGETRESPONSE_APPTYPE ClusterGetResponse_appType = iota
    HPE_VM_CLUSTERGETRESPONSE_APPTYPE
)

func (i ClusterGetResponse_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseClusterGetResponse_appType(v string) (any, error) {
    result := VMWARE_CLUSTERGETRESPONSE_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_CLUSTERGETRESPONSE_APPTYPE
        case "HPE_VM":
            result = HPE_VM_CLUSTERGETRESPONSE_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeClusterGetResponse_appType(values []ClusterGetResponse_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClusterGetResponse_appType) isMultiValue() bool {
    return false
}
