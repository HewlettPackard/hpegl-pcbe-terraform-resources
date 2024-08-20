package item
// Provides the current protection status of this resource. - UNPROTECTED - No policy assigned, No recovery points exists - LAPSED      - No policy assigned, at least one recovery points exists - PENDING     - Policy assigned, No recovery points exists - PARTIAL     - Policy assigned, At least one recovery point exists - PROTECTED   - Policy assigned, most recent run of every configured schedule is successful - PAUSED      - Policy assigned, one or more of the schedules are paused - UNSUPPORTED - No policy can be assigned
type DatastoresGetResponse_protectionStatus int

const (
    UNPROTECTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS DatastoresGetResponse_protectionStatus = iota
    LAPSED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
    PENDING_DATASTORESGETRESPONSE_PROTECTIONSTATUS
    PARTIAL_DATASTORESGETRESPONSE_PROTECTIONSTATUS
    PROTECTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
    PAUSED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
    UNSUPPORTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
)

func (i DatastoresGetResponse_protectionStatus) String() string {
    return []string{"UNPROTECTED", "LAPSED", "PENDING", "PARTIAL", "PROTECTED", "PAUSED", "UNSUPPORTED"}[i]
}
func ParseDatastoresGetResponse_protectionStatus(v string) (any, error) {
    result := UNPROTECTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
    switch v {
        case "UNPROTECTED":
            result = UNPROTECTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        case "LAPSED":
            result = LAPSED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        case "PENDING":
            result = PENDING_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        case "PARTIAL":
            result = PARTIAL_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        case "PROTECTED":
            result = PROTECTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        case "PAUSED":
            result = PAUSED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        case "UNSUPPORTED":
            result = UNSUPPORTED_DATASTORESGETRESPONSE_PROTECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_protectionStatus(values []DatastoresGetResponse_protectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_protectionStatus) isMultiValue() bool {
    return false
}
