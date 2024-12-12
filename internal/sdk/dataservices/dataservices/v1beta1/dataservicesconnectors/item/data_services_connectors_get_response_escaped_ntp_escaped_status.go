package item
// Status of NTP server
type DataServicesConnectorsGetResponse_ntp_status int

const (
    OK_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS DataServicesConnectorsGetResponse_ntp_status = iota
    WARNING_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS
    ERROR_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS
)

func (i DataServicesConnectorsGetResponse_ntp_status) String() string {
    return []string{"OK", "WARNING", "ERROR"}[i]
}
func ParseDataServicesConnectorsGetResponse_ntp_status(v string) (any, error) {
    result := OK_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS
    switch v {
        case "OK":
            result = OK_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS
        case "WARNING":
            result = WARNING_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS
        case "ERROR":
            result = ERROR_DATASERVICESCONNECTORSGETRESPONSE_NTP_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_ntp_status(values []DataServicesConnectorsGetResponse_ntp_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_ntp_status) isMultiValue() bool {
    return false
}
