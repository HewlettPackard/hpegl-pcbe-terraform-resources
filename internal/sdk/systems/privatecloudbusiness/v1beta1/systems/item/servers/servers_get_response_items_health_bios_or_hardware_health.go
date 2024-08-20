package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_biosOrHardwareHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH ServersGetResponse_items_health_biosOrHardwareHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
)

func (i ServersGetResponse_items_health_biosOrHardwareHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_biosOrHardwareHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_BIOSORHARDWAREHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_biosOrHardwareHealth(values []ServersGetResponse_items_health_biosOrHardwareHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_biosOrHardwareHealth) isMultiValue() bool {
    return false
}
