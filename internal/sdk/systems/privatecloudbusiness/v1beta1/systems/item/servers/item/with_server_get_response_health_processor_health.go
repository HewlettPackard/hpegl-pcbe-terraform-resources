package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_processorHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH WithServerGetResponse_health_processorHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
)

func (i WithServerGetResponse_health_processorHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_processorHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_PROCESSORHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_processorHealth(values []WithServerGetResponse_health_processorHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_processorHealth) isMultiValue() bool {
    return false
}
