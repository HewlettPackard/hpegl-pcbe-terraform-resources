package cspaccounts
// The cloud service provider type of the account
type CspAccountsPostRequestBody_cspType int

const (
    AWS_CSPACCOUNTSPOSTREQUESTBODY_CSPTYPE CspAccountsPostRequestBody_cspType = iota
    AZURE_CSPACCOUNTSPOSTREQUESTBODY_CSPTYPE
)

func (i CspAccountsPostRequestBody_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseCspAccountsPostRequestBody_cspType(v string) (any, error) {
    result := AWS_CSPACCOUNTSPOSTREQUESTBODY_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_CSPACCOUNTSPOSTREQUESTBODY_CSPTYPE
        case "AZURE":
            result = AZURE_CSPACCOUNTSPOSTREQUESTBODY_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsPostRequestBody_cspType(values []CspAccountsPostRequestBody_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsPostRequestBody_cspType) isMultiValue() bool {
    return false
}
