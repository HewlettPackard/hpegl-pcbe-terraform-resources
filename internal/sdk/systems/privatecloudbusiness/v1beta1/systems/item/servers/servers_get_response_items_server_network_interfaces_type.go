package servers
// type
type ServersGetResponse_items_serverNetworkInterfaces_type int

const (
    MANAGEMENT_SERVERSGETRESPONSE_ITEMS_SERVERNETWORKINTERFACES_TYPE ServersGetResponse_items_serverNetworkInterfaces_type = iota
    DATA_SERVERSGETRESPONSE_ITEMS_SERVERNETWORKINTERFACES_TYPE
)

func (i ServersGetResponse_items_serverNetworkInterfaces_type) String() string {
    return []string{"MANAGEMENT", "DATA"}[i]
}
func ParseServersGetResponse_items_serverNetworkInterfaces_type(v string) (any, error) {
    result := MANAGEMENT_SERVERSGETRESPONSE_ITEMS_SERVERNETWORKINTERFACES_TYPE
    switch v {
        case "MANAGEMENT":
            result = MANAGEMENT_SERVERSGETRESPONSE_ITEMS_SERVERNETWORKINTERFACES_TYPE
        case "DATA":
            result = DATA_SERVERSGETRESPONSE_ITEMS_SERVERNETWORKINTERFACES_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_serverNetworkInterfaces_type(values []ServersGetResponse_items_serverNetworkInterfaces_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_serverNetworkInterfaces_type) isMultiValue() bool {
    return false
}
