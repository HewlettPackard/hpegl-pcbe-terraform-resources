package item
// A type of backup available for the asset.
type CspRdsInstancesGetResponse_backupInfo_type int

const (
    NATIVE_BACKUP_CSPRDSINSTANCESGETRESPONSE_BACKUPINFO_TYPE CspRdsInstancesGetResponse_backupInfo_type = iota
)

func (i CspRdsInstancesGetResponse_backupInfo_type) String() string {
    return []string{"NATIVE_BACKUP"}[i]
}
func ParseCspRdsInstancesGetResponse_backupInfo_type(v string) (any, error) {
    result := NATIVE_BACKUP_CSPRDSINSTANCESGETRESPONSE_BACKUPINFO_TYPE
    switch v {
        case "NATIVE_BACKUP":
            result = NATIVE_BACKUP_CSPRDSINSTANCESGETRESPONSE_BACKUPINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_backupInfo_type(values []CspRdsInstancesGetResponse_backupInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_backupInfo_type) isMultiValue() bool {
    return false
}
