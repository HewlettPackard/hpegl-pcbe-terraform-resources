package storagelocations
// The Cloud Service Provider (CSP) of the location.
type StorageLocationsGetResponse_items_cloudServiceProvider int

const (
    AWS_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER StorageLocationsGetResponse_items_cloudServiceProvider = iota
    AZURE_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER
    GCP_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER
)

func (i StorageLocationsGetResponse_items_cloudServiceProvider) String() string {
    return []string{"AWS", "AZURE", "GCP"}[i]
}
func ParseStorageLocationsGetResponse_items_cloudServiceProvider(v string) (any, error) {
    result := AWS_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER
    switch v {
        case "AWS":
            result = AWS_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER
        case "AZURE":
            result = AZURE_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER
        case "GCP":
            result = GCP_STORAGELOCATIONSGETRESPONSE_ITEMS_CLOUDSERVICEPROVIDER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageLocationsGetResponse_items_cloudServiceProvider(values []StorageLocationsGetResponse_items_cloudServiceProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageLocationsGetResponse_items_cloudServiceProvider) isMultiValue() bool {
    return false
}
