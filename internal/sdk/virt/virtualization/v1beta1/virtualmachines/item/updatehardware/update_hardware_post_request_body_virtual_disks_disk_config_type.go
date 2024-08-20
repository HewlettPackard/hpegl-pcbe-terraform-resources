package updatehardware
// The type of the disk.
type UpdateHardwarePostRequestBody_virtualDisks_diskConfig_type int

const (
    IDE_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE UpdateHardwarePostRequestBody_virtualDisks_diskConfig_type = iota
    SCSI_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
    SATA_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
    NVME_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
)

func (i UpdateHardwarePostRequestBody_virtualDisks_diskConfig_type) String() string {
    return []string{"IDE", "SCSI", "SATA", "NVME"}[i]
}
func ParseUpdateHardwarePostRequestBody_virtualDisks_diskConfig_type(v string) (any, error) {
    result := IDE_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
    switch v {
        case "IDE":
            result = IDE_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
        case "SCSI":
            result = SCSI_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
        case "SATA":
            result = SATA_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
        case "NVME":
            result = NVME_UPDATEHARDWAREPOSTREQUESTBODY_VIRTUALDISKS_DISKCONFIG_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUpdateHardwarePostRequestBody_virtualDisks_diskConfig_type(values []UpdateHardwarePostRequestBody_virtualDisks_diskConfig_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UpdateHardwarePostRequestBody_virtualDisks_diskConfig_type) isMultiValue() bool {
    return false
}
