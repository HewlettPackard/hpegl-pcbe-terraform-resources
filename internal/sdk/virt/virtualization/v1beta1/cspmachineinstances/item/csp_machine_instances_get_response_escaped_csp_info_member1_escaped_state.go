package item
type CspMachineInstancesGetResponse_cspInfoMember1_state int

const (
    UNKNOWN_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE CspMachineInstancesGetResponse_cspInfoMember1_state = iota
    PENDING_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
    RUNNING_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
    SHUTTING_DOWN_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
    STOPPING_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
    STOPPED_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
    TERMINATED_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
)

func (i CspMachineInstancesGetResponse_cspInfoMember1_state) String() string {
    return []string{"UNKNOWN", "PENDING", "RUNNING", "SHUTTING_DOWN", "STOPPING", "STOPPED", "TERMINATED"}[i]
}
func ParseCspMachineInstancesGetResponse_cspInfoMember1_state(v string) (any, error) {
    result := UNKNOWN_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        case "PENDING":
            result = PENDING_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        case "RUNNING":
            result = RUNNING_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        case "SHUTTING_DOWN":
            result = SHUTTING_DOWN_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        case "STOPPING":
            result = STOPPING_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        case "STOPPED":
            result = STOPPED_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        case "TERMINATED":
            result = TERMINATED_CSPMACHINEINSTANCESGETRESPONSE_CSPINFOMEMBER1_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_cspInfoMember1_state(values []CspMachineInstancesGetResponse_cspInfoMember1_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_cspInfoMember1_state) isMultiValue() bool {
    return false
}
