package dataservicesconnectors
// Current status of Data Services Connector.
type DataServicesConnectorsGetResponse_items_status int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS DataServicesConnectorsGetResponse_items_status = iota
    WARNING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS
    ERROR_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS
)

func (i DataServicesConnectorsGetResponse_items_status) String() string {
    return []string{"OK", "WARNING", "ERROR"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_status(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_status(values []DataServicesConnectorsGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_status) isMultiValue() bool {
    return false
}
