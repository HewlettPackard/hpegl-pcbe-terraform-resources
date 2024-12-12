package item
// secret state
type WithServerGetResponse_iloAdminUserPasswordSecret_state int

const (
    UNSPECIFIED_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE WithServerGetResponse_iloAdminUserPasswordSecret_state = iota
    INPROGRESS_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
    INCONSISTENT_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
    INSYNC_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
)

func (i WithServerGetResponse_iloAdminUserPasswordSecret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseWithServerGetResponse_iloAdminUserPasswordSecret_state(v string) (any, error) {
    result := UNSPECIFIED_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
        case "INSYNC":
            result = INSYNC_WITHSERVERGETRESPONSE_ILOADMINUSERPASSWORDSECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_iloAdminUserPasswordSecret_state(values []WithServerGetResponse_iloAdminUserPasswordSecret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_iloAdminUserPasswordSecret_state) isMultiValue() bool {
    return false
}
