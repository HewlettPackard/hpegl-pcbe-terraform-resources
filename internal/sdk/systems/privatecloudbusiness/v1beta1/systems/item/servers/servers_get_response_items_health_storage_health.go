package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_storageHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH ServersGetResponse_items_health_storageHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
)

func (i ServersGetResponse_items_health_storageHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_storageHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_STORAGEHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_storageHealth(values []ServersGetResponse_items_health_storageHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_storageHealth) isMultiValue() bool {
    return false
}
