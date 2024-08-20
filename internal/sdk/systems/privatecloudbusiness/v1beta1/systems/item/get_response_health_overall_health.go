package item
type GetResponse_health_overallHealth int

const (
    OK_GETRESPONSE_HEALTH_OVERALLHEALTH GetResponse_health_overallHealth = iota
    WARNING_GETRESPONSE_HEALTH_OVERALLHEALTH
    CRITICAL_GETRESPONSE_HEALTH_OVERALLHEALTH
    MISSING_GETRESPONSE_HEALTH_OVERALLHEALTH
)

func (i GetResponse_health_overallHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseGetResponse_health_overallHealth(v string) (any, error) {
    result := OK_GETRESPONSE_HEALTH_OVERALLHEALTH
    switch v {
        case "OK":
            result = OK_GETRESPONSE_HEALTH_OVERALLHEALTH
        case "WARNING":
            result = WARNING_GETRESPONSE_HEALTH_OVERALLHEALTH
        case "CRITICAL":
            result = CRITICAL_GETRESPONSE_HEALTH_OVERALLHEALTH
        case "MISSING":
            result = MISSING_GETRESPONSE_HEALTH_OVERALLHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_health_overallHealth(values []GetResponse_health_overallHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_health_overallHealth) isMultiValue() bool {
    return false
}
