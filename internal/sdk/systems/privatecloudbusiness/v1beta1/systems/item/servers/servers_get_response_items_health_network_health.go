package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_networkHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH ServersGetResponse_items_health_networkHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
)

func (i ServersGetResponse_items_health_networkHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_networkHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_NETWORKHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_networkHealth(values []ServersGetResponse_items_health_networkHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_networkHealth) isMultiValue() bool {
    return false
}
