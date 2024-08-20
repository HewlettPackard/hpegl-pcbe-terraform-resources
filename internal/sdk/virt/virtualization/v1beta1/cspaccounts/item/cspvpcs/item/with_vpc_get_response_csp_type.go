package item
// The cloud service provider type of the account to which the resource is associated.
type WithVpcGetResponse_cspType int

const (
    AWS_WITHVPCGETRESPONSE_CSPTYPE WithVpcGetResponse_cspType = iota
    AZURE_WITHVPCGETRESPONSE_CSPTYPE
)

func (i WithVpcGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseWithVpcGetResponse_cspType(v string) (any, error) {
    result := AWS_WITHVPCGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_WITHVPCGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_WITHVPCGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithVpcGetResponse_cspType(values []WithVpcGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithVpcGetResponse_cspType) isMultiValue() bool {
    return false
}
