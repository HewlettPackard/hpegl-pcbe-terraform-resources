package virtualmachines
// Provides the current protection status of this resource. - UNPROTECTED - No policy assigned, No recovery points exists - LAPSED      - No policy assigned, at least one recovery points exists - PENDING     - Policy assigned, No recovery points exists - PARTIAL     - Policy assigned, At least one recovery point exists - PROTECTED   - Policy assigned, most recent run of every configured schedule is successful - PAUSED      - Policy assigned, one or more of the schedules are paused - UNSUPPORTED - No policy can be assigned
type VirtualMachinesGetResponse_items_protectionStatus int

const (
    UNPROTECTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS VirtualMachinesGetResponse_items_protectionStatus = iota
    LAPSED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PENDING_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PARTIAL_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PROTECTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PAUSED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    UNSUPPORTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
)

func (i VirtualMachinesGetResponse_items_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED", "UNSUPPORTED"}[i]
}
func ParseVirtualMachinesGetResponse_items_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "UNSUPPORTED":
            result = UNSUPPORTED_VIRTUALMACHINESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_items_protectionStatus(values []VirtualMachinesGetResponse_items_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_items_protectionStatus) isMultiValue() bool {
    return false
}
