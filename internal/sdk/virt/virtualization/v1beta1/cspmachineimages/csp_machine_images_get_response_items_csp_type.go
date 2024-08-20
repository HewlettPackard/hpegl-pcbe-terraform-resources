package cspmachineimages
// The cloud service provider type
type CspMachineImagesGetResponse_items_cspType int

const (
    AWS_CSPMACHINEIMAGESGETRESPONSE_ITEMS_CSPTYPE CspMachineImagesGetResponse_items_cspType = iota
    AZURE_CSPMACHINEIMAGESGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspMachineImagesGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspMachineImagesGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPMACHINEIMAGESGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPMACHINEIMAGESGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPMACHINEIMAGESGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineImagesGetResponse_items_cspType(values []CspMachineImagesGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineImagesGetResponse_items_cspType) isMultiValue() bool {
    return false
}
