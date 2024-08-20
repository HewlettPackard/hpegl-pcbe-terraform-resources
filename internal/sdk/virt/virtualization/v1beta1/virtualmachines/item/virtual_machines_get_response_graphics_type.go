package item
// Type of graphics adapter.
type VirtualMachinesGetResponse_graphics_type int

const (
    VNC_VIRTUALMACHINESGETRESPONSE_GRAPHICS_TYPE VirtualMachinesGetResponse_graphics_type = iota
    SPICE_VIRTUALMACHINESGETRESPONSE_GRAPHICS_TYPE
)

func (i VirtualMachinesGetResponse_graphics_type) String() string {
    return []string{"VNC", "SPICE"}[i]
}
func ParseVirtualMachinesGetResponse_graphics_type(v string) (any, error) {
    result := VNC_VIRTUALMACHINESGETRESPONSE_GRAPHICS_TYPE
    switch v {
        case "VNC":
            result = VNC_VIRTUALMACHINESGETRESPONSE_GRAPHICS_TYPE
        case "SPICE":
            result = SPICE_VIRTUALMACHINESGETRESPONSE_GRAPHICS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_graphics_type(values []VirtualMachinesGetResponse_graphics_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_graphics_type) isMultiValue() bool {
    return false
}
