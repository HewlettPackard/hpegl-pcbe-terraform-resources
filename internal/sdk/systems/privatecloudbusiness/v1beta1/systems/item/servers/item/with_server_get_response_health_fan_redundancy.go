package item
// Resource Component Redundancy Status. Indicated as ENUM with following mapping
type WithServerGetResponse_health_fanRedundancy int

const (
    REDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY WithServerGetResponse_health_fanRedundancy = iota
    NONREDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    FAILEDREDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    ABSENT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    DISABLED_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    OFFLINE_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    UNAVAILABLEOFFLINE_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    STANDBYOFFLINE_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
)

func (i WithServerGetResponse_health_fanRedundancy) String() string {
    return []string{"REDUNDANT", "NONREDUNDANT", "FAILEDREDUNDANT", "ABSENT", "DISABLED", "OFFLINE", "UNAVAILABLEOFFLINE", "STANDBYOFFLINE", "UNKNOWN"}[i]
}
func ParseWithServerGetResponse_health_fanRedundancy(v string) (any, error) {
    result := REDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
    switch v {
        case "REDUNDANT":
            result = REDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "NONREDUNDANT":
            result = NONREDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "FAILEDREDUNDANT":
            result = FAILEDREDUNDANT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "ABSENT":
            result = ABSENT_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "DISABLED":
            result = DISABLED_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "OFFLINE":
            result = OFFLINE_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "UNAVAILABLEOFFLINE":
            result = UNAVAILABLEOFFLINE_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "STANDBYOFFLINE":
            result = STANDBYOFFLINE_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        case "UNKNOWN":
            result = UNKNOWN_WITHSERVERGETRESPONSE_HEALTH_FANREDUNDANCY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_health_fanRedundancy(values []WithServerGetResponse_health_fanRedundancy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_health_fanRedundancy) isMultiValue() bool {
    return false
}
