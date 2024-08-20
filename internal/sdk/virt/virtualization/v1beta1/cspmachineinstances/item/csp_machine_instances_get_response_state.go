package item
// A DELETED state indicates the asset was already deleted from the cloud serviceprovider and will be deleted from the inventory when no backups or directprotection policy relationships remain.
type CspMachineInstancesGetResponse_state int

const (
    OK_CSPMACHINEINSTANCESGETRESPONSE_STATE CspMachineInstancesGetResponse_state = iota
    DELETED_CSPMACHINEINSTANCESGETRESPONSE_STATE
)

func (i CspMachineInstancesGetResponse_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseCspMachineInstancesGetResponse_state(v string) (any, error) {
    result := OK_CSPMACHINEINSTANCESGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_CSPMACHINEINSTANCESGETRESPONSE_STATE
        case "DELETED":
            result = DELETED_CSPMACHINEINSTANCESGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstancesGetResponse_state(values []CspMachineInstancesGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstancesGetResponse_state) isMultiValue() bool {
    return false
}
