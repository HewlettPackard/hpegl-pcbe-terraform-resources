package item
// The cloud service provider type of the account to which the resource is associated.
type WithSubscriptionGetResponse_cspType int

const (
    AWS_WITHSUBSCRIPTIONGETRESPONSE_CSPTYPE WithSubscriptionGetResponse_cspType = iota
    AZURE_WITHSUBSCRIPTIONGETRESPONSE_CSPTYPE
)

func (i WithSubscriptionGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseWithSubscriptionGetResponse_cspType(v string) (any, error) {
    result := AWS_WITHSUBSCRIPTIONGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_WITHSUBSCRIPTIONGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_WITHSUBSCRIPTIONGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSubscriptionGetResponse_cspType(values []WithSubscriptionGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSubscriptionGetResponse_cspType) isMultiValue() bool {
    return false
}
