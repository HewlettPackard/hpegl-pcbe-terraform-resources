package item
// Indicates the level of protection provided for the asset, based on the existence ofprotection jobs, whether the protection jobs are suspended, the existence of backups,and the success of recent backup attempts.
type CspRdsInstancesGetResponse_protectionStatus int

const (
    UNPROTECTED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS CspRdsInstancesGetResponse_protectionStatus = iota
    LAPSED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PENDING_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PARTIAL_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PROTECTED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
    PAUSED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
)

func (i CspRdsInstancesGetResponse_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED"}[i]
}
func ParseCspRdsInstancesGetResponse_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_protectionStatus(values []CspRdsInstancesGetResponse_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_protectionStatus) isMultiValue() bool {
    return false
}
