package cspaccounts
// The status of the last validation of cloud service provider permissions and access
type CspAccountsGetResponse_items_validationStatus int

const (
    UNVALIDATED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS CspAccountsGetResponse_items_validationStatus = iota
    PASSED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS
    FAILED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS
)

func (i CspAccountsGetResponse_items_validationStatus) String() string {
    return []string{"UNVALIDATED", "PASSED", "FAILED"}[i]
}
func ParseCspAccountsGetResponse_items_validationStatus(v string) (any, error) {
    result := UNVALIDATED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS
    switch v {
        case "UNVALIDATED":
            result = UNVALIDATED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS
        case "PASSED":
            result = PASSED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS
        case "FAILED":
            result = FAILED_CSPACCOUNTSGETRESPONSE_ITEMS_VALIDATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsGetResponse_items_validationStatus(values []CspAccountsGetResponse_items_validationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsGetResponse_items_validationStatus) isMultiValue() bool {
    return false
}
