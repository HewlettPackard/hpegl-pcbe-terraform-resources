package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithSwitchGetResponse_health_powerSuppliesHealth int

const (
    OK_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH WithSwitchGetResponse_health_powerSuppliesHealth = iota
    WARNING_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    MISSING_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
)

func (i WithSwitchGetResponse_health_powerSuppliesHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithSwitchGetResponse_health_powerSuppliesHealth(v string) (any, error) {
    result := OK_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    switch v {
        case "OK":
            result = OK_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "WARNING":
            result = WARNING_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "MISSING":
            result = MISSING_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSWITCHGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSwitchGetResponse_health_powerSuppliesHealth(values []WithSwitchGetResponse_health_powerSuppliesHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSwitchGetResponse_health_powerSuppliesHealth) isMultiValue() bool {
    return false
}
