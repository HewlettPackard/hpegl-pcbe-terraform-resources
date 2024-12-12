package item
type PatchResponse_health_storage int

const (
    OK_PATCHRESPONSE_HEALTH_STORAGE PatchResponse_health_storage = iota
    WARNING_PATCHRESPONSE_HEALTH_STORAGE
    CRITICAL_PATCHRESPONSE_HEALTH_STORAGE
    MISSING_PATCHRESPONSE_HEALTH_STORAGE
)

func (i PatchResponse_health_storage) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParsePatchResponse_health_storage(v string) (any, error) {
    result := OK_PATCHRESPONSE_HEALTH_STORAGE
    switch v {
        case "OK":
            result = OK_PATCHRESPONSE_HEALTH_STORAGE
        case "WARNING":
            result = WARNING_PATCHRESPONSE_HEALTH_STORAGE
        case "CRITICAL":
            result = CRITICAL_PATCHRESPONSE_HEALTH_STORAGE
        case "MISSING":
            result = MISSING_PATCHRESPONSE_HEALTH_STORAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_health_storage(values []PatchResponse_health_storage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_health_storage) isMultiValue() bool {
    return false
}
