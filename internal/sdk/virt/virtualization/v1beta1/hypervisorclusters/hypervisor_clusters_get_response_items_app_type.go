package hypervisorclusters
// Application type of this asset
type HypervisorClustersGetResponse_items_appType int

const (
    VMWARE_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_APPTYPE HypervisorClustersGetResponse_items_appType = iota
    HPE_VM_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_APPTYPE
)

func (i HypervisorClustersGetResponse_items_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseHypervisorClustersGetResponse_items_appType(v string) (any, error) {
    result := VMWARE_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_APPTYPE
        case "HPE_VM":
            result = HPE_VM_HYPERVISORCLUSTERSGETRESPONSE_ITEMS_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorClustersGetResponse_items_appType(values []HypervisorClustersGetResponse_items_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorClustersGetResponse_items_appType) isMultiValue() bool {
    return false
}
