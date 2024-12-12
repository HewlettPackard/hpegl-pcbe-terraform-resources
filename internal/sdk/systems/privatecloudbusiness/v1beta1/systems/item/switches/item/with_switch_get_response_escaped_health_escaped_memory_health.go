package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithSwitchGetResponse_health_memoryHealth int

const (
    OK_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH WithSwitchGetResponse_health_memoryHealth = iota
    WARNING_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
    CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
    MISSING_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
    UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
)

func (i WithSwitchGetResponse_health_memoryHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithSwitchGetResponse_health_memoryHealth(v string) (any, error) {
    result := OK_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
    switch v {
        case "OK":
            result = OK_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
        case "WARNING":
            result = WARNING_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
        case "MISSING":
            result = MISSING_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_MEMORYHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSwitchGetResponse_health_memoryHealth(values []WithSwitchGetResponse_health_memoryHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSwitchGetResponse_health_memoryHealth) isMultiValue() bool {
    return false
}
