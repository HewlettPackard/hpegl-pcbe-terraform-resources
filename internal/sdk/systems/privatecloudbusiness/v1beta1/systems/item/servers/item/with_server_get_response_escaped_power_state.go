package item
// Server Power state.
type WithServerGetResponse_powerState int

const (
    ON_WITHSERVERGETRESPONSE_POWERSTATE WithServerGetResponse_powerState = iota
    RESET_WITHSERVERGETRESPONSE_POWERSTATE
    UNKNOWN_WITHSERVERGETRESPONSE_POWERSTATE
)

func (i WithServerGetResponse_powerState) String() string {
    return []string{"ON", "RESET", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_powerState(v string) (any, error) {
    result := ON_WITHSERVERGETRESPONSE_POWERSTATE
    switch v {
        case "ON":
            result = ON_WITHSERVERGETRESPONSE_POWERSTATE
        case "RESET":
            result = RESET_WITHSERVERGETRESPONSE_POWERSTATE
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_POWERSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_powerState(values []WithServerGetResponse_powerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_powerState) isMultiValue() bool {
    return false
}
