package switches
// Resource Component Health Status. Indicated as ENUM with following mapping
type SwitchesGetResponse_items_health_fansHealth int

const (
    OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH SwitchesGetResponse_items_health_fansHealth = iota
    WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
    CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
    MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
    UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
)

func (i SwitchesGetResponse_items_health_fansHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseSwitchesGetResponse_items_health_fansHealth(v string) (any, error) {
    result := OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
    switch v {
        case "OK":
            result = OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
        case "WARNING":
            result = WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
        case "CRITICAL":
            result = CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
        case "MISSING":
            result = MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_FANSHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSwitchesGetResponse_items_health_fansHealth(values []SwitchesGetResponse_items_health_fansHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SwitchesGetResponse_items_health_fansHealth) isMultiValue() bool {
    return false
}
