package item
// The status of the most recent backup attempt for this schedule.  It is set toNOT_RUN when a PAUSED state no longer applies, regardless of the result of anyprior backup attempts.  May also be reset to NOT_RUN for a volume if the protectedasset is a machine instance and the volume is detached from and re-attached to themachine instance, even if it was attached in the most recent backup attempt.
type CspVolumesGetResponse_protectionJobInfo_scheduleInfo_status int

const (
    NOT_RUN_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS CspVolumesGetResponse_protectionJobInfo_scheduleInfo_status = iota
    SUCCESS_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
    FAILURE_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
    PAUSED_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
)

func (i CspVolumesGetResponse_protectionJobInfo_scheduleInfo_status) String() string {
    return []string{"NOT_RUN", "SUCCESS", "FAILURE", "PAUSED"}[i]
}
func ParseCspVolumesGetResponse_protectionJobInfo_scheduleInfo_status(v string) (any, error) {
    result := NOT_RUN_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
    switch v {
        case "NOT_RUN":
            result = NOT_RUN_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        case "SUCCESS":
            result = SUCCESS_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        case "FAILURE":
            result = FAILURE_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        case "PAUSED":
            result = PAUSED_CSPVOLUMESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVolumesGetResponse_protectionJobInfo_scheduleInfo_status(values []CspVolumesGetResponse_protectionJobInfo_scheduleInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVolumesGetResponse_protectionJobInfo_scheduleInfo_status) isMultiValue() bool {
    return false
}
