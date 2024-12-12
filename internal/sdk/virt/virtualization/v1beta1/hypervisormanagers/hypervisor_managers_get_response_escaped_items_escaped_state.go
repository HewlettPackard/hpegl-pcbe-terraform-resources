package hypervisormanagers
// The current state of the hypervisor manager object
type HypervisorManagersGetResponse_items_state int

const (
    OK_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE HypervisorManagersGetResponse_items_state = iota
    ERROR_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    INITIALIZING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    CREATING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    DELETING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    UPDATING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    REFRESHING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    CONFIGURED_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
)

func (i HypervisorManagersGetResponse_items_state) String() string {
    return []string{"OK", "ERROR", "INITIALIZING", "CREATING", "DELETING", "UPDATING", "REFRESHING", "CONFIGURED"}[i]
}
func ParseHypervisorManagersGetResponse_items_state(v string) (any, error) {
    result := OK_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "INITIALIZING":
            result = INITIALIZING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "CREATING":
            result = CREATING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "DELETING":
            result = DELETING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "UPDATING":
            result = UPDATING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        case "CONFIGURED":
            result = CONFIGURED_HYPERVISORMANAGERSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorManagersGetResponse_items_state(values []HypervisorManagersGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorManagersGetResponse_items_state) isMultiValue() bool {
    return false
}
