package cspmachineinstancetypes
// The cloud service provider type
type CspMachineInstanceTypesGetResponse_items_cspType int

const (
    AWS_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPTYPE CspMachineInstanceTypesGetResponse_items_cspType = iota
    AZURE_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspMachineInstanceTypesGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_items_cspType(values []CspMachineInstanceTypesGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_items_cspType) isMultiValue() bool {
    return false
}
