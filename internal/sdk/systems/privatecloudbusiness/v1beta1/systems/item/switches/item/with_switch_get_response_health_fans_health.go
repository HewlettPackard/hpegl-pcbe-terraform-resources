package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithSwitchGetResponse_health_fansHealth int

const (
    OK_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH WithSwitchGetResponse_health_fansHealth = iota
    WARNING_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
    CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
    MISSING_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
    UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
)

func (i WithSwitchGetResponse_health_fansHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithSwitchGetResponse_health_fansHealth(v string) (any, error) {
    result := OK_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
    switch v {
        case "OK":
            result = OK_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
        case "WARNING":
            result = WARNING_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
        case "MISSING":
            result = MISSING_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_FANSHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSwitchGetResponse_health_fansHealth(values []WithSwitchGetResponse_health_fansHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSwitchGetResponse_health_fansHealth) isMultiValue() bool {
    return false
}
