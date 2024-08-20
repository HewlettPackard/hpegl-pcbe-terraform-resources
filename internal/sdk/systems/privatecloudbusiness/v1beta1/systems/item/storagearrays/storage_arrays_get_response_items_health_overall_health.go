package storagearrays
type StorageArraysGetResponse_items_health_overallHealth int

const (
    OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH StorageArraysGetResponse_items_health_overallHealth = iota
    WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
    CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
    MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
)

func (i StorageArraysGetResponse_items_health_overallHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParseStorageArraysGetResponse_items_health_overallHealth(v string) (any, error) {
    result := OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
    switch v {
        case "OK":
            result = OK_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        case "WARNING":
            result = WARNING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        case "CRITICAL":
            result = CRITICAL_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        case "MISSING":
            result = MISSING_STORAGEARRAYSGETRESPONSE_ITEMS_HEALTH_OVERALLHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageArraysGetResponse_items_health_overallHealth(values []StorageArraysGetResponse_items_health_overallHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageArraysGetResponse_items_health_overallHealth) isMultiValue() bool {
    return false
}
