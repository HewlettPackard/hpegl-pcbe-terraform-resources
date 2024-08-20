package hypervisorhosts
// Status of the adapter
type HypervisorHostsGetResponse_items_storageAdaptersInfo_status int

const (
    OFFLINE_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS HypervisorHostsGetResponse_items_storageAdaptersInfo_status = iota
    ONLINE_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS
    UNKNOWN_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS
)

func (i HypervisorHostsGetResponse_items_storageAdaptersInfo_status) String() string {
    return []string{"OFFLINE", "ONLINE", "UNKNOWN"}[i]
}
func ParseHypervisorHostsGetResponse_items_storageAdaptersInfo_status(v string) (any, error) {
    result := OFFLINE_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS
    switch v {
        case "OFFLINE":
            result = OFFLINE_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS
        case "ONLINE":
            result = ONLINE_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS
        case "UNKNOWN":
            result = UNKNOWN_HYPERVISORHOSTSGETRESPONSE_ITEMS_STORAGEADAPTERSINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorHostsGetResponse_items_storageAdaptersInfo_status(values []HypervisorHostsGetResponse_items_storageAdaptersInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorHostsGetResponse_items_storageAdaptersInfo_status) isMultiValue() bool {
    return false
}
