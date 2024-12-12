package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithSwitchGetResponse_health_temperatureSensorsHealth int

const (
    OK_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH WithSwitchGetResponse_health_temperatureSensorsHealth = iota
    WARNING_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
    CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
    MISSING_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
    UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
)

func (i WithSwitchGetResponse_health_temperatureSensorsHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithSwitchGetResponse_health_temperatureSensorsHealth(v string) (any, error) {
    result := OK_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
    switch v {
        case "OK":
            result = OK_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
        case "WARNING":
            result = WARNING_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
        case "MISSING":
            result = MISSING_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_TEMPERATURESENSORSHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSwitchGetResponse_health_temperatureSensorsHealth(values []WithSwitchGetResponse_health_temperatureSensorsHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSwitchGetResponse_health_temperatureSensorsHealth) isMultiValue() bool {
    return false
}
