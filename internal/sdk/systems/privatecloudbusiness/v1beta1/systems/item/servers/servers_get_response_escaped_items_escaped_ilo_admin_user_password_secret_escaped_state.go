package servers
// secret state
type ServersGetResponse_items_iloAdminUserPasswordSecret_state int

const (
    UNSPECIFIED_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE ServersGetResponse_items_iloAdminUserPasswordSecret_state = iota
    INPROGRESS_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
    INCONSISTENT_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
    INSYNC_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
)

func (i ServersGetResponse_items_iloAdminUserPasswordSecret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseServersGetResponse_items_iloAdminUserPasswordSecret_state(v string) (any, error) {
    result := UNSPECIFIED_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
        case "INSYNC":
            result = INSYNC_SERVERSGETRESPONSE_ITEMS_ILOADMINUSERPASSWORDSECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_iloAdminUserPasswordSecret_state(values []ServersGetResponse_items_iloAdminUserPasswordSecret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_iloAdminUserPasswordSecret_state) isMultiValue() bool {
    return false
}
