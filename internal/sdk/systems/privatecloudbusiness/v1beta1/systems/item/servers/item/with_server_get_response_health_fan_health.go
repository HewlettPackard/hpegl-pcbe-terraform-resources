package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_fanHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH WithServerGetResponse_health_fanHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
)

func (i WithServerGetResponse_health_fanHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_fanHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_FANHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_fanHealth(values []WithServerGetResponse_health_fanHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_fanHealth) isMultiValue() bool {
    return false
}
