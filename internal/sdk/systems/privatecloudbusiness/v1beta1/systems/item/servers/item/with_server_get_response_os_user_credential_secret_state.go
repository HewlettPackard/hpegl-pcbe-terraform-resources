package item
// secret state
type WithServerGetResponse_osUserCredential_secret_state int

const (
    UNSPECIFIED_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE WithServerGetResponse_osUserCredential_secret_state = iota
    INPROGRESS_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
    INCONSISTENT_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
    INSYNC_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
)

func (i WithServerGetResponse_osUserCredential_secret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseWithServerGetResponse_osUserCredential_secret_state(v string) (any, error) {
    result := UNSPECIFIED_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
        case "INSYNC":
            result = INSYNC_WITHSERVERGETRESPONSE_OSUSERCREDENTIAL_SECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_osUserCredential_secret_state(values []WithServerGetResponse_osUserCredential_secret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_osUserCredential_secret_state) isMultiValue() bool {
    return false
}
