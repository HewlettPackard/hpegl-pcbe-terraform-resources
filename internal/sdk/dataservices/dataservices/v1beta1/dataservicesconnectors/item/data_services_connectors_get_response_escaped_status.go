package item
// Current status of Data Services Connector.
type DataServicesConnectorsGetResponse_status int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_STATUS DataServicesConnectorsGetResponse_status = iota
    WARNING_DATASERVICESCONNECTORSGETRESPONSE_STATUS
    ERROR_DATASERVICESCONNECTORSGETRESPONSE_STATUS
)

func (i DataServicesConnectorsGetResponse_status) String() string {
    return []string{"OK", "WARNING", "ERROR"}[i]
}
func ParseDataServicesConnectorsGetResponse_status(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_DATASERVICESCONNECTORSGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_DATASERVICESCONNECTORSGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_status(values []DataServicesConnectorsGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_status) isMultiValue() bool {
    return false
}
