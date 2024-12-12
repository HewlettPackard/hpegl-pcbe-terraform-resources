package cspsubscriptions
// The cloud service provider type of the account to which the resource is associated.
type CspSubscriptionsGetResponse_items_cspType int

const (
    AWS_CSPSUBSCRIPTIONSGETRESPONSE_ITEMS_CSPTYPE CspSubscriptionsGetResponse_items_cspType = iota
    AZURE_CSPSUBSCRIPTIONSGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspSubscriptionsGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspSubscriptionsGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPSUBSCRIPTIONSGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPSUBSCRIPTIONSGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPSUBSCRIPTIONSGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspSubscriptionsGetResponse_items_cspType(values []CspSubscriptionsGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspSubscriptionsGetResponse_items_cspType) isMultiValue() bool {
    return false
}
