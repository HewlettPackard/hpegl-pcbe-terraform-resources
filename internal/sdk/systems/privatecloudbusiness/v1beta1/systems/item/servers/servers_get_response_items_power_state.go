package servers
// Server Power state.
type ServersGetResponse_items_powerState int

const (
    ON_SERVERSGETRESPONSE_ITEMS_POWERSTATE ServersGetResponse_items_powerState = iota
    RESET_SERVERSGETRESPONSE_ITEMS_POWERSTATE
    UNKNOWN_SERVERSGETRESPONSE_ITEMS_POWERSTATE
)

func (i ServersGetResponse_items_powerState) String() string {
    return []string{"ON", "RESET", "UNKNOWN"}[i]
}
func ParseServersGetResponse_items_powerState(v string) (any, error) {
    result := ON_SERVERSGETRESPONSE_ITEMS_POWERSTATE
    switch v {
        case "ON":
            result = ON_SERVERSGETRESPONSE_ITEMS_POWERSTATE
        case "RESET":
            result = RESET_SERVERSGETRESPONSE_ITEMS_POWERSTATE
        case "UNKNOWN":
            result = UNKNOWN_SERVERSGETRESPONSE_ITEMS_POWERSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_powerState(values []ServersGetResponse_items_powerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_powerState) isMultiValue() bool {
    return false
}
