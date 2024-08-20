package cspaccounts
// The result status of the most recent attempt to refresh inventory with the cloud service provider
type CspAccountsPostResponse_inventoryRefreshInfo_status int

const (
    OK_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS CspAccountsPostResponse_inventoryRefreshInfo_status = iota
    WARNING_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
    ERROR_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
    UNKNOWN_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
)

func (i CspAccountsPostResponse_inventoryRefreshInfo_status) String() string {
    return []string{"OK", "WARNING", "ERROR", "UNKNOWN"}[i]
}
func ParseCspAccountsPostResponse_inventoryRefreshInfo_status(v string) (any, error) {
    result := OK_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
    switch v {
        case "OK":
            result = OK_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
        case "WARNING":
            result = WARNING_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
        case "ERROR":
            result = ERROR_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
        case "UNKNOWN":
            result = UNKNOWN_CSPACCOUNTSPOSTRESPONSE_INVENTORYREFRESHINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsPostResponse_inventoryRefreshInfo_status(values []CspAccountsPostResponse_inventoryRefreshInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsPostResponse_inventoryRefreshInfo_status) isMultiValue() bool {
    return false
}
