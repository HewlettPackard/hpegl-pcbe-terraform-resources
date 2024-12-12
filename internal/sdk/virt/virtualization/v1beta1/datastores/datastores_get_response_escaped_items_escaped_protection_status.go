package datastores
// Provides the current protection status of this resource. - UNPROTECTED - No policy assigned, No recovery points exists - LAPSED      - No policy assigned, at least one recovery points exists - PENDING     - Policy assigned, No recovery points exists - PARTIAL     - Policy assigned, At least one recovery point exists - PROTECTED   - Policy assigned, most recent run of every configured schedule is successful - PAUSED      - Policy assigned, one or more of the schedules are paused - UNSUPPORTED - No policy can be assigned
type DatastoresGetResponse_items_protectionStatus int

const (
    UNPROTECTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS DatastoresGetResponse_items_protectionStatus = iota
    LAPSED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PENDING_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PARTIAL_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PROTECTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    PAUSED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    UNSUPPORTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
)

func (i DatastoresGetResponse_items_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED", "UNSUPPORTED"}[i]
}
func ParseDatastoresGetResponse_items_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        case "UNSUPPORTED":
            result = UNSUPPORTED_DATASTORESGETRESPONSE_ITEMS_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_protectionStatus(values []DatastoresGetResponse_items_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_protectionStatus) isMultiValue() bool {
    return false
}
