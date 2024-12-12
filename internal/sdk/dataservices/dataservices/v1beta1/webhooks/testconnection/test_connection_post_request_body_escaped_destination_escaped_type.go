package testconnection
// An enum indicating the type of endpoint that events are published to.- `CLOUD_EVENT`: A customer controlled HTTPS endpoint to which DSCC cloud events are posted.
type TestConnectionPostRequestBody_destination_type int

const (
    CLOUD_EVENT_TESTCONNECTIONPOSTREQUESTBODY_DESTINATION_TYPE TestConnectionPostRequestBody_destination_type = iota
)

func (i TestConnectionPostRequestBody_destination_type) String() string {
    return []string{"CLOUD_EVENT"}[i]
}
func ParseTestConnectionPostRequestBody_destination_type(v string) (any, error) {
    result := CLOUD_EVENT_TESTCONNECTIONPOSTREQUESTBODY_DESTINATION_TYPE
    switch v {
        case "CLOUD_EVENT":
            result = CLOUD_EVENT_TESTCONNECTIONPOSTREQUESTBODY_DESTINATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTestConnectionPostRequestBody_destination_type(values []TestConnectionPostRequestBody_destination_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TestConnectionPostRequestBody_destination_type) isMultiValue() bool {
    return false
}
