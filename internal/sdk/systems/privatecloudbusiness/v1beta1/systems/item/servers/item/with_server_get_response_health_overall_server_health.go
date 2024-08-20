package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_overallServerHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH WithServerGetResponse_health_overallServerHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
)

func (i WithServerGetResponse_health_overallServerHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_overallServerHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_OVERALLSERVERHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_overallServerHealth(values []WithServerGetResponse_health_overallServerHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_overallServerHealth) isMultiValue() bool {
    return false
}
