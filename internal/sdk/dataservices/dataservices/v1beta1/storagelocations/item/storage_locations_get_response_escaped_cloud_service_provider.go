package item
// The Cloud Service Provider (CSP) of the location.
type StorageLocationsGetResponse_cloudServiceProvider int

const (
    AWS_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER StorageLocationsGetResponse_cloudServiceProvider = iota
    AZURE_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER
    GCP_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER
)

func (i StorageLocationsGetResponse_cloudServiceProvider) String() string {
    return []string{"AWS", "AZURE", "GCP"}[i]
}
func ParseStorageLocationsGetResponse_cloudServiceProvider(v string) (any, error) {
    result := AWS_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER
    switch v {
        case "AWS":
            result = AWS_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER
        case "AZURE":
            result = AZURE_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER
        case "GCP":
            result = GCP_STORAGELOCATIONSGETRESPONSE_CLOUDSERVICEPROVIDER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStorageLocationsGetResponse_cloudServiceProvider(values []StorageLocationsGetResponse_cloudServiceProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StorageLocationsGetResponse_cloudServiceProvider) isMultiValue() bool {
    return false
}
