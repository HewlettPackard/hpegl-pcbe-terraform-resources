package datastores
// The current status of the datastore.
type DatastoresGetResponse_items_status int

const (
    OK_DATASTORESGETRESPONSE_ITEMS_STATUS DatastoresGetResponse_items_status = iota
    ERROR_DATASTORESGETRESPONSE_ITEMS_STATUS
    WARNING_DATASTORESGETRESPONSE_ITEMS_STATUS
)

func (i DatastoresGetResponse_items_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseDatastoresGetResponse_items_status(v string) (any, error) {
    result := OK_DATASTORESGETRESPONSE_ITEMS_STATUS
    switch v {
        case "OK":
            result = OK_DATASTORESGETRESPONSE_ITEMS_STATUS
        case "ERROR":
            result = ERROR_DATASTORESGETRESPONSE_ITEMS_STATUS
        case "WARNING":
            result = WARNING_DATASTORESGETRESPONSE_ITEMS_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_status(values []DatastoresGetResponse_items_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_status) isMultiValue() bool {
    return false
}
