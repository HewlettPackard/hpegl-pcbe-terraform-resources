package item
type PatchResponse_health_overallHealth int

const (
    OK_PATCHRESPONSE_HEALTH_OVERALLHEALTH PatchResponse_health_overallHealth = iota
    WARNING_PATCHRESPONSE_HEALTH_OVERALLHEALTH
    CRITICAL_PATCHRESPONSE_HEALTH_OVERALLHEALTH
    MISSING_PATCHRESPONSE_HEALTH_OVERALLHEALTH
)

func (i PatchResponse_health_overallHealth) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParsePatchResponse_health_overallHealth(v string) (any, error) {
    result := OK_PATCHRESPONSE_HEALTH_OVERALLHEALTH
    switch v {
        case "OK":
            result = OK_PATCHRESPONSE_HEALTH_OVERALLHEALTH
        case "WARNING":
            result = WARNING_PATCHRESPONSE_HEALTH_OVERALLHEALTH
        case "CRITICAL":
            result = CRITICAL_PATCHRESPONSE_HEALTH_OVERALLHEALTH
        case "MISSING":
            result = MISSING_PATCHRESPONSE_HEALTH_OVERALLHEALTH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_health_overallHealth(values []PatchResponse_health_overallHealth) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_health_overallHealth) isMultiValue() bool {
    return false
}
