package item
type PatchResponse_health_servers int

const (
    OK_PATCHRESPONSE_HEALTH_SERVERS PatchResponse_health_servers = iota
    WARNING_PATCHRESPONSE_HEALTH_SERVERS
    CRITICAL_PATCHRESPONSE_HEALTH_SERVERS
    MISSING_PATCHRESPONSE_HEALTH_SERVERS
)

func (i PatchResponse_health_servers) String() string {
    return []string{"OK", "WARNING", "CRITICAL", "MISSING"}[i]
}
func ParsePatchResponse_health_servers(v string) (any, error) {
    result := OK_PATCHRESPONSE_HEALTH_SERVERS
    switch v {
        case "OK":
            result = OK_PATCHRESPONSE_HEALTH_SERVERS
        case "WARNING":
            result = WARNING_PATCHRESPONSE_HEALTH_SERVERS
        case "CRITICAL":
            result = CRITICAL_PATCHRESPONSE_HEALTH_SERVERS
        case "MISSING":
            result = MISSING_PATCHRESPONSE_HEALTH_SERVERS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_health_servers(values []PatchResponse_health_servers) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_health_servers) isMultiValue() bool {
    return false
}
