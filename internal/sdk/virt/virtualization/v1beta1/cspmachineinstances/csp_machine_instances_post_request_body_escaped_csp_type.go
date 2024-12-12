package cspmachineinstances
// The cloud service provider type
type CspMachineInstancesPostRequestBody_cspType int

const (
    AWS_CSPMACHINEINSTANCESPOSTREQUESTBODY_CSPTYPE CspMachineInstancesPostRequestBody_cspType = iota
    AZURE_CSPMACHINEINSTANCESPOSTREQUESTBODY_CSPTYPE
)

func (i CspMachineInstancesPostRequestBody_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspMachineInstancesPostRequestBody_cspType(v string) (any, error) {
    result := AWS_CSPMACHINEINSTANCESPOSTREQUESTBODY_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPMACHINEINSTANCESPOSTREQUESTBODY_CSPTYPE
        case "AZURE":
            result = AZURE_CSPMACHINEINSTANCESPOSTREQUESTBODY_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesPostRequestBody_cspType(values []CspMachineInstancesPostRequestBody_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesPostRequestBody_cspType) isMultiValue() bool {
    return false
}
