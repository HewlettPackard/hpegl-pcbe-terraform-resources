package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_memoryHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH WithServerGetResponse_health_memoryHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
)

func (i WithServerGetResponse_health_memoryHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_memoryHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_MEMORYHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_memoryHealth(values []WithServerGetResponse_health_memoryHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_memoryHealth) isMultiValue() bool {
    return false
}
