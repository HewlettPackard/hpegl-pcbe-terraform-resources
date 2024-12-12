package item
// The current status of the datastore.
type DatastoresGetResponse_status int

const (
    OK_DATASTORESGETRESPONSE_STATUS DatastoresGetResponse_status = iota
    ERROR_DATASTORESGETRESPONSE_STATUS
    WARNING_DATASTORESGETRESPONSE_STATUS
)

func (i DatastoresGetResponse_status) String() string {
    return []string{"OK", "ERROR", "WARNING"}[i]
}
func ParseDatastoresGetResponse_status(v string) (any, error) {
    result := OK_DATASTORESGETRESPONSE_STATUS
    switch v {
        case "OK":
            result = OK_DATASTORESGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_DATASTORESGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_DATASTORESGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_status(values []DatastoresGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_status) isMultiValue() bool {
    return false
}
