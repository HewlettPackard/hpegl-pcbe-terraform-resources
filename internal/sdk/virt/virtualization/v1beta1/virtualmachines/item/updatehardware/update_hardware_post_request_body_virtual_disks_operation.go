package updatehardware
// Operation to be performed on a disk. Minimum size required to  add a disk is 1MB. Maximum capacity for a disk depends on the backing storage.
type UpdateHardwarePostRequestBody_virtualDisks_operation int

const (
    ADD_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION UpdateHardwarePostRequestBody_virtualDisks_operation = iota
    EDIT_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION
    DELETE_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION
)

func (i UpdateHardwarePostRequestBody_virtualDisks_operation) String() string {
    return []string{"ADD", "EDIT", "DELETE"}[i]
}
func ParseUpdateHardwarePostRequestBody_virtualDisks_operation(v string) (any, error) {
    result := ADD_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION
    switch v {
        case "ADD":
            result = ADD_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION
        case "EDIT":
            result = EDIT_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION
        case "DELETE":
            result = DELETE_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_OPERATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUpdateHardwarePostRequestBody_virtualDisks_operation(values []UpdateHardwarePostRequestBody_virtualDisks_operation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UpdateHardwarePostRequestBody_virtualDisks_operation) isMultiValue() bool {
    return false
}
