package dataservicesconnectors
// Method for how data and time is set on the Data Services Connector.
type DataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet int

const (
    NTP_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DATETIME_METHODDATETIMESET DataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet = iota
)

func (i DataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet) String() string {
    return []string{"NTP"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet(v string) (any, error) {
    result := NTP_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DATETIME_METHODDATETIMESET
    switch v {
        case "NTP":
            result = NTP_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DATETIME_METHODDATETIMESET
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet(values []DataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_dateTime_methodDateTimeSet) isMultiValue() bool {
    return false
}