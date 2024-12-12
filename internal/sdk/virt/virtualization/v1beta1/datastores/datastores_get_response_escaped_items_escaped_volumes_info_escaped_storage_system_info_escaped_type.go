package datastores
// Type of storage system.
type DatastoresGetResponse_items_volumesInfo_storageSystemInfo_type int

const (
    NIMBLE_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE DatastoresGetResponse_items_volumesInfo_storageSystemInfo_type = iota
    THREEPAR_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    PRIMERA_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    ALLETRA_6000_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    ALLETRA_9000_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
)

func (i DatastoresGetResponse_items_volumesInfo_storageSystemInfo_type) String() string {
    return []string{"NIMBLE", "THREEPAR", "PRIMERA", "ALLETRA_6000", "ALLETRA_9000"}[i]
}
func ParseDatastoresGetResponse_items_volumesInfo_storageSystemInfo_type(v string) (any, error) {
    result := NIMBLE_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    switch v {
        case "NIMBLE":
            result = NIMBLE_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "THREEPAR":
            result = THREEPAR_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "PRIMERA":
            result = PRIMERA_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "ALLETRA_6000":
            result = ALLETRA_6000_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "ALLETRA_9000":
            result = ALLETRA_9000_DATASTORESGETRESPONSE_ITEMS_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_volumesInfo_storageSystemInfo_type(values []DatastoresGetResponse_items_volumesInfo_storageSystemInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_volumesInfo_storageSystemInfo_type) isMultiValue() bool {
    return false
}
