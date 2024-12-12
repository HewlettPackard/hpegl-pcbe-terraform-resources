package dataservicesconnectors
// Connection state of Data Services Connector to DSCC.
type DataServicesConnectorsGetResponse_items_connectionState int

const (
    CONNECTED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_CONNECTIONSTATE DataServicesConnectorsGetResponse_items_connectionState = iota
    DISCONNECTED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_CONNECTIONSTATE
)

func (i DataServicesConnectorsGetResponse_items_connectionState) String() string {
    return []string{"CONNECTED", "DISCONNECTED"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_connectionState(v string) (any, error) {
    result := CONNECTED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_CONNECTIONSTATE
    switch v {
        case "CONNECTED":
            result = CONNECTED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_CONNECTIONSTATE
        case "DISCONNECTED":
            result = DISCONNECTED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_CONNECTIONSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_connectionState(values []DataServicesConnectorsGetResponse_items_connectionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_connectionState) isMultiValue() bool {
    return false
}
