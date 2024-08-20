package cspvolumes
// The cloud service provider type of the account to which the resource is associated.
type CspVolumesGetResponse_items_cspType int

const (
    AWS_CSPVOLUMESGETRESPONSE_ITEMS_CSPTYPE CspVolumesGetResponse_items_cspType = iota
    AZURE_CSPVOLUMESGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspVolumesGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspVolumesGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPVOLUMESGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPVOLUMESGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPVOLUMESGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVolumesGetResponse_items_cspType(values []CspVolumesGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVolumesGetResponse_items_cspType) isMultiValue() bool {
    return false
}
