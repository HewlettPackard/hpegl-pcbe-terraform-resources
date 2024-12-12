package dataservicesconnectors
// State of NTP server
type DataServicesConnectorsGetResponse_items_ntp_state int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_NTP_STATE DataServicesConnectorsGetResponse_items_ntp_state = iota
    DEGRADED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_NTP_STATE
)

func (i DataServicesConnectorsGetResponse_items_ntp_state) String() string {
    return []string{"OK", "DEGRADED"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_ntp_state(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_NTP_STATE
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_NTP_STATE
        case "DEGRADED":
            result = DEGRADED_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_NTP_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_ntp_state(values []DataServicesConnectorsGetResponse_items_ntp_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_ntp_state) isMultiValue() bool {
    return false
}
