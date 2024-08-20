package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_smartStorageBatteryHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH ServersGetResponse_items_health_smartStorageBatteryHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
)

func (i ServersGetResponse_items_health_smartStorageBatteryHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_smartStorageBatteryHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_SMARTSTORAGEBATTERYHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_smartStorageBatteryHealth(values []ServersGetResponse_items_health_smartStorageBatteryHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_smartStorageBatteryHealth) isMultiValue() bool {
    return false
}
