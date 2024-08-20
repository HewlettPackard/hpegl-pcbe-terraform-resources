package tags
// The current state of the hypervisor tag.
type TagsGetResponse_items_state int

const (
    OK_TAGSGETRESPONSE_ITEMS_STATE TagsGetResponse_items_state = iota
    DELETED_TAGSGETRESPONSE_ITEMS_STATE
)

func (i TagsGetResponse_items_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseTagsGetResponse_items_state(v string) (any, error) {
    result := OK_TAGSGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_TAGSGETRESPONSE_ITEMS_STATE
        case "DELETED":
            result = DELETED_TAGSGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTagsGetResponse_items_state(values []TagsGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TagsGetResponse_items_state) isMultiValue() bool {
    return false
}
