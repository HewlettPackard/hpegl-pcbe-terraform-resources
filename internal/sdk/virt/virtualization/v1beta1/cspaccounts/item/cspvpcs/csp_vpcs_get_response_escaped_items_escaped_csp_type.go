package cspvpcs
// The cloud service provider type of the account to which the resource is associated.
type CspVpcsGetResponse_items_cspType int

const (
    AWS_CSPVPCSGETRESPONSE_ITEMS_CSPTYPE CspVpcsGetResponse_items_cspType = iota
    AZURE_CSPVPCSGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspVpcsGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspVpcsGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPVPCSGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPVPCSGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPVPCSGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVpcsGetResponse_items_cspType(values []CspVpcsGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVpcsGetResponse_items_cspType) isMultiValue() bool {
    return false
}
