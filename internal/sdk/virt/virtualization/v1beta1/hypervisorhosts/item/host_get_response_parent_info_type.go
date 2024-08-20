package item
// The type of the parent.
type HostGetResponse_parentInfo_type int

const (
    CLUSTER_HOSTGETRESPONSE_PARENTINFO_TYPE HostGetResponse_parentInfo_type = iota
    FOLDER_HOSTGETRESPONSE_PARENTINFO_TYPE
)

func (i HostGetResponse_parentInfo_type) String() string {
    return []string{"CLUSTER", "FOLDER"}[i]
}
func ParseHostGetResponse_parentInfo_type(v string) (any, error) {
    result := CLUSTER_HOSTGETRESPONSE_PARENTINFO_TYPE
    switch v {
        case "CLUSTER":
            result = CLUSTER_HOSTGETRESPONSE_PARENTINFO_TYPE
        case "FOLDER":
            result = FOLDER_HOSTGETRESPONSE_PARENTINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostGetResponse_parentInfo_type(values []HostGetResponse_parentInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostGetResponse_parentInfo_type) isMultiValue() bool {
    return false
}
