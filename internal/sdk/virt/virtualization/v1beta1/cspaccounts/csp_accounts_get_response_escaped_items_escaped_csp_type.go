package cspaccounts
// The cloud service provider type of the account
type CspAccountsGetResponse_items_cspType int

const (
    AWS_CSPACCOUNTSGETRESPONSE_ITEMS_CSPTYPE CspAccountsGetResponse_items_cspType = iota
    AZURE_CSPACCOUNTSGETRESPONSE_ITEMS_CSPTYPE
)

func (i CspAccountsGetResponse_items_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspAccountsGetResponse_items_cspType(v string) (any, error) {
    result := AWS_CSPACCOUNTSGETRESPONSE_ITEMS_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPACCOUNTSGETRESPONSE_ITEMS_CSPTYPE
        case "AZURE":
            result = AZURE_CSPACCOUNTSGETRESPONSE_ITEMS_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsGetResponse_items_cspType(values []CspAccountsGetResponse_items_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsGetResponse_items_cspType) isMultiValue() bool {
    return false
}
