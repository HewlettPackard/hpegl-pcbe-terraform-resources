package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithSwitchGetResponse_health_cpuHealth int

const (
    OK_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH WithSwitchGetResponse_health_cpuHealth = iota
    WARNING_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
    CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
    MISSING_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
    UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
)

func (i WithSwitchGetResponse_health_cpuHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithSwitchGetResponse_health_cpuHealth(v string) (any, error) {
    result := OK_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
    switch v {
        case "OK":
            result = OK_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
        case "WARNING":
            result = WARNING_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
        case "MISSING":
            result = MISSING_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_CPUHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSwitchGetResponse_health_cpuHealth(values []WithSwitchGetResponse_health_cpuHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSwitchGetResponse_health_cpuHealth) isMultiValue() bool {
    return false
}
