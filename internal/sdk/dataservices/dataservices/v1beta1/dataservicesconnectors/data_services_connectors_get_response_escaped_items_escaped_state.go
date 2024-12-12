package dataservicesconnectors
// Current state of Data Services Connector.
type DataServicesConnectorsGetResponse_items_state int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE DataServicesConnectorsGetResponse_items_state = iota
    UPGRADING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
    UPDATING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
    UNKNOWN_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
)

func (i DataServicesConnectorsGetResponse_items_state) String() string {
    return []string{"OK", "UPGRADING", "UPDATING", "UNKNOWN"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_state(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
        case "UPGRADING":
            result = UPGRADING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
        case "UPDATING":
            result = UPDATING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
        case "UNKNOWN":
            result = UNKNOWN_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_state(values []DataServicesConnectorsGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_state) isMultiValue() bool {
    return false
}
