package item
// The stabilities that can be assigned to a Software Release.
type SoftwareReleasesGetResponse_stability int

const (
    GENERAL_AVAILABILITY_SOFTWARERELEASESGETRESPONSE_STABILITY SoftwareReleasesGetResponse_stability = iota
)

func (i SoftwareReleasesGetResponse_stability) String() string {
    return []string{"GENERAL_AVAILABILITY"}[i]
}
func ParseSoftwareReleasesGetResponse_stability(v string) (any, error) {
    result := GENERAL_AVAILABILITY_SOFTWARERELEASESGETRESPONSE_STABILITY
    switch v {
        case "GENERAL_AVAILABILITY":
            result = GENERAL_AVAILABILITY_SOFTWARERELEASESGETRESPONSE_STABILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSoftwareReleasesGetResponse_stability(values []SoftwareReleasesGetResponse_stability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SoftwareReleasesGetResponse_stability) isMultiValue() bool {
    return false
}
