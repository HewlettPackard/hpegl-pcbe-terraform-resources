package switches
// Resource Component Health Status. Indicated as ENUM with following mapping
type SwitchesGetResponse_items_health_overallSwitchHealth int

const (
    OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH SwitchesGetResponse_items_health_overallSwitchHealth = iota
    WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
    CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
    MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
    UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
)

func (i SwitchesGetResponse_items_health_overallSwitchHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseSwitchesGetResponse_items_health_overallSwitchHealth(v string) (any, error) {
    result := OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
    switch v {
        case "OK":
            result = OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
        case "WARNING":
            result = WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
        case "CRITICAL":
            result = CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
        case "MISSING":
            result = MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_OVERALLSWITCHHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSwitchesGetResponse_items_health_overallSwitchHealth(values []SwitchesGetResponse_items_health_overallSwitchHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SwitchesGetResponse_items_health_overallSwitchHealth) isMultiValue() bool {
    return false
}