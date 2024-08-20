package customizeguestos
// IP address protocol. This is a mandatory parameter for linuxConfig customization.
type CustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type int

const (
    STATIC_CUSTOMIZEGUESTOSPOSTREQUESTBODY_LINUXCONFIG_NETWORKINTERFACES_NETWORKADAPTER_IPV4_TYPE CustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type = iota
    DHCP_CUSTOMIZEGUESTOSPOSTREQUESTBODY_LINUXCONFIG_NETWORKINTERFACES_NETWORKADAPTER_IPV4_TYPE
)

func (i CustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type) String() string {
    return []string{"STATIC", "DHCP"}[i]
}
func ParseCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type(v string) (any, error) {
    result := STATIC_CUSTOMIZEGUESTOSPOSTREQUESTBODY_LINUXCONFIG_NETWORKINTERFACES_NETWORKADAPTER_IPV4_TYPE
    switch v {
        case "STATIC":
            result = STATIC_CUSTOMIZEGUESTOSPOSTREQUESTBODY_LINUXCONFIG_NETWORKINTERFACES_NETWORKADAPTER_IPV4_TYPE
        case "DHCP":
            result = DHCP_CUSTOMIZEGUESTOSPOSTREQUESTBODY_LINUXCONFIG_NETWORKINTERFACES_NETWORKADAPTER_IPV4_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type(values []CustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CustomizeGuestOsPostRequestBody_linuxConfig_networkInterfaces_networkAdapter_ipv4_type) isMultiValue() bool {
    return false
}
