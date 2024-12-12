package datastores
// Type of the datastore.
type DatastoresGetResponse_items_datastoreType int

const (
    VMFS_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE DatastoresGetResponse_items_datastoreType = iota
    VVOL_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
    NFS_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
    VSAN_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
)

func (i DatastoresGetResponse_items_datastoreType) String() string {
    return []string{"VMFS", "VVOL", "NFS", "VSAN"}[i]
}
func ParseDatastoresGetResponse_items_datastoreType(v string) (any, error) {
    result := VMFS_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
    switch v {
        case "VMFS":
            result = VMFS_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
        case "VVOL":
            result = VVOL_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
        case "NFS":
            result = NFS_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
        case "VSAN":
            result = VSAN_DATASTORESGETRESPONSE_ITEMS_DATASTORETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_datastoreType(values []DatastoresGetResponse_items_datastoreType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_datastoreType) isMultiValue() bool {
    return false
}
