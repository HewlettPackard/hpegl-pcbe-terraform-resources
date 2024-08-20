package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_temperaturesHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH WithServerGetResponse_health_temperaturesHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
)

func (i WithServerGetResponse_health_temperaturesHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_temperaturesHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_TEMPERATURESHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_temperaturesHealth(values []WithServerGetResponse_health_temperaturesHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_temperaturesHealth) isMultiValue() bool {
    return false
}
