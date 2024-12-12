package datastores
// The current state of the datastore
type DatastoresGetResponse_items_state int

const (
    OK_DATASTORESGETRESPONSE_ITEMS_STATE DatastoresGetResponse_items_state = iota
    ERROR_DATASTORESGETRESPONSE_ITEMS_STATE
    CREATING_DATASTORESGETRESPONSE_ITEMS_STATE
    DELETING_DATASTORESGETRESPONSE_ITEMS_STATE
    UPDATING_DATASTORESGETRESPONSE_ITEMS_STATE
    REFRESHING_DATASTORESGETRESPONSE_ITEMS_STATE
    RESTORING_DATASTORESGETRESPONSE_ITEMS_STATE
    MOUNTED_DATASTORESGETRESPONSE_ITEMS_STATE
    DELETED_DATASTORESGETRESPONSE_ITEMS_STATE
)

func (i DatastoresGetResponse_items_state) String() string {
    return []string{"OK", "ERROR", "CREATING", "DELETING", "UPDATING", "REFRESHING", "RESTORING", "MOUNTED", "DELETED"}[i]
}
func ParseDatastoresGetResponse_items_state(v string) (any, error) {
    result := OK_DATASTORESGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_DATASTORESGETRESPONSE_ITEMS_STATE
        case "ERROR":
            result = ERROR_DATASTORESGETRESPONSE_ITEMS_STATE
        case "CREATING":
            result = CREATING_DATASTORESGETRESPONSE_ITEMS_STATE
        case "DELETING":
            result = DELETING_DATASTORESGETRESPONSE_ITEMS_STATE
        case "UPDATING":
            result = UPDATING_DATASTORESGETRESPONSE_ITEMS_STATE
        case "REFRESHING":
            result = REFRESHING_DATASTORESGETRESPONSE_ITEMS_STATE
        case "RESTORING":
            result = RESTORING_DATASTORESGETRESPONSE_ITEMS_STATE
        case "MOUNTED":
            result = MOUNTED_DATASTORESGETRESPONSE_ITEMS_STATE
        case "DELETED":
            result = DELETED_DATASTORESGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_state(values []DatastoresGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_state) isMultiValue() bool {
    return false
}
