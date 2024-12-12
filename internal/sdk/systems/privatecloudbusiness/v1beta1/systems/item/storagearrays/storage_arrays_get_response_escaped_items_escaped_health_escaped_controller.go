package storagearrays
type StorageArraysGetResponse_items_health_controller int

const (
    OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER StorageArraysGetResponse_items_health_controller = iota
    WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
    CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
    MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
)

func (i StorageArraysGetResponse_items_health_controller) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseStorageArraysGetResponse_items_health_controller(v string) (any, error) {
    result := OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
    switch v {
        case "OK":
            result = OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
        case "WARNING":
            result = WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
        case "CRITICAL":
            result = CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
        case "MISSING":
            result = MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_CONTROLLER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageArraysGetResponse_items_health_controller(values []StorageArraysGetResponse_items_health_controller) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageArraysGetResponse_items_health_controller) isMultiValue() bool {
    return false
}
