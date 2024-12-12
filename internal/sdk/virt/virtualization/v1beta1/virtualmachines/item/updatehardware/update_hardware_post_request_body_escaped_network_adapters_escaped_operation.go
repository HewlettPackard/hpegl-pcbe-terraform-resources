package updatehardware
// Operation to be performed on a network adapter.
type UpdateHardwarePostRequestBody_networkAdapters_operation int

const (
    ADD_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION UpdateHardwarePostRequestBody_networkAdapters_operation = iota
    EDIT_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION
    DELETE_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION
)

func (i UpdateHardwarePostRequestBody_networkAdapters_operation) String() string {
    return []string{"ADD", "EDIT", "DELETE"}[i]
}
func ParseUpdateHardwarePostRequestBody_networkAdapters_operation(v string) (any, error) {
    result := ADD_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION
    switch v {
        case "ADD":
            result = ADD_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION
        case "EDIT":
            result = EDIT_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION
        case "DELETE":
            result = DELETE_UPDATEHARDWAREPOSTREQUESTBODY_NETWORKADAPTERS_OPERATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUpdateHardwarePostRequestBody_networkAdapters_operation(values []UpdateHardwarePostRequestBody_networkAdapters_operation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UpdateHardwarePostRequestBody_networkAdapters_operation) isMultiValue() bool {
    return false
}
