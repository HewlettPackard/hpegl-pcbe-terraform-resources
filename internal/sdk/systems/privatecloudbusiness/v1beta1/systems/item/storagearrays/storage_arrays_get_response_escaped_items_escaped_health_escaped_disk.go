package storagearrays
type StorageArraysGetResponse_items_health_disk int

const (
    OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK StorageArraysGetResponse_items_health_disk = iota
    WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
    CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
    MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
)

func (i StorageArraysGetResponse_items_health_disk) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseStorageArraysGetResponse_items_health_disk(v string) (any, error) {
    result := OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
    switch v {
        case "OK":
            result = OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
        case "WARNING":
            result = WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
        case "CRITICAL":
            result = CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
        case "MISSING":
            result = MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_DISK
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageArraysGetResponse_items_health_disk(values []StorageArraysGetResponse_items_health_disk) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageArraysGetResponse_items_health_disk) isMultiValue() bool {
    return false
}
