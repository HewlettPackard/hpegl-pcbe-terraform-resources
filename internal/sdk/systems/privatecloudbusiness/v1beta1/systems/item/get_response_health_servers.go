package item
type GetResponse_health_servers int

const (
    OK_GETRESPONSE_HEALTH_SERVERS GetResponse_health_servers = iota
    WARNING_GETRESPONSE_HEALTH_SERVERS
    CRITICAL_GETRESPONSE_HEALTH_SERVERS
    MISSING_GETRESPONSE_HEALTH_SERVERS
)

func (i GetResponse_health_servers) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseGetResponse_health_servers(v string) (any, error) {
    result := OK_GETRESPONSE_HEALTH_SERVERS
    switch v {
        case "OK":
            result = OK_GETRESPONSE_HEALTH_SERVERS
        case "WARNING":
            result = WARNING_GETRESPONSE_HEALTH_SERVERS
        case "CRITICAL":
            result = CRITICAL_GETRESPONSE_HEALTH_SERVERS
        case "MISSING":
            result = MISSING_GETRESPONSE_HEALTH_SERVERS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_health_servers(values []GetResponse_health_servers) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_health_servers) isMultiValue() bool {
    return false
}
