package item
// The cloud service provider type of the account to which the resource is associated.
type WithSubnetGetResponse_cspType int

const (
    AWS_WITHSUBNETGETRESPONSE_CSPTYPE WithSubnetGetResponse_cspType = iota
    AZURE_WITHSUBNETGETRESPONSE_CSPTYPE
)

func (i WithSubnetGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseWithSubnetGetResponse_cspType(v string) (any, error) {
    result := AWS_WITHSUBNETGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_WITHSUBNETGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_WITHSUBNETGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSubnetGetResponse_cspType(values []WithSubnetGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSubnetGetResponse_cspType) isMultiValue() bool {
    return false
}
