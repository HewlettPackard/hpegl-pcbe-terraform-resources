package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_memoryHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH ServersGetResponse_items_health_memoryHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
)

func (i ServersGetResponse_items_health_memoryHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_memoryHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_memoryHealth(values []ServersGetResponse_items_health_memoryHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_memoryHealth) isMultiValue() bool {
    return false
}
