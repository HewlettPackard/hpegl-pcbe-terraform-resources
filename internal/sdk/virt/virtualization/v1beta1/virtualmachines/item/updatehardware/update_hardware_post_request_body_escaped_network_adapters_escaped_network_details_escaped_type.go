package updatehardware
// Type of the backing network.
type UpdateHardwarePostRequestBody_networkAdapters_networkDetails_type int

const (
    STANDARD_PORT_GROUP_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_NETWORKDETAILS_TYPE UpdateHardwarePostRequestBody_networkAdapters_networkDetails_type = iota
)

func (i UpdateHardwarePostRequestBody_networkAdapters_networkDetails_type) String() string {
    return []string{"STANDARD_PORT_GROUP"}[i]
}
func ParseUpdateHardwarePostRequestBody_networkAdapters_networkDetails_type(v string) (any, error) {
    result := STANDARD_PORT_GROUP_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_NETWORKDETAILS_TYPE
    switch v {
        case "STANDARD_PORT_GROUP":
            result = STANDARD_PORT_GROUP_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_NETWORKDETAILS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUpdateHardwarePostRequestBody_networkAdapters_networkDetails_type(values []UpdateHardwarePostRequestBody_networkAdapters_networkDetails_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UpdateHardwarePostRequestBody_networkAdapters_networkDetails_type) isMultiValue() bool {
    return false
}
