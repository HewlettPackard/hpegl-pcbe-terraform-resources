package servers
// type
type ServersGetResponse_items_iloNetworkInfo_networkInterfaces_type int

const (
    MANAGEMENT_SERVERSGETRESPONSE_ITEMS_ILONETWORKINFO_NETWORKINTERFACES_TYPE ServersGetResponse_items_iloNetworkInfo_networkInterfaces_type = iota
    DATA_SERVERSGETRESPONSE_ITEMS_ILONETWORKINFO_NETWORKINTERFACES_TYPE
)

func (i ServersGetResponse_items_iloNetworkInfo_networkInterfaces_type) String() string {
    return []string{"MANAGEMENT", "DATA"}[i]
}
func ParseServersGetResponse_items_iloNetworkInfo_networkInterfaces_type(v string) (any, error) {
    result := MANAGEMENT_SERVERSGETRESPONSE_ITEMS_ILONETWORKINFO_NETWORKINTERFACES_TYPE
    switch v {
        case "MANAGEMENT":
            result = MANAGEMENT_SERVERSGETRESPONSE_ITEMS_ILONETWORKINFO_NETWORKINTERFACES_TYPE
        case "DATA":
            result = DATA_SERVERSGETRESPONSE_ITEMS_ILONETWORKINFO_NETWORKINTERFACES_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_iloNetworkInfo_networkInterfaces_type(values []ServersGetResponse_items_iloNetworkInfo_networkInterfaces_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_iloNetworkInfo_networkInterfaces_type) isMultiValue() bool {
    return false
}
