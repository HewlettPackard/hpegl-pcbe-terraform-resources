package servers
// secret state
type ServersGetResponse_items_osUserCredential_secret_state int

const (
    UNSPECIFIED_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE ServersGetResponse_items_osUserCredential_secret_state = iota
    INPROGRESS_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
    INCONSISTENT_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
    INSYNC_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
)

func (i ServersGetResponse_items_osUserCredential_secret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseServersGetResponse_items_osUserCredential_secret_state(v string) (any, error) {
    result := UNSPECIFIED_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
        case "INSYNC":
            result = INSYNC_SERVERSGETRESPONSE_ITEMS_OSUSERCREDENTIAL_SECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_osUserCredential_secret_state(values []ServersGetResponse_items_osUserCredential_secret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_osUserCredential_secret_state) isMultiValue() bool {
    return false
}
