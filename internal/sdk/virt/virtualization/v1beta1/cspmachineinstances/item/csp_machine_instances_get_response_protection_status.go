package item
// Indicates the level of protection provided for the asset, based on the existence ofprotection jobs, whether the protection jobs are suspended, the existence of backups,and the success of recent backup attempts.
type CspMachineInstancesGetResponse_protectionStatus int

const (
    UNPROTECTED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS CspMachineInstancesGetResponse_protectionStatus = iota
    LAPSED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PENDING_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PARTIAL_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PROTECTED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PAUSED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
)

func (i CspMachineInstancesGetResponse_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED"}[i]
}
func ParseCspMachineInstancesGetResponse_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_CSPMACHINEINSTANCESGETRESPONSE_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_protectionStatus(values []CspMachineInstancesGetResponse_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_protectionStatus) isMultiValue() bool {
    return false
}
