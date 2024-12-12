package item
// new value of the "state" setting
type DualAuthOperationsPatchRequestBody_state int

const (
    APPROVED_DUALAUTHOPERATIONSPATCHREQUESTBODY_STATE DualAuthOperationsPatchRequestBody_state = iota
    CANCELLED_DUALAUTHOPERATIONSPATCHREQUESTBODY_STATE
)

func (i DualAuthOperationsPatchRequestBody_state) String() string {
    return []string{"APPROVED", "CANCELLED"}[i]
}
func ParseDualAuthOperationsPatchRequestBody_state(v string) (any, error) {
    result := APPROVED_DUALAUTHOPERATIONSPATCHREQUESTBODY_STATE
    switch v {
        case "APPROVED":
            result = APPROVED_DUALAUTHOPERATIONSPATCHREQUESTBODY_STATE
        case "CANCELLED":
            result = CANCELLED_DUALAUTHOPERATIONSPATCHREQUESTBODY_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDualAuthOperationsPatchRequestBody_state(values []DualAuthOperationsPatchRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DualAuthOperationsPatchRequestBody_state) isMultiValue() bool {
    return false
}
