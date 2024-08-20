package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_storageHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH WithServerGetResponse_health_storageHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
)

func (i WithServerGetResponse_health_storageHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_storageHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_STORAGEHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_storageHealth(values []WithServerGetResponse_health_storageHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_storageHealth) isMultiValue() bool {
    return false
}
