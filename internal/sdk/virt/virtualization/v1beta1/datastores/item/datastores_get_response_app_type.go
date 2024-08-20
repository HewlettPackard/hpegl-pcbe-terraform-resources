package item
// Application type of this asset
type DatastoresGetResponse_appType int

const (
    VMWARE_DATASTORESGETRESPONSE_APPTYPE DatastoresGetResponse_appType = iota
    HPE_VM_DATASTORESGETRESPONSE_APPTYPE
)

func (i DatastoresGetResponse_appType) String() string {
    return []string{"VMWARE", "HPE_VM"}[i]
}
func ParseDatastoresGetResponse_appType(v string) (any, error) {
    result := VMWARE_DATASTORESGETRESPONSE_APPTYPE
    switch v {
        case "VMWARE":
            result = VMWARE_DATASTORESGETRESPONSE_APPTYPE
        case "HPE_VM":
            result = HPE_VM_DATASTORESGETRESPONSE_APPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_appType(values []DatastoresGetResponse_appType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_appType) isMultiValue() bool {
    return false
}
