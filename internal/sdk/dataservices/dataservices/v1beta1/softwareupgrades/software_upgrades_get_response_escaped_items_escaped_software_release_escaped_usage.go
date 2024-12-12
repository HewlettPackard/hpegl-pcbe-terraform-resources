package softwareupgrades
// An indicator of how a Software Release should be used:- `INSTALL`: Used for an initial install, such as the OVA for a virtual machine.- `UPGRADE`: Upgrades an existing release.- `UNIVERSAL`: Acts as both an install and upgrade media.
type SoftwareUpgradesGetResponse_items_softwareRelease_usage int

const (
    INSTALL_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE SoftwareUpgradesGetResponse_items_softwareRelease_usage = iota
    UPGRADE_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE
    UNIVERSAL_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE
)

func (i SoftwareUpgradesGetResponse_items_softwareRelease_usage) String() string {
    return []string{"INSTALL", "UPGRADE", "UNIVERSAL"}[i]
}
func ParseSoftwareUpgradesGetResponse_items_softwareRelease_usage(v string) (any, error) {
    result := INSTALL_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE
    switch v {
        case "INSTALL":
            result = INSTALL_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE
        case "UPGRADE":
            result = UPGRADE_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE
        case "UNIVERSAL":
            result = UNIVERSAL_SOFTWAREUPGRADESGETRESPONSE_ITEMS_SOFTWARERELEASE_USAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSoftwareUpgradesGetResponse_items_softwareRelease_usage(values []SoftwareUpgradesGetResponse_items_softwareRelease_usage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SoftwareUpgradesGetResponse_items_softwareRelease_usage) isMultiValue() bool {
    return false
}
