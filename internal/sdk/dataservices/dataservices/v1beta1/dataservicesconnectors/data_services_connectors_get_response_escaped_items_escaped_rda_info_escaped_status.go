package dataservicesconnectors
// Current status of RDA.
type DataServicesConnectorsGetResponse_items_rdaInfo_status int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS DataServicesConnectorsGetResponse_items_rdaInfo_status = iota
    WARNING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS
    ERROR_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS
)

func (i DataServicesConnectorsGetResponse_items_rdaInfo_status) String() string {
    return []string{"OK", "WARNING", "ERROR"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_rdaInfo_status(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS
        case "WARNING":
            result = WARNING_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS
        case "ERROR":
            result = ERROR_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_rdaInfo_status(values []DataServicesConnectorsGetResponse_items_rdaInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_rdaInfo_status) isMultiValue() bool {
    return false
}
