package csprdsinstances
// The result status of the most recent attempt to refresh the inventory with the cloud service provider
type CspRdsInstancesGetResponse_items_refreshInfo_status int

const (
    OK_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS CspRdsInstancesGetResponse_items_refreshInfo_status = iota
    WARNING_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
    ERROR_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
    UNKNOWN_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
)

func (i CspRdsInstancesGetResponse_items_refreshInfo_status) String() string {
    return []string{"OK", "WARNING", "ERROR", "UNKNOWN"}[i]
}
func ParseCspRdsInstancesGetResponse_items_refreshInfo_status(v string) (any, error) {
    result := OK_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
    switch v {
        case "OK":
            result = OK_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
        case "WARNING":
            result = WARNING_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
        case "ERROR":
            result = ERROR_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
        case "UNKNOWN":
            result = UNKNOWN_CSPRDSINSTANCESGETRESPONSE_ITEMS_REFRESHINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_items_refreshInfo_status(values []CspRdsInstancesGetResponse_items_refreshInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_items_refreshInfo_status) isMultiValue() bool {
    return false
}
