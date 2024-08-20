package item
// Status of the adapter
type HostGetResponse_storageAdaptersInfo_status int

const (
    OFFLINE_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS HostGetResponse_storageAdaptersInfo_status = iota
    ONLINE_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS
    UNKNOWN_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS
)

func (i HostGetResponse_storageAdaptersInfo_status) String() string {
    return []string{"OFFLINE", "ONLINE", "UNKNOWN"}[i]
}
func ParseHostGetResponse_storageAdaptersInfo_status(v string) (any, error) {
    result := OFFLINE_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS
    switch v {
        case "OFFLINE":
            result = OFFLINE_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS
        case "ONLINE":
            result = ONLINE_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS
        case "UNKNOWN":
            result = UNKNOWN_HOSTGETRESPONSE_STORAGEADAPTERSINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostGetResponse_storageAdaptersInfo_status(values []HostGetResponse_storageAdaptersInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostGetResponse_storageAdaptersInfo_status) isMultiValue() bool {
    return false
}
