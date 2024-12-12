package installrelease
// An indicator of how a Software Release should be used:- `INSTALL`: Used for an initial install, such as the OVA for a virtual machine.- `UPGRADE`: Upgrades an existing release.- `UNIVERSAL`: Acts as both an install and upgrade media.
type InstallReleaseGetResponse_usage int

const (
    INSTALL_INSTALLRELEASEGETRESPONSE_USAGE InstallReleaseGetResponse_usage = iota
    UPGRADE_INSTALLRELEASEGETRESPONSE_USAGE
    UNIVERSAL_INSTALLRELEASEGETRESPONSE_USAGE
)

func (i InstallReleaseGetResponse_usage) String() string {
    return []string{"INSTALL", "UPGRADE", "UNIVERSAL"}[i]
}
func ParseInstallReleaseGetResponse_usage(v string) (any, error) {
    result := INSTALL_INSTALLRELEASEGETRESPONSE_USAGE
    switch v {
        case "INSTALL":
            result = INSTALL_INSTALLRELEASEGETRESPONSE_USAGE
        case "UPGRADE":
            result = UPGRADE_INSTALLRELEASEGETRESPONSE_USAGE
        case "UNIVERSAL":
            result = UNIVERSAL_INSTALLRELEASEGETRESPONSE_USAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeInstallReleaseGetResponse_usage(values []InstallReleaseGetResponse_usage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InstallReleaseGetResponse_usage) isMultiValue() bool {
    return false
}
