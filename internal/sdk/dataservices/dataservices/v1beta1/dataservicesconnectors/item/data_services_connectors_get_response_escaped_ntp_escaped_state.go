package item
// State of NTP server
type DataServicesConnectorsGetResponse_ntp_state int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATE DataServicesConnectorsGetResponse_ntp_state = iota
    DEGRADED_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATE
)

func (i DataServicesConnectorsGetResponse_ntp_state) String() string {
    return []string{"OK", "DEGRADED"}[i]
}
func ParseDataServicesConnectorsGetResponse_ntp_state(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATE
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATE
        case "DEGRADED":
            result = DEGRADED_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_ntp_state(values []DataServicesConnectorsGetResponse_ntp_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_ntp_state) isMultiValue() bool {
    return false
}
