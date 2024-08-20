package cspresourcegroups
// The cloud service provider type of the account to which the resource is associated.
type CspResourceGroupsGetResponse_items_cspType int

const (
    AWS_CSPRESOURCEGROUPSGETRESPONSE_ITEMS_CSPTYPE CspResourceGroupsGetResponse_items_cspType = iota
    AZURE_CSPRESOURCEGROUPSGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspResourceGroupsGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspResourceGroupsGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPRESOURCEGROUPSGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPRESOURCEGROUPSGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPRESOURCEGROUPSGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspResourceGroupsGetResponse_items_cspType(values []CspResourceGroupsGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspResourceGroupsGetResponse_items_cspType) isMultiValue() bool {
    return false
}
