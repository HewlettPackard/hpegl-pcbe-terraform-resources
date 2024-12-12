package networks
// The current status of the hypervisor network. Status is derived and abstracted to a 'standard status' based on the status reported by the hypervisor manager.
type NetworksGetResponse_items_status int

const (
    OK_NETWORKSGETRESPONSE_ITEMS_STATUS NetworksGetResponse_items_status = iota
    ERROR_NETWORKSGETRESPONSE_ITEMS_STATUS
    WARNING_NETWORKSGETRESPONSE_ITEMS_STATUS
)

func (i NetworksGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseNetworksGetResponse_items_status(v string) (any, error) {
    result := OK_NETWORKSGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_NETWORKSGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_NETWORKSGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_NETWORKSGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworksGetResponse_items_status(values []NetworksGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworksGetResponse_items_status) isMultiValue() bool {
    return false
}
