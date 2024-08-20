package storagearrays
type StorageArraysGetResponse_items_health_network int

const (
    OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK StorageArraysGetResponse_items_health_network = iota
    WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
    CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
    MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
)

func (i StorageArraysGetResponse_items_health_network) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseStorageArraysGetResponse_items_health_network(v string) (any, error) {
    result := OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
    switch v {
        case "OK":
            result = OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
        case "WARNING":
            result = WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
        case "CRITICAL":
            result = CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
        case "MISSING":
            result = MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_NETWORK
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageArraysGetResponse_items_health_network(values []StorageArraysGetResponse_items_health_network) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageArraysGetResponse_items_health_network) isMultiValue() bool {
    return false
}
