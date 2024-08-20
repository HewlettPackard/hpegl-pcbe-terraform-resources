package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_fanHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH ServersGetResponse_items_health_fanHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
)

func (i ServersGetResponse_items_health_fanHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_fanHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_FANHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_fanHealth(values []ServersGetResponse_items_health_fanHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_fanHealth) isMultiValue() bool {
    return false
}
