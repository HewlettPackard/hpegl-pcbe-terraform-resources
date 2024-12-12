package switches
// Resource Component Health Status. Indicated as ENUM with following mapping
type SwitchesGetResponse_items_health_cpuHealth int

const (
    OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH SwitchesGetResponse_items_health_cpuHealth = iota
    WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
    CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
    MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
    UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
)

func (i SwitchesGetResponse_items_health_cpuHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseSwitchesGetResponse_items_health_cpuHealth(v string) (any, error) {
    result := OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
    switch v {
        case "OK":
            result = OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
        case "WARNING":
            result = WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
        case "CRITICAL":
            result = CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
        case "MISSING":
            result = MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_CPUHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSwitchesGetResponse_items_health_cpuHealth(values []SwitchesGetResponse_items_health_cpuHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SwitchesGetResponse_items_health_cpuHealth) isMultiValue() bool {
    return false
}
