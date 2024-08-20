package item
// The cloud service provider type of the account to which the resource is associated.
type CspMachineInstancesGetResponse_cspType int

const (
    AWS_CSPMACHINEINSTANCESGETRESPONSE_CSPTYPE CspMachineInstancesGetResponse_cspType = iota
    AZURE_CSPMACHINEINSTANCESGETRESPONSE_CSPTYPE
)

func (i CspMachineInstancesGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspMachineInstancesGetResponse_cspType(v string) (any, error) {
    result := AWS_CSPMACHINEINSTANCESGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPMACHINEINSTANCESGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_CSPMACHINEINSTANCESGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_cspType(values []CspMachineInstancesGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_cspType) isMultiValue() bool {
    return false
}
