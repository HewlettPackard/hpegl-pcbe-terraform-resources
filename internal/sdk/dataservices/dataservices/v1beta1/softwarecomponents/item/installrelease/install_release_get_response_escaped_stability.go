package installrelease
// The stabilities that can be assigned to a Software Release.
type InstallReleaseGetResponse_stability int

const (
    GENERAL_AVAILABILITY_INSTALLRELEASEGETRESPONSE_STABILITY InstallReleaseGetResponse_stability = iota
)

func (i InstallReleaseGetResponse_stability) String() string {
    return []string{"GENERAL_AVAILABILITY"}[i]
}
func ParseInstallReleaseGetResponse_stability(v string) (any, error) {
    result := GENERAL_AVAILABILITY_INSTALLRELEASEGETRESPONSE_STABILITY
    switch v {
        case "GENERAL_AVAILABILITY":
            result = GENERAL_AVAILABILITY_INSTALLRELEASEGETRESPONSE_STABILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeInstallReleaseGetResponse_stability(values []InstallReleaseGetResponse_stability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InstallReleaseGetResponse_stability) isMultiValue() bool {
    return false
}
