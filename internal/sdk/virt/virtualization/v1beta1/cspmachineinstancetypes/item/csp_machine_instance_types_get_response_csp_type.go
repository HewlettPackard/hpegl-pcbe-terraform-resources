package item
// The cloud service provider type
type CspMachineInstanceTypesGetResponse_cspType int

const (
    AWS_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPTYPE CspMachineInstanceTypesGetResponse_cspType = iota
    AZURE_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPTYPE
)

func (i CspMachineInstanceTypesGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_cspType(v string) (any, error) {
    result := AWS_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_cspType(values []CspMachineInstanceTypesGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_cspType) isMultiValue() bool {
    return false
}
