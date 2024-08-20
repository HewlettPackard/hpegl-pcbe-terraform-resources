package item
// Current state of the system
type PatchResponse_state int

const (
    OFFLINE_PATCHRESPONSE_STATE PatchResponse_state = iota
    ONLINE_PATCHRESPONSE_STATE
    INITIALIZED_PATCHRESPONSE_STATE
    UNINITIALIZED_PATCHRESPONSE_STATE
)

func (i PatchResponse_state) String() string {
    return []string{"OFFLINE", "ONLINE", "INITIALIZED", "UNINITIALIZED"}[i]
}
func ParsePatchResponse_state(v string) (any, error) {
    result := OFFLINE_PATCHRESPONSE_STATE
    switch v {
        case "OFFLINE":
            result = OFFLINE_PATCHRESPONSE_STATE
        case "ONLINE":
            result = ONLINE_PATCHRESPONSE_STATE
        case "INITIALIZED":
            result = INITIALIZED_PATCHRESPONSE_STATE
        case "UNINITIALIZED":
            result = UNINITIALIZED_PATCHRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_state(values []PatchResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_state) isMultiValue() bool {
    return false
}
