package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_smartStorageBatteryHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH WithServerGetResponse_health_smartStorageBatteryHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
)

func (i WithServerGetResponse_health_smartStorageBatteryHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_smartStorageBatteryHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_SMARTSTORAGEBATTERYHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_smartStorageBatteryHealth(values []WithServerGetResponse_health_smartStorageBatteryHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_smartStorageBatteryHealth) isMultiValue() bool {
    return false
}
