package updatehardware
// Type of the network adapter. Mandatory parameter for ADD operation.
type UpdateHardwarePostRequestBody_networkAdapters_type int

const (
    E1000_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE UpdateHardwarePostRequestBody_networkAdapters_type = iota
    E1000E_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
    PCNET32_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
    VMXNET_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
    VMXNET2_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
    VMXNET3_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
)

func (i UpdateHardwarePostRequestBody_networkAdapters_type) String() string {
    return []string{"E1000", "E1000E", "PCNET32", "VMXNET", "VMXNET2", "VMXNET3"}[i]
}
func ParseUpdateHardwarePostRequestBody_networkAdapters_type(v string) (any, error) {
    result := E1000_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
    switch v {
        case "E1000":
            result = E1000_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
        case "E1000E":
            result = E1000E_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
        case "PCNET32":
            result = PCNET32_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
        case "VMXNET":
            result = VMXNET_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
        case "VMXNET2":
            result = VMXNET2_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
        case "VMXNET3":
            result = VMXNET3_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUpdateHardwarePostRequestBody_networkAdapters_type(values []UpdateHardwarePostRequestBody_networkAdapters_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UpdateHardwarePostRequestBody_networkAdapters_type) isMultiValue() bool {
    return false
}
