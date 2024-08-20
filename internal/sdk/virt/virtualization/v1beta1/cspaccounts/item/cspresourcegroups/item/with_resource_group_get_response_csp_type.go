package item
// The cloud service provider type of the account to which the resource is associated.
type WithResourceGroupGetResponse_cspType int

const (
    AWS_WITHRESOURCEGROUPGETRESPONSE_CSPTYPE WithResourceGroupGetResponse_cspType = iota
    AZURE_WITHRESOURCEGROUPGETRESPONSE_CSPTYPE
)

func (i WithResourceGroupGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseWithResourceGroupGetResponse_cspType(v string) (any, error) {
    result := AWS_WITHRESOURCEGROUPGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_WITHRESOURCEGROUPGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_WITHRESOURCEGROUPGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithResourceGroupGetResponse_cspType(values []WithResourceGroupGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithResourceGroupGetResponse_cspType) isMultiValue() bool {
    return false
}
