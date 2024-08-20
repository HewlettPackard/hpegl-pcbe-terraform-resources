package cspvolumes
// Indicates the level of protection provided for the asset, based on the existence ofprotection jobs, whether the protection jobs are suspended, the existence of backups,and the success of recent backup attempts.
type CspVolumesGetResponse_items_protectionStatus int

const (
    UNPROTECTED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS CspVolumesGetResponse_items_protectionStatus = iota
    LAPSED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PENDING_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PARTIAL_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PROTECTED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PAUSED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
)

func (i CspVolumesGetResponse_items_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED"}[i]
}
func ParseCspVolumesGetResponse_items_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_CSPVOLUMESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVolumesGetResponse_items_protectionStatus(values []CspVolumesGetResponse_items_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVolumesGetResponse_items_protectionStatus) isMultiValue() bool {
    return false
}
