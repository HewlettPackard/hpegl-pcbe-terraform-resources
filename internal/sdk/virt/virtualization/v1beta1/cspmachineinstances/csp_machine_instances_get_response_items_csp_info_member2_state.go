package cspmachineinstances
// Power state of the machine instance.
type CspMachineInstancesGetResponse_items_cspInfoMember2_state int

const (
    UNKNOWN_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE CspMachineInstancesGetResponse_items_cspInfoMember2_state = iota
    CREATING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    STARTING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    RUNNING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    STOPPING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    STOPPED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    DEALLOCATING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    DEALLOCATED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
)

func (i CspMachineInstancesGetResponse_items_cspInfoMember2_state) String() string {
    return []string{"UNKNOWN", "CREATING", "STARTING", "RUNNING", "STOPPING", "STOPPED", "DEALLOCATING", "DEALLOCATED"}[i]
}
func ParseCspMachineInstancesGetResponse_items_cspInfoMember2_state(v string) (any, error) {
    result := UNKNOWN_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "CREATING":
            result = CREATING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "STARTING":
            result = STARTING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "RUNNING":
            result = RUNNING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "STOPPING":
            result = STOPPING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "STOPPED":
            result = STOPPED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "DEALLOCATING":
            result = DEALLOCATING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        case "DEALLOCATED":
            result = DEALLOCATED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_CSPINFOMEMBER2_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_items_cspInfoMember2_state(values []CspMachineInstancesGetResponse_items_cspInfoMember2_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_items_cspInfoMember2_state) isMultiValue() bool {
    return false
}
