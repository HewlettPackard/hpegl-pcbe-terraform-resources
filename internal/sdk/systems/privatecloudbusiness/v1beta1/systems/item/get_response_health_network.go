package item
type GetResponse_health_network int

const (
    OK_GETRESPONSE_HEALTH_NETWORK GetResponse_health_network = iota
    WARNING_GETRESPONSE_HEALTH_NETWORK
    CRITICAL_GETRESPONSE_HEALTH_NETWORK
    MISSING_GETRESPONSE_HEALTH_NETWORK
)

func (i GetResponse_health_network) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseGetResponse_health_network(v string) (any, error) {
    result := OK_GETRESPONSE_HEALTH_NETWORK
    switch v {
        case "OK":
            result = OK_GETRESPONSE_HEALTH_NETWORK
        case "WARNING":
            result = WARNING_GETRESPONSE_HEALTH_NETWORK
        case "CRITICAL":
            result = CRITICAL_GETRESPONSE_HEALTH_NETWORK
        case "MISSING":
            result = MISSING_GETRESPONSE_HEALTH_NETWORK
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_health_network(values []GetResponse_health_network) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_health_network) isMultiValue() bool {
    return false
}
