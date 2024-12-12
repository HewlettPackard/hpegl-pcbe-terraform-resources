package item
// Indicates whether the address is assigned statically or via DHCP.
type DataServicesConnectorsGetResponse_interfaces_network_nic_addressType int

const (
    DHCP_DATASERVICESCONNECTORSGETRESPONSE_INTERFACES_NETWORK_NIC_ADDRESSTYPE DataServicesConnectorsGetResponse_interfaces_network_nic_addressType = iota
    STATIC_DATASERVICESCONNECTORSGETRESPONSE_INTERFACES_NETWORK_NIC_ADDRESSTYPE
)

func (i DataServicesConnectorsGetResponse_interfaces_network_nic_addressType) String() string {
    return []string{"DHCP", "STATIC"}[i]
}
func ParseDataServicesConnectorsGetResponse_interfaces_network_nic_addressType(v string) (any, error) {
    result := DHCP_DATASERVICESCONNECTORSGETRESPONSE_INTERFACES_NETWORK_NIC_ADDRESSTYPE
    switch v {
        case "DHCP":
            result = DHCP_DATASERVICESCONNECTORSGETRESPONSE_INTERFACES_NETWORK_NIC_ADDRESSTYPE
        case "STATIC":
            result = STATIC_DATASERVICESCONNECTORSGETRESPONSE_INTERFACES_NETWORK_NIC_ADDRESSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_interfaces_network_nic_addressType(values []DataServicesConnectorsGetResponse_interfaces_network_nic_addressType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_interfaces_network_nic_addressType) isMultiValue() bool {
    return false
}
