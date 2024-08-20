package item
// The status of the most recent backup attempt for this schedule.  It is set toNOT_RUN when a PAUSED state no longer applies, regardless of the result of anyprior backup attempts.
type CspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status int

const (
    NOT_RUN_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS CspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status = iota
    SUCCESS_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
    FAILURE_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
    PAUSED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
)

func (i CspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status) String() string {
    return []string{"NOT_RUN", "SUCCESS", "FAILURE", "PAUSED"}[i]
}
func ParseCspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status(v string) (any, error) {
    result := NOT_RUN_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
    switch v {
        case "NOT_RUN":
            result = NOT_RUN_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        case "SUCCESS":
            result = SUCCESS_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        case "FAILURE":
            result = FAILURE_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        case "PAUSED":
            result = PAUSED_CSPRDSINSTANCESGETRESPONSE_PROTECTIONJOBINFO_SCHEDULEINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status(values []CspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_protectionJobInfo_scheduleInfo_status) isMultiValue() bool {
    return false
}
