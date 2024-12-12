package virtualmachines
// Reflects if network is available or deleted from vCenter.
type VirtualMachinesGetResponse_items_networkAdapters_networkDetails_state int

const (
    AVAILABLE_VIRTUALMACHINESGETRESPONSE_ITEMS_NETWORKADAPTERS_NETWORKDETAILS_STATE VirtualMachinesGetResponse_items_networkAdapters_networkDetails_state = iota
    DELETED_VIRTUALMACHINESGETRESPONSE_ITEMS_NETWORKADAPTERS_NETWORKDETAILS_STATE
)

func (i VirtualMachinesGetResponse_items_networkAdapters_networkDetails_state) String() string {
    return []string{"AVAILABLE", "DELETED"}[i]
}
func ParseVirtualMachinesGetResponse_items_networkAdapters_networkDetails_state(v string) (any, error) {
    result := AVAILABLE_VIRTUALMACHINESGETRESPONSE_ITEMS_NETWORKADAPTERS_NETWORKDETAILS_STATE
    switch v {
        case "AVAILABLE":
            result = AVAILABLE_VIRTUALMACHINESGETRESPONSE_ITEMS_NETWORKADAPTERS_NETWORKDETAILS_STATE
        case "DELETED":
            result = DELETED_VIRTUALMACHINESGETRESPONSE_ITEMS_NETWORKADAPTERS_NETWORKDETAILS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_networkAdapters_networkDetails_state(values []VirtualMachinesGetResponse_items_networkAdapters_networkDetails_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_networkAdapters_networkDetails_state) isMultiValue() bool {
    return false
}
