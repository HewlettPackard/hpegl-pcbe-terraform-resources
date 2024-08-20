package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_biosOrHardwareHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH WithServerGetResponse_health_biosOrHardwareHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
)

func (i WithServerGetResponse_health_biosOrHardwareHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_biosOrHardwareHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_BIOSORHARDWAREHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_biosOrHardwareHealth(values []WithServerGetResponse_health_biosOrHardwareHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_biosOrHardwareHealth) isMultiValue() bool {
    return false
}
