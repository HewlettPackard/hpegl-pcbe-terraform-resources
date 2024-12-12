package dataservicesconnectors
// Current state of RDA.
type DataServicesConnectorsGetResponse_items_rdaInfo_state int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATE DataServicesConnectorsGetResponse_items_rdaInfo_state = iota
    UNKNOWN_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATE
)

func (i DataServicesConnectorsGetResponse_items_rdaInfo_state) String() string {
    return []string{"OK", "UNKNOWN"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_rdaInfo_state(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATE
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATE
        case "UNKNOWN":
            result = UNKNOWN_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_RDAINFO_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_rdaInfo_state(values []DataServicesConnectorsGetResponse_items_rdaInfo_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_rdaInfo_state) isMultiValue() bool {
    return false
}
