package systems
type SystemsGetResponse_items_health_network int

const (
    OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK SystemsGetResponse_items_health_network = iota
    WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
    CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
    MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
)

func (i SystemsGetResponse_items_health_network) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseSystemsGetResponse_items_health_network(v string) (any, error) {
    result := OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
    switch v {
        case "OK":
            result = OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
        case "WARNING":
            result = WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
        case "CRITICAL":
            result = CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
        case "MISSING":
            result = MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_NETWORK
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_health_network(values []SystemsGetResponse_items_health_network) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_health_network) isMultiValue() bool {
    return false
}
