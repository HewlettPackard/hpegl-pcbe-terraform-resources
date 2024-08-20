package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_overallServerHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH ServersGetResponse_items_health_overallServerHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
)

func (i ServersGetResponse_items_health_overallServerHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_overallServerHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_OVERALLSERVERHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_overallServerHealth(values []ServersGetResponse_items_health_overallServerHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_overallServerHealth) isMultiValue() bool {
    return false
}
