package cspmachineinstances
// A type of backup available for the asset.
type CspMachineInstancesGetResponse_items_backupInfo_backupType int

const (
    NATIVE_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE CspMachineInstancesGetResponse_items_backupInfo_backupType = iota
    HPE_CLOUD_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE
    STAGING_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE
)

func (i CspMachineInstancesGetResponse_items_backupInfo_backupType) String() string {
    return []string{"NATIVE_BACKUP", "HPE_CLOUD_BACKUP", "STAGING_BACKUP"}[i]
}
func ParseCspMachineInstancesGetResponse_items_backupInfo_backupType(v string) (any, error) {
    result := NATIVE_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE
    switch v {
        case "NATIVE_BACKUP":
            result = NATIVE_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE
        case "HPE_CLOUD_BACKUP":
            result = HPE_CLOUD_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE
        case "STAGING_BACKUP":
            result = STAGING_BACKUP_CSPMACHINEINSTANCESGETRESPONSE_ITEMS_BACKUPINFO_BACKUPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_items_backupInfo_backupType(values []CspMachineInstancesGetResponse_items_backupInfo_backupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_items_backupInfo_backupType) isMultiValue() bool {
    return false
}
