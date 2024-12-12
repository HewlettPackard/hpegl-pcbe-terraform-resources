package systems
type SystemsGetResponse_items_health_servers int

const (
    OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS SystemsGetResponse_items_health_servers = iota
    WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
    CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
    MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
)

func (i SystemsGetResponse_items_health_servers) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseSystemsGetResponse_items_health_servers(v string) (any, error) {
    result := OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
    switch v {
        case "OK":
            result = OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
        case "WARNING":
            result = WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
        case "CRITICAL":
            result = CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
        case "MISSING":
            result = MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_SERVERS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_health_servers(values []SystemsGetResponse_items_health_servers) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_health_servers) isMultiValue() bool {
    return false
}
