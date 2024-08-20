package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_processorHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH ServersGetResponse_items_health_processorHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
)

func (i ServersGetResponse_items_health_processorHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_processorHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_PROCESSORHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_processorHealth(values []ServersGetResponse_items_health_processorHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_processorHealth) isMultiValue() bool {
    return false
}
