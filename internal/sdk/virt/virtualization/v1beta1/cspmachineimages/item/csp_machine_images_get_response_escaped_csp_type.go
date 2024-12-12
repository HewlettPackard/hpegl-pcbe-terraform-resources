package item
// The cloud service provider type
type CspMachineImagesGetResponse_cspType int

const (
    AWS_CSPMACHINEIMAGESGETRESPONSE_CSPTYPE CspMachineImagesGetResponse_cspType = iota
    AZURE_CSPMACHINEIMAGESGETRESPONSE_CSPTYPE
)

func (i CspMachineImagesGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspMachineImagesGetResponse_cspType(v string) (any, error) {
    result := AWS_CSPMACHINEIMAGESGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPMACHINEIMAGESGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_CSPMACHINEIMAGESGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineImagesGetResponse_cspType(values []CspMachineImagesGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineImagesGetResponse_cspType) isMultiValue() bool {
    return false
}
