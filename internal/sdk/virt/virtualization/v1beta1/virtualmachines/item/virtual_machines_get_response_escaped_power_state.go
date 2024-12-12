package item
// This provides the information power state of the virtual machine.
type VirtualMachinesGetResponse_powerState int

const (
    POWERED_ON_VIRTUALMACHINESGETRESPONSE_POWERSTATE VirtualMachinesGetResponse_powerState = iota
    POWERED_OFF_VIRTUALMACHINESGETRESPONSE_POWERSTATE
    SUSPENDED_VIRTUALMACHINESGETRESPONSE_POWERSTATE
    UNKNOWN_VIRTUALMACHINESGETRESPONSE_POWERSTATE
)

func (i VirtualMachinesGetResponse_powerState) String() string {
    return []string{"POWERED_ON", "POWERED_OFF", "SUSPENDED", "UNKNOWN"}[i]
}
func ParseVirtualMachinesGetResponse_powerState(v string) (any, error) {
    result := POWERED_ON_VIRTUALMACHINESGETRESPONSE_POWERSTATE
    switch v {
        case "POWERED_ON":
            result = POWERED_ON_VIRTUALMACHINESGETRESPONSE_POWERSTATE
        case "POWERED_OFF":
            result = POWERED_OFF_VIRTUALMACHINESGETRESPONSE_POWERSTATE
        case "SUSPENDED":
            result = SUSPENDED_VIRTUALMACHINESGETRESPONSE_POWERSTATE
        case "UNKNOWN":
            result = UNKNOWN_VIRTUALMACHINESGETRESPONSE_POWERSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_powerState(values []VirtualMachinesGetResponse_powerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_powerState) isMultiValue() bool {
    return false
}
