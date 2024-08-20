package item
// The current state of the virtual machine
type VirtualMachinesGetResponse_state int

const (
    OK_VIRTUALMACHINESGETRESPONSE_STATE VirtualMachinesGetResponse_state = iota
    UNAVAILABLE_VIRTUALMACHINESGETRESPONSE_STATE
    ERROR_VIRTUALMACHINESGETRESPONSE_STATE
    CREATING_VIRTUALMACHINESGETRESPONSE_STATE
    DELETING_VIRTUALMACHINESGETRESPONSE_STATE
    UPDATING_VIRTUALMACHINESGETRESPONSE_STATE
    REFRESHING_VIRTUALMACHINESGETRESPONSE_STATE
    RESTORING_VIRTUALMACHINESGETRESPONSE_STATE
    RESTORE_FAILED_VIRTUALMACHINESGETRESPONSE_STATE
    DELETED_VIRTUALMACHINESGETRESPONSE_STATE
)

func (i VirtualMachinesGetResponse_state) String() string {
    return []string{"OK", "UNAVAILABLE", "ERROR", "CREATING", "DELETING", "UPDATING", "REFRESHING", "RESTORING", "RESTORE_FAILED", "DELETED"}[i]
}
func ParseVirtualMachinesGetResponse_state(v string) (any, error) {
    result := OK_VIRTUALMACHINESGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_VIRTUALMACHINESGETRESPONSE_STATE
        case "UNAVAILABLE":
            result = UNAVAILABLE_VIRTUALMACHINESGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_VIRTUALMACHINESGETRESPONSE_STATE
        case "CREATING":
            result = CREATING_VIRTUALMACHINESGETRESPONSE_STATE
        case "DELETING":
            result = DELETING_VIRTUALMACHINESGETRESPONSE_STATE
        case "UPDATING":
            result = UPDATING_VIRTUALMACHINESGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_VIRTUALMACHINESGETRESPONSE_STATE
        case "RESTORING":
            result = RESTORING_VIRTUALMACHINESGETRESPONSE_STATE
        case "RESTORE_FAILED":
            result = RESTORE_FAILED_VIRTUALMACHINESGETRESPONSE_STATE
        case "DELETED":
            result = DELETED_VIRTUALMACHINESGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_state(values []VirtualMachinesGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_state) isMultiValue() bool {
    return false
}
