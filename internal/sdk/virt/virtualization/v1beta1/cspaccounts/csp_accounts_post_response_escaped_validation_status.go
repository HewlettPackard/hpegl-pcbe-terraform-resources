package cspaccounts
// The status of the last validation of cloud service provider permissions and access
type CspAccountsPostResponse_validationStatus int

const (
    UNVALIDATED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS CspAccountsPostResponse_validationStatus = iota
    PASSED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS
    FAILED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS
)

func (i CspAccountsPostResponse_validationStatus) String() string {
    return []string{"UNVALIDATED", "PASSED", "FAILED"}[i]
}
func ParseCspAccountsPostResponse_validationStatus(v string) (any, error) {
    result := UNVALIDATED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS
    switch v {
        case "UNVALIDATED":
            result = UNVALIDATED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS
        case "PASSED":
            result = PASSED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS
        case "FAILED":
            result = FAILED_CSPACCOUNTSPOSTRESPONSE_VALIDATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsPostResponse_validationStatus(values []CspAccountsPostResponse_validationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsPostResponse_validationStatus) isMultiValue() bool {
    return false
}
