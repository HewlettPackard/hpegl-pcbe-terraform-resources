package systems
type SystemsGetResponse_items_health_storage int

const (
    OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE SystemsGetResponse_items_health_storage = iota
    WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
    CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
    MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
)

func (i SystemsGetResponse_items_health_storage) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseSystemsGetResponse_items_health_storage(v string) (any, error) {
    result := OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
    switch v {
        case "OK":
            result = OK_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
        case "WARNING":
            result = WARNING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
        case "CRITICAL":
            result = CRITICAL_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
        case "MISSING":
            result = MISSING_SYSTEMSGETRESPONSE_ITEMS_HEALTH_STORAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_health_storage(values []SystemsGetResponse_items_health_storage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_health_storage) isMultiValue() bool {
    return false
}
