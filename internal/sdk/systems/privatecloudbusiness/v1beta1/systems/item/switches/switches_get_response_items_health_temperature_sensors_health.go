package switches
// Resource Component Health Status. Indicated as ENUM with following mapping
type SwitchesGetResponse_items_health_temperatureSensorsHealth int

const (
    OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH SwitchesGetResponse_items_health_temperatureSensorsHealth = iota
    WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
    CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
    MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
    UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
)

func (i SwitchesGetResponse_items_health_temperatureSensorsHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseSwitchesGetResponse_items_health_temperatureSensorsHealth(v string) (any, error) {
    result := OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
    switch v {
        case "OK":
            result = OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
        case "WARNING":
            result = WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
        case "CRITICAL":
            result = CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
        case "MISSING":
            result = MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_TEMPERATURESENSORSHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSwitchesGetResponse_items_health_temperatureSensorsHealth(values []SwitchesGetResponse_items_health_temperatureSensorsHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SwitchesGetResponse_items_health_temperatureSensorsHealth) isMultiValue() bool {
    return false
}
