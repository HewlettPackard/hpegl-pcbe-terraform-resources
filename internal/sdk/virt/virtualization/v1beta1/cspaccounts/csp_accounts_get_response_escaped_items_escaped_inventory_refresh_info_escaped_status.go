package cspaccounts
// The result status of the most recent attempt to refresh inventory with the cloud service provider
type CspAccountsGetResponse_items_inventoryRefreshInfo_status int

const (
    OK_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS CspAccountsGetResponse_items_inventoryRefreshInfo_status = iota
    WARNING_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
    ERROR_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
    UNKNOWN_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
)

func (i CspAccountsGetResponse_items_inventoryRefreshInfo_status) String() string {
    return []string{"OK", "WARNING", "ERROR", "UNKNOWN"}[i]
}
func ParseCspAccountsGetResponse_items_inventoryRefreshInfo_status(v string) (any, error) {
    result := OK_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
    switch v {
        case "OK":
            result = OK_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
        case "WARNING":
            result = WARNING_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
        case "ERROR":
            result = ERROR_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
        case "UNKNOWN":
            result = UNKNOWN_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsGetResponse_items_inventoryRefreshInfo_status(values []CspAccountsGetResponse_items_inventoryRefreshInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsGetResponse_items_inventoryRefreshInfo_status) isMultiValue() bool {
    return false
}
