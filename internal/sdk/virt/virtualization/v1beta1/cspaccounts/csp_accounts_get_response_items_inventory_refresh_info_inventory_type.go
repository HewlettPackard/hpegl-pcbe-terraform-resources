package cspaccounts
// A category of inventory resources refreshed with the cloud service provider
type CspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType int

const (
    MACHINE_INSTANCES_AND_VOLUMES_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE CspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType = iota
    RDS_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE
    KUBERNETES_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE
)

func (i CspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType) String() string {
    return []string{"MACHINE_INSTANCES_AND_VOLUMES", "RDS", "KUBERNETES"}[i]
}
func ParseCspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType(v string) (any, error) {
    result := MACHINE_INSTANCES_AND_VOLUMES_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE
    switch v {
        case "MACHINE_INSTANCES_AND_VOLUMES":
            result = MACHINE_INSTANCES_AND_VOLUMES_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE
        case "RDS":
            result = RDS_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE
        case "KUBERNETES":
            result = KUBERNETES_CSPACCOUNTSGETRESPONSE_ITEMS_INVENTORYREFRESHINFO_INVENTORYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType(values []CspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspAccountsGetResponse_items_inventoryRefreshInfo_inventoryType) isMultiValue() bool {
    return false
}
