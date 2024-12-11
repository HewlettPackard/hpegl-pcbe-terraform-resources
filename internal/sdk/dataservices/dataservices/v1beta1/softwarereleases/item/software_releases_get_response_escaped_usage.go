package item
// An indicator of how a Software Release should be used:- `INSTALL`: Used for an initial install, such as the OVA for a virtual machine.- `UPGRADE`: Upgrades an existing release.- `UNIVERSAL`: Acts as both an install and upgrade media.
type SoftwareReleasesGetResponse_usage int

const (
    INSTALL_SOFTWARERELEASESGETRESPONSE_USAGE SoftwareReleasesGetResponse_usage = iota
    UPGRADE_SOFTWARERELEASESGETRESPONSE_USAGE
    UNIVERSAL_SOFTWARERELEASESGETRESPONSE_USAGE
)

func (i SoftwareReleasesGetResponse_usage) String() string {
    return []string{"INSTALL", "UPGRADE", "UNIVERSAL"}[i]
}
func ParseSoftwareReleasesGetResponse_usage(v string) (any, error) {
    result := INSTALL_SOFTWARERELEASESGETRESPONSE_USAGE
    switch v {
        case "INSTALL":
            result = INSTALL_SOFTWARERELEASESGETRESPONSE_USAGE
        case "UPGRADE":
            result = UPGRADE_SOFTWARERELEASESGETRESPONSE_USAGE
        case "UNIVERSAL":
            result = UNIVERSAL_SOFTWARERELEASESGETRESPONSE_USAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSoftwareReleasesGetResponse_usage(values []SoftwareReleasesGetResponse_usage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SoftwareReleasesGetResponse_usage) isMultiValue() bool {
    return false
}