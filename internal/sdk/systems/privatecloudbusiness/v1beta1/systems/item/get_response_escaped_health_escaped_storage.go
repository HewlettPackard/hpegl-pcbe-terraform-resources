package item
type GetResponse_health_storage int

const (
    OK_GETRESPONSE_HEALTH_STORAGE GetResponse_health_storage = iota
    WARNING_GETRESPONSE_HEALTH_STORAGE
    CRITICAL_GETRESPONSE_HEALTH_STORAGE
    MISSING_GETRESPONSE_HEALTH_STORAGE
)

func (i GetResponse_health_storage) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseGetResponse_health_storage(v string) (any, error) {
    result := OK_GETRESPONSE_HEALTH_STORAGE
    switch v {
        case "OK":
            result = OK_GETRESPONSE_HEALTH_STORAGE
        case "WARNING":
            result = WARNING_GETRESPONSE_HEALTH_STORAGE
        case "CRITICAL":
            result = CRITICAL_GETRESPONSE_HEALTH_STORAGE
        case "MISSING":
            result = MISSING_GETRESPONSE_HEALTH_STORAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_health_storage(values []GetResponse_health_storage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_health_storage) isMultiValue() bool {
    return false
}
