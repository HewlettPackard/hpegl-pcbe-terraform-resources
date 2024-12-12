package hypervisorhosts
// The type of the parent.
type HypervisorHostsGetResponse_items_parentInfo_type int

const (
    CLUSTER_HYPERVISORHOSTSGETRESPONSE_ITEMS_PARENTINFO_TYPE HypervisorHostsGetResponse_items_parentInfo_type = iota
    FOLDER_HYPERVISORHOSTSGETRESPONSE_ITEMS_PARENTINFO_TYPE
)

func (i HypervisorHostsGetResponse_items_parentInfo_type) String() string {
    return []string{"CLUSTER", "FOLDER"}[i]
}
func ParseHypervisorHostsGetResponse_items_parentInfo_type(v string) (any, error) {
    result := CLUSTER_HYPERVISORHOSTSGETRESPONSE_ITEMS_PARENTINFO_TYPE
    switch v {
        case "CLUSTER":
            result = CLUSTER_HYPERVISORHOSTSGETRESPONSE_ITEMS_PARENTINFO_TYPE
        case "FOLDER":
            result = FOLDER_HYPERVISORHOSTSGETRESPONSE_ITEMS_PARENTINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorHostsGetResponse_items_parentInfo_type(values []HypervisorHostsGetResponse_items_parentInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorHostsGetResponse_items_parentInfo_type) isMultiValue() bool {
    return false
}
