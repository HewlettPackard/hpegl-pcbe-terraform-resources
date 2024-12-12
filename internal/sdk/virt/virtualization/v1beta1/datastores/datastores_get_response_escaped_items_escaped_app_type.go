package datastores
// Application type of this asset
type DatastoresGetResponse_items_appType int

const (
    VMWARE_DATASTORESGETRESPONSE_ITEMS_APPTYPE DatastoresGetResponse_items_appType = iota
    HPE_VM_DATASTORESGETRESPONSE_ITEMS_APPTYPE
)

func (i DatastoresGetResponse_items_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseDatastoresGetResponse_items_appType(v string) (any, error) {
    result := VMWARE_DATASTORESGETRESPONSE_ITEMS_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_DATASTORESGETRESPONSE_ITEMS_APPTYPE
        case "HPE_VM":
            result = HPE_VM_DATASTORESGETRESPONSE_ITEMS_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_appType(values []DatastoresGetResponse_items_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_appType) isMultiValue() bool {
    return false
}
