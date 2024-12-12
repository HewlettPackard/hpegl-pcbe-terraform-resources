package item
// Type of the datastore.
type DatastoresGetResponse_datastoreType int

const (
    VMFS_DATASTORESGETRESPONSE_DATASTORETYPE DatastoresGetResponse_datastoreType = iota
    VVOL_DATASTORESGETRESPONSE_DATASTORETYPE
    NFS_DATASTORESGETRESPONSE_DATASTORETYPE
    VSAN_DATASTORESGETRESPONSE_DATASTORETYPE
)

func (i DatastoresGetResponse_datastoreType) String() string {
    return []string{"VMFS", "VVOL", "NFS", "VSAN"}[i]
}
func ParseDatastoresGetResponse_datastoreType(v string) (any, error) {
    result := VMFS_DATASTORESGETRESPONSE_DATASTORETYPE
    switch v {
        case "VMFS":
            result = VMFS_DATASTORESGETRESPONSE_DATASTORETYPE
        case "VVOL":
            result = VVOL_DATASTORESGETRESPONSE_DATASTORETYPE
        case "NFS":
            result = NFS_DATASTORESGETRESPONSE_DATASTORETYPE
        case "VSAN":
            result = VSAN_DATASTORESGETRESPONSE_DATASTORETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_datastoreType(values []DatastoresGetResponse_datastoreType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_datastoreType) isMultiValue() bool {
    return false
}
