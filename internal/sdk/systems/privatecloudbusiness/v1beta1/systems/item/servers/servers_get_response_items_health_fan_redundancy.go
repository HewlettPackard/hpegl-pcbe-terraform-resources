package servers
// Resource Component Redundancy Status. Indicated as ENUM with following mapping
type ServersGetResponse_items_health_fanRedundancy int

const (
    REDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY ServersGetResponse_items_health_fanRedundancy = iota
    NONREDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    FAILEDREDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    ABSENT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    DISABLED_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    OFFLINE_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    UNAVAILABLEOFFLINE_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    STANDBYOFFLINE_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
)

func (i ServersGetResponse_items_health_fanRedundancy) String() string {
    return []string{"REDUNDANT", "NONREDUNDANT", "FAILEDREDUNDANT", "ABSENT", "DISABLED", "OFFLINE", "UNAVAILABLEOFFLINE", "STANDBYOFFLINE", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_health_fanRedundancy(v string) (any, error) {
    result := REDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
    switch v {
        case "REDUNDANT":
            result = REDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "NONREDUNDANT":
            result = NONREDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "FAILEDREDUNDANT":
            result = FAILEDREDUNDANT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "ABSENT":
            result = ABSENT_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "DISABLED":
            result = DISABLED_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "OFFLINE":
            result = OFFLINE_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "UNAVAILABLEOFFLINE":
            result = UNAVAILABLEOFFLINE_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "STANDBYOFFLINE":
            result = STANDBYOFFLINE_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_HEALTH_FANREDUNDANCY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_health_fanRedundancy(values []ServersGetResponse_items_health_fanRedundancy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_health_fanRedundancy) isMultiValue() bool {
    return false
}
