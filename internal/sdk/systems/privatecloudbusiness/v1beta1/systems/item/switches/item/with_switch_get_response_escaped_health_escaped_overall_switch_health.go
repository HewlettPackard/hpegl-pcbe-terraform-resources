package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithSwitchGetResponse_health_overallSwitchHealth int

const (
    OK_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH WithSwitchGetResponse_health_overallSwitchHealth = iota
    WARNING_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
    CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
    MISSING_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
    UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
)

func (i WithSwitchGetResponse_health_overallSwitchHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithSwitchGetResponse_health_overallSwitchHealth(v string) (any, error) {
    result := OK_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
    switch v {
        case "OK":
            result = OK_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
        case "WARNING":
            result = WARNING_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
        case "MISSING":
            result = MISSING_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_OVERALLSWITCHHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSwitchGetResponse_health_overallSwitchHealth(values []WithSwitchGetResponse_health_overallSwitchHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSwitchGetResponse_health_overallSwitchHealth) isMultiValue() bool {
    return false
}
