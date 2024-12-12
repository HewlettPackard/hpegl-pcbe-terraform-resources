package csprdsinstances
// A DELETED state indicates the asset was already deleted from the cloud serviceprovider and will be deleted from the inventory when no backups orprotection policy relationships remain.
type CspRdsInstancesGetResponse_items_state int

const (
    OK_CSPRDSINSTANCESGETRESPONSE_ITEMS_STATE CspRdsInstancesGetResponse_items_state = iota
    DELETED_CSPRDSINSTANCESGETRESPONSE_ITEMS_STATE
)

func (i CspRdsInstancesGetResponse_items_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseCspRdsInstancesGetResponse_items_state(v string) (any, error) {
    result := OK_CSPRDSINSTANCESGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_CSPRDSINSTANCESGETRESPONSE_ITEMS_STATE
        case "DELETED":
            result = DELETED_CSPRDSINSTANCESGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspRdsInstancesGetResponse_items_state(values []CspRdsInstancesGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspRdsInstancesGetResponse_items_state) isMultiValue() bool {
    return false
}
