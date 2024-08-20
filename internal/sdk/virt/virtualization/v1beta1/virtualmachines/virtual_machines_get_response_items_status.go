package virtualmachines
// The current status of the virtual machine.
type VirtualMachinesGetResponse_items_status int

const (
    OK_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS VirtualMachinesGetResponse_items_status = iota
    ERROR_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS
    WARNING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS
)

func (i VirtualMachinesGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseVirtualMachinesGetResponse_items_status(v string) (any, error) {
    result := OK_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_status(values []VirtualMachinesGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_status) isMultiValue() bool {
    return false
}
