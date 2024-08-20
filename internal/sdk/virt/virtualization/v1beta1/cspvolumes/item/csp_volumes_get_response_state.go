package item
// A DELETED state indicates the asset was already deleted from the cloud serviceprovider and will be deleted from the inventory when no backups or directprotection policy relationships remain.
type CspVolumesGetResponse_state int

const (
    OK_CSPVOLUMESGETRESPONSE_STATE CspVolumesGetResponse_state = iota
    DELETED_CSPVOLUMESGETRESPONSE_STATE
)

func (i CspVolumesGetResponse_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseCspVolumesGetResponse_state(v string) (any, error) {
    result := OK_CSPVOLUMESGETRESPONSE_STATE
    switch v {
        case "OK":
            result = OK_CSPVOLUMESGETRESPONSE_STATE
        case "DELETED":
            result = DELETED_CSPVOLUMESGETRESPONSE_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVolumesGetResponse_state(values []CspVolumesGetResponse_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVolumesGetResponse_state) isMultiValue() bool {
    return false
}
