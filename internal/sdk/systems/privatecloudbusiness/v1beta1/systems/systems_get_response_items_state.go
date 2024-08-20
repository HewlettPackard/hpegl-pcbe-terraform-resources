package systems
// Current state of the system
type SystemsGetResponse_items_state int

const (
    OFFLINE_SYSTEMSGETRESPONSE_ITEMS_STATE SystemsGetResponse_items_state = iota
    ONLINE_SYSTEMSGETRESPONSE_ITEMS_STATE
    INITIALIZED_SYSTEMSGETRESPONSE_ITEMS_STATE
    UNINITIALIZED_SYSTEMSGETRESPONSE_ITEMS_STATE
)

func (i SystemsGetResponse_items_state) String() string {
    return []string{"OFFLINE", "ONLINE", "INITIALIZED", "UNINITIALIZED"}[i]
}
func ParseSystemsGetResponse_items_state(v string) (any, error) {
    result := OFFLINE_SYSTEMSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OFFLINE":
            result = OFFLINE_SYSTEMSGETRESPONSE_ITEMS_STATE
        case "ONLINE":
            result = ONLINE_SYSTEMSGETRESPONSE_ITEMS_STATE
        case "INITIALIZED":
            result = INITIALIZED_SYSTEMSGETRESPONSE_ITEMS_STATE
        case "UNINITIALIZED":
            result = UNINITIALIZED_SYSTEMSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_state(values []SystemsGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_state) isMultiValue() bool {
    return false
}
