package item
// The current state of the hypervisor tag.
type TagGetResponse_state int

const (
    OK_TAGGETRESPONSE_STATE TagGetResponse_state = iota
    DELETED_TAGGETRESPONSE_STATE
)

func (i TagGetResponse_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseTagGetResponse_state(v string) (any, error) {
    result := OK_TAGGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_TAGGETRESPONSE_STATE
        case "DELETED":
            result = DELETED_TAGGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTagGetResponse_state(values []TagGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TagGetResponse_state) isMultiValue() bool {
    return false
}
