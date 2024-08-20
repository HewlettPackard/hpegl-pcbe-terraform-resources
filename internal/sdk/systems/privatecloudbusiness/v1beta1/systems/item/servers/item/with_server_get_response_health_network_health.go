package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_networkHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH WithServerGetResponse_health_networkHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
)

func (i WithServerGetResponse_health_networkHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_networkHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_NETWORKHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_networkHealth(values []WithServerGetResponse_health_networkHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_networkHealth) isMultiValue() bool {
    return false
}
