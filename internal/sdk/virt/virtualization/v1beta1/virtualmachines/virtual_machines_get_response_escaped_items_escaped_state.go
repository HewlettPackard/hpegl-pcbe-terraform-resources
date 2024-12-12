package virtualmachines
// The current state of the virtual machine
type VirtualMachinesGetResponse_items_state int

const (
    OK_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE VirtualMachinesGetResponse_items_state = iota
    UNAVAILABLE_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    ERROR_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    CREATING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    DELETING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    UPDATING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    REFRESHING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    RESTORING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    RESTORE_FAILED_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    DELETED_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
)

func (i VirtualMachinesGetResponse_items_state) String() string {
    return []string{"OK", "UNAVAILABLE", "ERROR", "CREATING", "DELETING", "UPDATING", "REFRESHING", "RESTORING", "RESTORE_FAILED", "DELETED"}[i]
}
func ParseVirtualMachinesGetResponse_items_state(v string) (any, error) {
    result := OK_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "UNAVAILABLE":
            result = UNAVAILABLE_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "CREATING":
            result = CREATING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "DELETING":
            result = DELETING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "UPDATING":
            result = UPDATING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "RESTORING":
            result = RESTORING_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "RESTORE_FAILED":
            result = RESTORE_FAILED_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        case "DELETED":
            result = DELETED_VIRTUALMACHINESGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_state(values []VirtualMachinesGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_state) isMultiValue() bool {
    return false
}
