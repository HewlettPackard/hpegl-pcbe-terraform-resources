package networks
// The current state of the hypervisor network.
type NetworksGetResponse_items_state int

const (
    OK_NETWORKSGETRESPONSE_ITEMS_STATE NetworksGetResponse_items_state = iota
    ERROR_NETWORKSGETRESPONSE_ITEMS_STATE
    REFRESHING_NETWORKSGETRESPONSE_ITEMS_STATE
)

func (i NetworksGetResponse_items_state) String() string {
    return []string{"OK", "ERROR", "REFRESHING"}[i]
}
func ParseNetworksGetResponse_items_state(v string) (any, error) {
    result := OK_NETWORKSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_NETWORKSGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_NETWORKSGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_NETWORKSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworksGetResponse_items_state(values []NetworksGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworksGetResponse_items_state) isMultiValue() bool {
    return false
}
