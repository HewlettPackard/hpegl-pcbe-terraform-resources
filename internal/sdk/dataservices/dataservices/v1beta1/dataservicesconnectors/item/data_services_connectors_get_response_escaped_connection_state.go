package item
// Connection state of Data Services Connector to DSCC.
type DataServicesConnectorsGetResponse_connectionState int

const (
    CONNECTED_DATASERVICESCONNECTORSGETRESPONSE_CONNECTIONSTATE DataServicesConnectorsGetResponse_connectionState = iota
    DISCONNECTED_DATASERVICESCONNECTORSGETRESPONSE_CONNECTIONSTATE
)

func (i DataServicesConnectorsGetResponse_connectionState) String() string {
    return []string{"CONNECTED", "DISCONNECTED"}[i]
}
func ParseDataServicesConnectorsGetResponse_connectionState(v string) (any, error) {
    result := CONNECTED_DATASERVICESCONNECTORSGETRESPONSE_CONNECTIONSTATE
    switch v {
        case "CONNECTED":
            result = CONNECTED_DATASERVICESCONNECTORSGETRESPONSE_CONNECTIONSTATE
        case "DISCONNECTED":
            result = DISCONNECTED_DATASERVICESCONNECTORSGETRESPONSE_CONNECTIONSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_connectionState(values []DataServicesConnectorsGetResponse_connectionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_connectionState) isMultiValue() bool {
    return false
}
