package virtualmachines
// This provides the information power state of the virtual machine.
type VirtualMachinesGetResponse_items_powerState int

const (
    POWERED_ON_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE VirtualMachinesGetResponse_items_powerState = iota
    POWERED_OFF_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
    SUSPENDED_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
    UNKNOWN_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
)

func (i VirtualMachinesGetResponse_items_powerState) String() string {
    return []string{"POWERED_ON", "POWERED_OFF", "SUSPENDED", "UNKNOWN"}[i]
}
func ParseVirtualMachinesGetResponse_items_powerState(v string) (any, error) {
    result := POWERED_ON_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
    switch v {
        case "POWERED_ON":
            result = POWERED_ON_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
        case "POWERED_OFF":
            result = POWERED_OFF_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
        case "SUSPENDED":
            result = SUSPENDED_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
        case "UNKNOWN":
            result = UNKNOWN_VIRTUALMACHINESGETRESPONSE_ITEMS_POWERSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_powerState(values []VirtualMachinesGetResponse_items_powerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_powerState) isMultiValue() bool {
    return false
}
