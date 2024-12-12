package item
// Current state of Data Services Connector.
type DataServicesConnectorsGetResponse_state int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_STATE DataServicesConnectorsGetResponse_state = iota
    UPGRADING_DATASERVICESCONNECTORSGETRESPONSE_STATE
    UPDATING_DATASERVICESCONNECTORSGETRESPONSE_STATE
    UNKNOWN_DATASERVICESCONNECTORSGETRESPONSE_STATE
)

func (i DataServicesConnectorsGetResponse_state) String() string {
    return []string{"OK", "UPGRADING", "UPDATING", "UNKNOWN"}[i]
}
func ParseDataServicesConnectorsGetResponse_state(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_STATE
        case "UPGRADING":
            result = UPGRADING_DATASERVICESCONNECTORSGETRESPONSE_STATE
        case "UPDATING":
            result = UPDATING_DATASERVICESCONNECTORSGETRESPONSE_STATE
        case "UNKNOWN":
            result = UNKNOWN_DATASERVICESCONNECTORSGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_state(values []DataServicesConnectorsGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_state) isMultiValue() bool {
    return false
}
