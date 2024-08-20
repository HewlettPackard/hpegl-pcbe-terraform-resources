package item
// Resource Component Health Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_powerSuppliesHealth int

const (
    OK_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH WithServerGetResponse_health_powerSuppliesHealth = iota
    WARNING_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    CRITICAL_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    MISSING_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
)

func (i WithServerGetResponse_health_powerSuppliesHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_powerSuppliesHealth(v string) (any, error) {
    result := OK_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
    switch v {
        case "OK":
            result = OK_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "WARNING":
            result = WARNING_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "CRITICAL":
            result = CRITICAL_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "MISSING":
            result = MISSING_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_POWERSUPPLIESHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_powerSuppliesHealth(values []WithServerGetResponse_health_powerSuppliesHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_powerSuppliesHealth) isMultiValue() bool {
    return false
}
