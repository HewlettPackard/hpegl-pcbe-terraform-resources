package switches
// Resource Component Health Status. Indicated as ENUM with following mapping
type SwitchesGetResponse_items_health_memoryHealth int

const (
    OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH SwitchesGetResponse_items_health_memoryHealth = iota
    WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
)

func (i SwitchesGetResponse_items_health_memoryHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseSwitchesGetResponse_items_health_memoryHealth(v string) (any, error) {
    result := OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
    switch v {
        case "OK":
            result = OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "WARNING":
            result = WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "CRITICAL":
            result = CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "MISSING":
            result = MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_MEMORYHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSwitchesGetResponse_items_health_memoryHealth(values []SwitchesGetResponse_items_health_memoryHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SwitchesGetResponse_items_health_memoryHealth) isMultiValue() bool {
    return false
}
