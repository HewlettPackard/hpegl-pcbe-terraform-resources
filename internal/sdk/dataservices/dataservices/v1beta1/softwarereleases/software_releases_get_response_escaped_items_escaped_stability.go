package softwarereleases
// The stabilities that can be assigned to a Software Release.
type SoftwareReleasesGetResponse_items_stability int

const (
    GENERAL_AVAILABILITY_SOFTWARERELEASESGETRESPONSE_ITEMS_STABILITY SoftwareReleasesGetResponse_items_stability = iota
)

func (i SoftwareReleasesGetResponse_items_stability) String() string {
    return []string{"GENERAL_AVAILABILITY"}[i]
}
func ParseSoftwareReleasesGetResponse_items_stability(v string) (any, error) {
    result := GENERAL_AVAILABILITY_SOFTWARERELEASESGETRESPONSE_ITEMS_STABILITY
    switch v {
        case "GENERAL_AVAILABILITY":
            result = GENERAL_AVAILABILITY_SOFTWARERELEASESGETRESPONSE_ITEMS_STABILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSoftwareReleasesGetResponse_items_stability(values []SoftwareReleasesGetResponse_items_stability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SoftwareReleasesGetResponse_items_stability) isMultiValue() bool {
    return false
}
