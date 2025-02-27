package item
// Status of the software update:  * `UP_TO_DATE` - Already up to date  * `UPDATE_AVAILABLE` - One or more updates are available  * `PENDING` - Waiting for update operation (precheck or update) to begin  * `PRECHECK_IN_PROGRESS` - A software update precheck operation is in progress  * `PRECHECK_FAILED` - The previous software update precheck operation has failed  * `PRECHECK_COMPLETE` - The previous software update precheck operation has completed successfully  * `DOWNLOAD_IN_PROGRESS` - A software download operation is in progress  * `DOWNLOAD_FAILED` - The previous software download operation has failed  * `DOWNLOAD_COMPLETE` - The previous software download operation has completed successfully  * `UPDATE_IN_PROGRESS` - A software update operation is in progress  * `UPDATE_COMPLETE` - The previous software update operation has completed successfully  * `UPDATE_FAILED` - The previous software update operation has failed  * `NOT_READY` - Not ready for update (e.g. when current version is not available, so no update paths exist)
type GetResponse_softwareInfo_currentUpdateStatus_state int

const (
    UP_TO_DATE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE GetResponse_softwareInfo_currentUpdateStatus_state = iota
    UPDATE_AVAILABLE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    PENDING_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    PRECHECK_IN_PROGRESS_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    PRECHECK_FAILED_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    PRECHECK_COMPLETE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    DOWNLOAD_IN_PROGRESS_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    DOWNLOAD_FAILED_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    DOWNLOAD_COMPLETE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    UPDATE_IN_PROGRESS_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    UPDATE_COMPLETE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    UPDATE_FAILED_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    NOT_READY_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
)

func (i GetResponse_softwareInfo_currentUpdateStatus_state) String() string {
    return []string{"UP_TO_DATE", "UPDATE_AVAILABLE", "PENDING", "PRECHECK_IN_PROGRESS", "PRECHECK_FAILED", "PRECHECK_COMPLETE", "DOWNLOAD_IN_PROGRESS", "DOWNLOAD_FAILED", "DOWNLOAD_COMPLETE", "UPDATE_IN_PROGRESS", "UPDATE_COMPLETE", "UPDATE_FAILED", "NOT_READY"}[i]
}
func ParseGetResponse_softwareInfo_currentUpdateStatus_state(v string) (any, error) {
    result := UP_TO_DATE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
    switch v {
        case "UP_TO_DATE":
            result = UP_TO_DATE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "UPDATE_AVAILABLE":
            result = UPDATE_AVAILABLE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "PENDING":
            result = PENDING_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "PRECHECK_IN_PROGRESS":
            result = PRECHECK_IN_PROGRESS_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "PRECHECK_FAILED":
            result = PRECHECK_FAILED_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "PRECHECK_COMPLETE":
            result = PRECHECK_COMPLETE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "DOWNLOAD_IN_PROGRESS":
            result = DOWNLOAD_IN_PROGRESS_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "DOWNLOAD_FAILED":
            result = DOWNLOAD_FAILED_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "DOWNLOAD_COMPLETE":
            result = DOWNLOAD_COMPLETE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "UPDATE_IN_PROGRESS":
            result = UPDATE_IN_PROGRESS_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "UPDATE_COMPLETE":
            result = UPDATE_COMPLETE_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "UPDATE_FAILED":
            result = UPDATE_FAILED_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        case "NOT_READY":
            result = NOT_READY_GETRESPONSE_SOFTWAREINFO_CURRENTUPDATESTATUS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_softwareInfo_currentUpdateStatus_state(values []GetResponse_softwareInfo_currentUpdateStatus_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_softwareInfo_currentUpdateStatus_state) isMultiValue() bool {
    return false
}
