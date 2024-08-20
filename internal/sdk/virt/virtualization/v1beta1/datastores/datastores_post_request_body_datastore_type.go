package datastores
// Supported datastore types are VMFS or vVOL
type DatastoresPostRequestBody_datastoreType int

const (
    VMFS_DATASTORESPOSTREQUESTBODY_DATASTORETYPE DatastoresPostRequestBody_datastoreType = iota
    VVOL_DATASTORESPOSTREQUESTBODY_DATASTORETYPE
)

func (i DatastoresPostRequestBody_datastoreType) String() string {
    return []string{"VMFS", "VVOL"}[i]
}
func ParseDatastoresPostRequestBody_datastoreType(v string) (any, error) {
    result := VMFS_DATASTORESPOSTREQUESTBODY_DATASTORETYPE
    switch v {
        case "VMFS":
            result = VMFS_DATASTORESPOSTREQUESTBODY_DATASTORETYPE
        case "VVOL":
            result = VVOL_DATASTORESPOSTREQUESTBODY_DATASTORETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresPostRequestBody_datastoreType(values []DatastoresPostRequestBody_datastoreType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresPostRequestBody_datastoreType) isMultiValue() bool {
    return false
}
