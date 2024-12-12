package softwareupgrades
// The stabilities that can be assigned to a Software Release.
type SoftwareUpgradesGetResponse_items_softwareRelease_stability int

const (
    GENERAL_AVAILABILITY_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_STABILITY SoftwareUpgradesGetResponse_items_softwareRelease_stability = iota
)

func (i SoftwareUpgradesGetResponse_items_softwareRelease_stability) String() string {
    return []string{"GENERAL_AVAILABILITY"}[i]
}
func ParseSoftwareUpgradesGetResponse_items_softwareRelease_stability(v string) (any, error) {
    result := GENERAL_AVAILABILITY_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_STABILITY
    switch v {
        case "GENERAL_AVAILABILITY":
            result = GENERAL_AVAILABILITY_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_STABILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSoftwareUpgradesGetResponse_items_softwareRelease_stability(values []SoftwareUpgradesGetResponse_items_softwareRelease_stability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SoftwareUpgradesGetResponse_items_softwareRelease_stability) isMultiValue() bool {
    return false
}
