package storagearrays
type StorageArraysGetResponse_items_health_shelf int

const (
    OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF StorageArraysGetResponse_items_health_shelf = iota
    WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
    CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
    MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
)

func (i StorageArraysGetResponse_items_health_shelf) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseStorageArraysGetResponse_items_health_shelf(v string) (any, error) {
    result := OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
    switch v {
        case "OK":
            result = OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
        case "WARNING":
            result = WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
        case "CRITICAL":
            result = CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
        case "MISSING":
            result = MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_SHELF
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageArraysGetResponse_items_health_shelf(values []StorageArraysGetResponse_items_health_shelf) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageArraysGetResponse_items_health_shelf) isMultiValue() bool {
    return false
}
