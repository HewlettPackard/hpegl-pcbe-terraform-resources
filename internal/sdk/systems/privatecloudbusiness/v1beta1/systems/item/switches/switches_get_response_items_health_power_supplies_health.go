package switches
// Resource Component Health Status. Indicated as ENUM with following mapping
type SwitchesGetResponse_items_health_powerSuppliesHealth int

const (
    OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH SwitchesGetResponse_items_health_powerSuppliesHealth = iota
    WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
)

func (i SwitchesGetResponse_items_health_powerSuppliesHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseSwitchesGetResponse_items_health_powerSuppliesHealth(v string) (any, error) {
    result := OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    switch v {
        case "OK":
            result = OK_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "WARNING":
            result = WARNING_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "CRITICAL":
            result = CRITICAL_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "MISSING":
            result = MISSING_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SWITCHESGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSwitchesGetResponse_items_health_powerSuppliesHealth(values []SwitchesGetResponse_items_health_powerSuppliesHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SwitchesGetResponse_items_health_powerSuppliesHealth) isMultiValue() bool {
    return false
}
