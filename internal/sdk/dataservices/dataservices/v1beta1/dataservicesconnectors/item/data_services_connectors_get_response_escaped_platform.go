package item
// Hypervisor platform.
type DataServicesConnectorsGetResponse_platform int

const (
    VMWARE_ESXI_DATASERVICESCONNECTORSGETRESPONSE_PLATFORM DataServicesConnectorsGetResponse_platform = iota
)

func (i DataServicesConnectorsGetResponse_platform) String() string {
    return []string{"VMWARE_ESXI"}[i]
}
func ParseDataServicesConnectorsGetResponse_platform(v string) (any, error) {
    result := VMWARE_ESXI_DATASERVICESCONNECTORSGETRESPONSE_PLATFORM
    switch v {
        case "VMWARE_ESXI":
            result = VMWARE_ESXI_DATASERVICESCONNECTORSGETRESPONSE_PLATFORM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_platform(values []DataServicesConnectorsGetResponse_platform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_platform) isMultiValue() bool {
    return false
}
