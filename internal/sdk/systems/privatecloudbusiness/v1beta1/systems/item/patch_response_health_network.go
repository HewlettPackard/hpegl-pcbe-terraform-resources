package item
type PatchResponse_health_network int

const (
    OK_PATCHRESPONSE_HEALTH_NETWORK PatchResponse_health_network = iota
    WARNING_PATCHRESPONSE_HEALTH_NETWORK
    CRITICAL_PATCHRESPONSE_HEALTH_NETWORK
    MISSING_PATCHRESPONSE_HEALTH_NETWORK
)

func (i PatchResponse_health_network) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParsePatchResponse_health_network(v string) (any, error) {
    result := OK_PATCHRESPONSE_HEALTH_NETWORK
    switch v {
        case "OK":
            result = OK_PATCHRESPONSE_HEALTH_NETWORK
        case "WARNING":
            result = WARNING_PATCHRESPONSE_HEALTH_NETWORK
        case "CRITICAL":
            result = CRITICAL_PATCHRESPONSE_HEALTH_NETWORK
        case "MISSING":
            result = MISSING_PATCHRESPONSE_HEALTH_NETWORK
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_health_network(values []PatchResponse_health_network) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_health_network) isMultiValue() bool {
    return false
}
