package item
// Current state of the system
type GetResponse_state int

const (
    OFFLINE_GETRESPONSE_STATE GetResponse_state = iota
    ONLINE_GETRESPONSE_STATE
    INITIALIZED_GETRESPONSE_STATE
    UNINITIALIZED_GETRESPONSE_STATE
)

func (i GetResponse_state) String() string {
    return []string{"OFFLINE", "ONLINE", "INITIALIZED", "UNINITIALIZED"}[i]
}
func ParseGetResponse_state(v string) (any, error) {
    result := OFFLINE_GETRESPONSE_STATE
    switch v {
        case "OFFLINE":
            result = OFFLINE_GETRESPONSE_STATE
        case "ONLINE":
            result = ONLINE_GETRESPONSE_STATE
        case "INITIALIZED":
            result = INITIALIZED_GETRESPONSE_STATE
        case "UNINITIALIZED":
            result = UNINITIALIZED_GETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_state(values []GetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_state) isMultiValue() bool {
    return false
}
