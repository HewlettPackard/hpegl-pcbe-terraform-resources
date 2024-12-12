package csprdsinstances
// Indicates the level of protection provided for the asset, based on the existence ofprotection jobs, whether the protection jobs are suspended, the existence of backups,and the success of recent backup attempts.
type CspRdsInstancesGetResponse_items_protectionStatus int

const (
    UNPROTECTED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS CspRdsInstancesGetResponse_items_protectionStatus = iota
    LAPSED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PENDING_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PARTIAL_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PROTECTED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PAUSED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
)

func (i CspRdsInstancesGetResponse_items_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED"}[i]
}
func ParseCspRdsInstancesGetResponse_items_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_CSPRDSINSTANCESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_items_protectionStatus(values []CspRdsInstancesGetResponse_items_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_items_protectionStatus) isMultiValue() bool {
    return false
}
