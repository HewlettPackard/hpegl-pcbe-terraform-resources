package cspsubnets
// The cloud service provider type of the account to which the resource is associated.
type CspSubnetsGetResponse_items_cspType int

const (
    AWS_CSPSUBNETSGETRESPONSE_ITEMS_CSPTYPE CspSubnetsGetResponse_items_cspType = iota
    AZURE_CSPSUBNETSGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspSubnetsGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspSubnetsGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPSUBNETSGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPSUBNETSGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPSUBNETSGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspSubnetsGetResponse_items_cspType(values []CspSubnetsGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspSubnetsGetResponse_items_cspType) isMultiValue() bool {
    return false
}
