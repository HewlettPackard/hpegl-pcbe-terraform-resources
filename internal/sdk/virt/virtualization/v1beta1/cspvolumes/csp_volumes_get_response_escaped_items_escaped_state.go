package cspvolumes
// A DELETED state indicates the asset was already deleted from the cloud serviceprovider and will be deleted from the inventory when no backups or directprotection policy relationships remain.
type CspVolumesGetResponse_items_state int

const (
    OK_CSPVOLUMESGETRESPONSE_ITEMS_STATE CspVolumesGetResponse_items_state = iota
    DELETED_CSPVOLUMESGETRESPONSE_ITEMS_STATE
)

func (i CspVolumesGetResponse_items_state) String() string {
    return []string{"OK", "DELETED"}[i]
}
func ParseCspVolumesGetResponse_items_state(v string) (any, error) {
    result := OK_CSPVOLUMESGETRESPONSE_ITEMS_STATE
    switch v {
        case "OK":
            result = OK_CSPVOLUMESGETRESPONSE_ITEMS_STATE
        case "DELETED":
            result = DELETED_CSPVOLUMESGETRESPONSE_ITEMS_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspVolumesGetResponse_items_state(values []CspVolumesGetResponse_items_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspVolumesGetResponse_items_state) isMultiValue() bool {
    return false
}
