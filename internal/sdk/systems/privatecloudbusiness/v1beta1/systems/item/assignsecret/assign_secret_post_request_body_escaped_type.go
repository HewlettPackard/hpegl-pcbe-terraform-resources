package assignsecret
// Type of the resource.
type AssignSecretPostRequestBody_type int

const (
    ILO_ASSIGNSECRETPOSTREQUESTBODY_TYPE AssignSecretPostRequestBody_type = iota
    ESXI_ASSIGNSECRETPOSTREQUESTBODY_TYPE
)

func (i AssignSecretPostRequestBody_type) String() string {
    return []string{"ilo", "esxi"}[i]
}
func ParseAssignSecretPostRequestBody_type(v string) (any, error) {
    result := ILO_ASSIGNSECRETPOSTREQUESTBODY_TYPE
    switch v {
        case "ilo":
            result = ILO_ASSIGNSECRETPOSTREQUESTBODY_TYPE
        case "esxi":
            result = ESXI_ASSIGNSECRETPOSTREQUESTBODY_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAssignSecretPostRequestBody_type(values []AssignSecretPostRequestBody_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AssignSecretPostRequestBody_type) isMultiValue() bool {
    return false
}
