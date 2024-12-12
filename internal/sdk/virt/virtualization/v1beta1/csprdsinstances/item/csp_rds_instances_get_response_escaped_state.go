package item
// A DELETED state indicates the asset was already deleted from the cloud serviceprovider and will be deleted from the inventory when no backups orprotection policy relationships remain.
type CspRdsInstancesGetResponse_state int

const (
    OK_CSPRDSINSTANCESGETRESPONSE_STATE CspRdsInstancesGetResponse_state = iota
    DELETED_CSPRDSINSTANCESGETRESPONSE_STATE
)

func (i CspRdsInstancesGetResponse_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseCspRdsInstancesGetResponse_state(v string) (any, error) {
    result := OK_CSPRDSINSTANCESGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_CSPRDSINSTANCESGETRESPONSE_STATE
        case "DELETED":
            result = DELETED_CSPRDSINSTANCESGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_state(values []CspRdsInstancesGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_state) isMultiValue() bool {
    return false
}
