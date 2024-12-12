package item
// type
type WithServerGetResponse_serverNetworkInterfaces_type int

const (
    MANAGEMENT_WITHSERVERGETRESPONSE_SERVERNETWORKINTERFACES_TYPE WithServerGetResponse_serverNetworkInterfaces_type = iota
    DATA_WITHSERVERGETRESPONSE_SERVERNETWORKINTERFACES_TYPE
)

func (i WithServerGetResponse_serverNetworkInterfaces_type) String() string {
    return []string{"MANAGEMENT", "DATA"}[i]
}
func ParseWithServerGetResponse_serverNetworkInterfaces_type(v string) (any, error) {
    result := MANAGEMENT_WITHSERVERGETRESPONSE_SERVERNETWORKINTERFACES_TYPE
    switch v {
        case "MANAGEMENT":
            result = MANAGEMENT_WITHSERVERGETRESPONSE_SERVERNETWORKINTERFACES_TYPE
        case "DATA":
            result = DATA_WITHSERVERGETRESPONSE_SERVERNETWORKINTERFACES_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_serverNetworkInterfaces_type(values []WithServerGetResponse_serverNetworkInterfaces_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_serverNetworkInterfaces_type) isMultiValue() bool {
    return false
}
