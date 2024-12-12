package systems
type SystemsGetResponse_items_health_overallHealth int

const (
    OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH SystemsGetResponse_items_health_overallHealth = iota
    WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
    CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
    MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
)

func (i SystemsGetResponse_items_health_overallHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseSystemsGetResponse_items_health_overallHealth(v string) (any, error) {
    result := OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
    switch v {
        case "OK":
            result = OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        case "WARNING":
            result = WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        case "CRITICAL":
            result = CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        case "MISSING":
            result = MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_health_overallHealth(values []SystemsGetResponse_items_health_overallHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_health_overallHealth) isMultiValue() bool {
    return false
}
