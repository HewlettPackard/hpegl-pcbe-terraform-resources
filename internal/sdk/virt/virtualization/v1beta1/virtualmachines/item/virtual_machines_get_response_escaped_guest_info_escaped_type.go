package item
// Operating system on this virtual machine.
type VirtualMachinesGetResponse_guestInfo_type int

const (
    WINDOWS_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE VirtualMachinesGetResponse_guestInfo_type = iota
    LINUX_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE
    OTHERS_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE
)

func (i VirtualMachinesGetResponse_guestInfo_type) String() string {
    return []string{"WINDOWS", "LINUX", "OTHERS"}[i]
}
func ParseVirtualMachinesGetResponse_guestInfo_type(v string) (any, error) {
    result := WINDOWS_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE
    switch v {
        case "WINDOWS":
            result = WINDOWS_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE
        case "LINUX":
            result = LINUX_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE
        case "OTHERS":
            result = OTHERS_VIRTUALMACHINESGETRESPONSE_GUESTINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_guestInfo_type(values []VirtualMachinesGetResponse_guestInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_guestInfo_type) isMultiValue() bool {
    return false
}
