package item
// The cloud service provider type of the account
type AccountGetResponse_cspType int

const (
    AWS_ACCOUNTGETRESPONSE_CSPTYPE AccountGetResponse_cspType = iota
    AZURE_ACCOUNTGETRESPONSE_CSPTYPE
)

func (i AccountGetResponse_cspType) String() string {
    return []string{"AWS", "AZURE"}[i]
}
func ParseAccountGetResponse_cspType(v string) (any, error) {
    result := AWS_ACCOUNTGETRESPONSE_CSPTYPE
    switch v {
        case "AWS":
            result = AWS_ACCOUNTGETRESPONSE_CSPTYPE
        case "AZURE":
            result = AZURE_ACCOUNTGETRESPONSE_CSPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAccountGetResponse_cspType(values []AccountGetResponse_cspType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccountGetResponse_cspType) isMultiValue() bool {
    return false
}
