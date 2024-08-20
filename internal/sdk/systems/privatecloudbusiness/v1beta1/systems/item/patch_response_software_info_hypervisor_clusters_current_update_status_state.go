package item
// Status of the software update:  * `UP_TO_DATE` - Already up to date  * `UPDATE_AVAILABLE` - One or more updates are available  * `PENDING` - Waiting for update operation (precheck or update) to begin  * `PRECHECK_IN_PROGRESS` - A software update precheck operation is in progress  * `PRECHECK_FAILED` - The previous software update precheck operation has failed  * `PRECHECK_COMPLETE` - The previous software update precheck operation has completed successfully  * `UPDATE_IN_PROGRESS` - A software update operation is in progress  * `UPDATE_COMPLETE` - The previous software update operation has completed successfully  * `UPDATE_FAILED` - The previous software update operation has failed  * `NOT_READY` - Not ready for update (e.g. when current version is not available, so no update paths exist)
type PatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state int

const (
    UP_TO_DATE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE PatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state = iota
    UPDATE_AVAILABLE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    PENDING_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    PRECHECK_IN_PROGRESS_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    PRECHECK_FAILED_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    PRECHECK_COMPLETE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    UPDATE_IN_PROGRESS_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    UPDATE_COMPLETE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    UPDATE_FAILED_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    NOT_READY_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
)

func (i PatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state) String() string {
    return []string{"UP_TO_DATE", "UPDATE_AVAILABLE", "PENDING", "PRECHECK_IN_PROGRESS", "PRECHECK_FAILED", "PRECHECK_COMPLETE", "UPDATE_IN_PROGRESS", "UPDATE_COMPLETE", "UPDATE_FAILED", "NOT_READY"}[i]
}
func ParsePatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state(v string) (any, error) {
    result := UP_TO_DATE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
    switch v {
        case "UP_TO_DATE":
            result = UP_TO_DATE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "UPDATE_AVAILABLE":
            result = UPDATE_AVAILABLE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "PENDING":
            result = PENDING_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "PRECHECK_IN_PROGRESS":
            result = PRECHECK_IN_PROGRESS_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "PRECHECK_FAILED":
            result = PRECHECK_FAILED_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "PRECHECK_COMPLETE":
            result = PRECHECK_COMPLETE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "UPDATE_IN_PROGRESS":
            result = UPDATE_IN_PROGRESS_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "UPDATE_COMPLETE":
            result = UPDATE_COMPLETE_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "UPDATE_FAILED":
            result = UPDATE_FAILED_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        case "NOT_READY":
            result = NOT_READY_PATCHRESPONSE_SOFTWAREINFO_HYPERVISORCLUSTERS_CURRENTUPDATESTATUS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state(values []PatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_softwareInfo_hypervisorClusters_currentUpdateStatus_state) isMultiValue() bool {
    return false
}
