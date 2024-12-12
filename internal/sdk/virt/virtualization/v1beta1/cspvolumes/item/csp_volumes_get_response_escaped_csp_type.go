package item
// The cloud service provider type of the account to which the resource is associated.
type CspVolumesGetResponse_cspType int

const (
    AWS_CSPVOLUMESGETRESPONSE_CSPTYPE CspVolumesGetResponse_cspType = iota
    AZURE_CSPVOLUMESGETRESPONSE_CSPTYPE
)

func (i CspVolumesGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspVolumesGetResponse_cspType(v string) (any, error) {
    result := AWS_CSPVOLUMESGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPVOLUMESGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_CSPVOLUMESGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVolumesGetResponse_cspType(values []CspVolumesGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVolumesGetResponse_cspType) isMultiValue() bool {
    return false
}
