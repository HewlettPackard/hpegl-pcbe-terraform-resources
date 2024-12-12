package cspmachineinstances
// Indicates the level of protection provided for the asset, based on the existence ofprotection jobs, whether the protection jobs are suspended, the existence of backups,and the success of recent backup attempts.
type CspMachineInstancesGetResponse_items_protectionStatus int

const (
    UNPROTECTED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS CspMachineInstancesGetResponse_items_protectionStatus = iota
    LAPSED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PENDING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PARTIAL_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PROTECTED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PAUSED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
)

func (i CspMachineInstancesGetResponse_items_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED"}[i]
}
func ParseCspMachineInstancesGetResponse_items_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_items_protectionStatus(values []CspMachineInstancesGetResponse_items_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_items_protectionStatus) isMultiValue() bool {
    return false
}
