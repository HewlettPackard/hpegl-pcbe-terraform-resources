package item
// type
type WithServerGetResponse_iloNetworkInfo_networkInterfaces_type int

const (
    MANAGEMENT_WITHSERVERGETRESPONSE_ILONETWORKINFO_NETWORKINTERFACES_TYPE WithServerGetResponse_iloNetworkInfo_networkInterfaces_type = iota
    DATA_WITHSERVERGETRESPONSE_ILONETWORKINFO_NETWORKINTERFACES_TYPE
)

func (i WithServerGetResponse_iloNetworkInfo_networkInterfaces_type) String() string {
    return []string{"MANAGEMENT", "DATA"}[i]
}
func ParseWithServerGetResponse_iloNetworkInfo_networkInterfaces_type(v string) (any, error) {
    result := MANAGEMENT_WITHSERVERGETRESPONSE_ILONETWORKINFO_NETWORKINTERFACES_TYPE
    switch v {
        case "MANAGEMENT":
            result = MANAGEMENT_WITHSERVERGETRESPONSE_ILONETWORKINFO_NETWORKINTERFACES_TYPE
        case "DATA":
            result = DATA_WITHSERVERGETRESPONSE_ILONETWORKINFO_NETWORKINTERFACES_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_iloNetworkInfo_networkInterfaces_type(values []WithServerGetResponse_iloNetworkInfo_networkInterfaces_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_iloNetworkInfo_networkInterfaces_type) isMultiValue() bool {
    return false
}
