package dataservicesconnectors
// Hypervisor platform.
type DataServicesConnectorsGetResponse_items_platform int

const (
    VMWARE_ESXI_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_PLATFORM DataServicesConnectorsGetResponse_items_platform = iota
)

func (i DataServicesConnectorsGetResponse_items_platform) String() string {
    return []string{"VMWARE_ESXI"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_platform(v string) (any, error) {
    result := VMWARE_ESXI_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_PLATFORM
    switch v {
        case "VMWARE_ESXI":
            result = VMWARE_ESXI_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_PLATFORM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_platform(values []DataServicesConnectorsGetResponse_items_platform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_platform) isMultiValue() bool {
    return false
}
