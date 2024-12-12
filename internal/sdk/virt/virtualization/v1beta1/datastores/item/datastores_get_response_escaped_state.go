package item
// The current state of the datastore
type DatastoresGetResponse_state int

const (
    OK_DATASTORESGETRESPONSE_STATE DatastoresGetResponse_state = iota
    ERROR_DATASTORESGETRESPONSE_STATE
    CREATING_DATASTORESGETRESPONSE_STATE
    DELETING_DATASTORESGETRESPONSE_STATE
    UPDATING_DATASTORESGETRESPONSE_STATE
    REFRESHING_DATASTORESGETRESPONSE_STATE
    RESTORING_DATASTORESGETRESPONSE_STATE
    MOUNTED_DATASTORESGETRESPONSE_STATE
    DELETED_DATASTORESGETRESPONSE_STATE
)

func (i DatastoresGetResponse_state) String() string {
    return []string{"OK", "ERROR", "CREATING", "DELETING", "UPDATING", "REFRESHING", "RESTORING", "MOUNTED", "DELETED"}[i]
}
func ParseDatastoresGetResponse_state(v string) (any, error) {
    result := OK_DATASTORESGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_DATASTORESGETRESPONSE_STATE
        case "ERROR":
            result = ERROR_DATASTORESGETRESPONSE_STATE
        case "CREATING":
            result = CREATING_DATASTORESGETRESPONSE_STATE
        case "DELETING":
            result = DELETING_DATASTORESGETRESPONSE_STATE
        case "UPDATING":
            result = UPDATING_DATASTORESGETRESPONSE_STATE
        case "REFRESHING":
            result = REFRESHING_DATASTORESGETRESPONSE_STATE
        case "RESTORING":
            result = RESTORING_DATASTORESGETRESPONSE_STATE
        case "MOUNTED":
            result = MOUNTED_DATASTORESGETRESPONSE_STATE
        case "DELETED":
            result = DELETED_DATASTORESGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_state(values []DatastoresGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_state) isMultiValue() bool {
    return false
}
