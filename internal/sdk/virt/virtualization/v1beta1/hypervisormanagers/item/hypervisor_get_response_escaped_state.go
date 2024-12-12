package item
// The current state of the hypervisor manager object
type HypervisorGetResponse_state int

const (
    OK_HYPERVISORGETRESPONSE_STATE HypervisorGetResponse_state = iota
    ERROR_HYPERVISORGETRESPONSE_STATE
    INITIALIZING_HYPERVISORGETRESPONSE_STATE
    CREATING_HYPERVISORGETRESPONSE_STATE
    DELETING_HYPERVISORGETRESPONSE_STATE
    UPDATING_HYPERVISORGETRESPONSE_STATE
    REFRESHING_HYPERVISORGETRESPONSE_STATE
    CONFIGURED_HYPERVISORGETRESPONSE_STATE
)

func (i HypervisorGetResponse_state) String() string {
    return []string{"OK", "ERROR", "INITIALIZING", "CREATING", "DELETING", "UPDATING", "REFRESHING", "CONFIGURED"}[i]
}
func ParseHypervisorGetResponse_state(v string) (any, error) {
    result := OK_HYPERVISORGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_HYPERVISORGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_HYPERVISORGETRESPONSE_STATE
        case "INITIALIZING":
            result = INITIALIZING_HYPERVISORGETRESPONSE_STATE
        case "CREATING":
            result = CREATING_HYPERVISORGETRESPONSE_STATE
        case "DELETING":
            result = DELETING_HYPERVISORGETRESPONSE_STATE
        case "UPDATING":
            result = UPDATING_HYPERVISORGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_HYPERVISORGETRESPONSE_STATE
        case "CONFIGURED":
            result = CONFIGURED_HYPERVISORGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorGetResponse_state(values []HypervisorGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorGetResponse_state) isMultiValue() bool {
    return false
}
