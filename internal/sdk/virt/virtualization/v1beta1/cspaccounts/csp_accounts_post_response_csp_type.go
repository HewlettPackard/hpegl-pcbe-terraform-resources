package cspaccounts
// The cloud service provider type of the account
type CspAccountsPostResponse_cspType int

const (
    AWS_CSPACCOUNTSPOSTRESPONSE_CSPTYPE CspAccountsPostResponse_cspType = iota
    AZURE_CSPACCOUNTSPOSTRESPONSE_CSPTYPE
)

func (i CspAccountsPostResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspAccountsPostResponse_cspType(v string) (any, error) {
    result := AWS_CSPACCOUNTSPOSTRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPACCOUNTSPOSTRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_CSPACCOUNTSPOSTRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsPostResponse_cspType(values []CspAccountsPostResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsPostResponse_cspType) isMultiValue() bool {
    return false
}
