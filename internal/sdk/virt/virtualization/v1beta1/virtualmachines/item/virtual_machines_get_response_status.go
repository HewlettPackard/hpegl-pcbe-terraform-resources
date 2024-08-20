package item
// The current status of the virtual machine.
type VirtualMachinesGetResponse_status int

const (
    OK_VIRTUALMACHINESGETRESPONSE_STATUS VirtualMachinesGetResponse_status = iota
    ERROR_VIRTUALMACHINESGETRESPONSE_STATUS
    WARNING_VIRTUALMACHINESGETRESPONSE_STATUS
)

func (i VirtualMachinesGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseVirtualMachinesGetResponse_status(v string) (any, error) {
    result := OK_VIRTUALMACHINESGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_VIRTUALMACHINESGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_VIRTUALMACHINESGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_VIRTUALMACHINESGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_status(values []VirtualMachinesGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_status) isMultiValue() bool {
    return false
}
