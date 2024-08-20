package servers
// Resource Component Health Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_powerSuppliesHealth int

const (
    OK_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH ServersGetResponse_items_health_powerSuppliesHealth = iota
    WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
)

func (i ServersGetResponse_items_health_powerSuppliesHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_powerSuppliesHealth(v string) (any, error) {
    result := OK_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
    switch v {
        case "OK":
            result = OK_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "WARNING":
            result = WARNING_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "CRITICAL":
            result = CRITICAL_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "MISSING":
            result = MISSING_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_POWERSUPPLIESHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_powerSuppliesHealth(values []ServersGetResponse_items_health_powerSuppliesHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_powerSuppliesHealth) isMultiValue() bool {
    return false
}
