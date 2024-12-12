package item
// The result status of the most recent attempt to refresh inventory with the cloud service provider
type AccountGetResponse_inventoryRefreshInfo_status int

const (
    OK_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS AccountGetResponse_inventoryRefreshInfo_status = iota
    WARNING_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
    ERROR_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
    UNKNOWN_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
)

func (i AccountGetResponse_inventoryRefreshInfo_status) String() string {
    return []string{"OK", "WARNING", "ERROR", "UNKNOWN"}[i]
}
func ParseAccountGetResponse_inventoryRefreshInfo_status(v string) (any, error) {
    result := OK_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
    switch v {
        case "OK":
            result = OK_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
        case "WARNING":
            result = WARNING_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
        case "ERROR":
            result = ERROR_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
        case "UNKNOWN":
            result = UNKNOWN_ACCOUNTGETRESPONSE_INVENTORYREFRESHINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAccountGetResponse_inventoryRefreshInfo_status(values []AccountGetResponse_inventoryRefreshInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccountGetResponse_inventoryRefreshInfo_status) isMultiValue() bool {
    return false
}
