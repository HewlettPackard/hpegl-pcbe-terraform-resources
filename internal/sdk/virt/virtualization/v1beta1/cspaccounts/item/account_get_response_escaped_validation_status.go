package item
// The status of the last validation of cloud service provider permissions and access
type AccountGetResponse_validationStatus int

const (
    UNVALIDATED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS AccountGetResponse_validationStatus = iota
    PASSED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS
    FAILED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS
)

func (i AccountGetResponse_validationStatus) String() string {
    return []string{"UNVALIDATED", "PASSED", "FAILED"}[i]
}
func ParseAccountGetResponse_validationStatus(v string) (any, error) {
    result := UNVALIDATED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS
    switch v {
        case "UNVALIDATED":
            result = UNVALIDATED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS
        case "PASSED":
            result = PASSED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS
        case "FAILED":
            result = FAILED_ACCOUNTGETRESPONSE_VALIDATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAccountGetResponse_validationStatus(values []AccountGetResponse_validationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccountGetResponse_validationStatus) isMultiValue() bool {
    return false
}
