package item
// Specifies how the MAC address is provided for the adapter.
type VirtualMachinesGetResponse_networkAdapters_macAddressType int

const (
    MANUAL_VIRTUALMACHINESGETRESPONSE_NETWORKADAPTERS_MACADDRESSTYPE VirtualMachinesGetResponse_networkAdapters_macAddressType = iota
    AUTOMATIC_VIRTUALMACHINESGETRESPONSE_NETWORKADAPTERS_MACADDRESSTYPE
)

func (i VirtualMachinesGetResponse_networkAdapters_macAddressType) String() string {
    return []string{"MANUAL", "AUTOMATIC"}[i]
}
func ParseVirtualMachinesGetResponse_networkAdapters_macAddressType(v string) (any, error) {
    result := MANUAL_VIRTUALMACHINESGETRESPONSE_NETWORKADAPTERS_MACADDRESSTYPE
    switch v {
        case "MANUAL":
            result = MANUAL_VIRTUALMACHINESGETRESPONSE_NETWORKADAPTERS_MACADDRESSTYPE
        case "AUTOMATIC":
            result = AUTOMATIC_VIRTUALMACHINESGETRESPONSE_NETWORKADAPTERS_MACADDRESSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_networkAdapters_macAddressType(values []VirtualMachinesGetResponse_networkAdapters_macAddressType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_networkAdapters_macAddressType) isMultiValue() bool {
    return false
}
