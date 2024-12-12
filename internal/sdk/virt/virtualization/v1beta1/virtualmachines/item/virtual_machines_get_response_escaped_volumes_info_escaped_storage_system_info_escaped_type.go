package item
// Type of storage system.
type VirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type int

const (
    NIMBLE_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE VirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type = iota
    THREEPAR_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    PRIMERA_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    ALLETRA_6000_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    ALLETRA_9000_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
)

func (i VirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type) String() string {
    return []string{"NIMBLE", "THREEPAR", "PRIMERA", "ALLETRA_6000", "ALLETRA_9000"}[i]
}
func ParseVirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type(v string) (any, error) {
    result := NIMBLE_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
    switch v {
        case "NIMBLE":
            result = NIMBLE_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "THREEPAR":
            result = THREEPAR_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "PRIMERA":
            result = PRIMERA_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "ALLETRA_6000":
            result = ALLETRA_6000_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        case "ALLETRA_9000":
            result = ALLETRA_9000_VIRTUALMACHINESGETRESPONSE_VOLUMESINFO_STORAGESYSTEMINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type(values []VirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_volumesInfo_storageSystemInfo_type) isMultiValue() bool {
    return false
}
