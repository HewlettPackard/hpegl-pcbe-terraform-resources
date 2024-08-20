package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_temperaturesHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH ServersGetResponse_items_health_temperaturesHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
)

func (i ServersGetResponse_items_health_temperaturesHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_temperaturesHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_TEMPERATURESHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_temperaturesHealth(values []ServersGetResponse_items_health_temperaturesHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_temperaturesHealth) isMultiValue() bool {
    return false
}
