package virtualmachines
// Specifies whether IP addresses are allocated by DHCP or static address
type VirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy int

const (
    DHCP_POLICY_VIRTUALMACHINESPOSTREQUESTBODY_NETWORKCONFIG_IPALLOCATIONPOLICY VirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy = iota
    FIXED_POLICY_VIRTUALMACHINESPOSTREQUESTBODY_NETWORKCONFIG_IPALLOCATIONPOLICY
)

func (i VirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy) String() string {
    return []string{"DHCP_POLICY", "FIXED_POLICY"}[i]
}
func ParseVirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy(v string) (any, error) {
    result := DHCP_POLICY_VIRTUALMACHINESPOSTREQUESTBODY_NETWORKCONFIG_IPALLOCATIONPOLICY
    switch v {
        case "DHCP_POLICY":
            result = DHCP_POLICY_VIRTUALMACHINESPOSTREQUESTBODY_NETWORKCONFIG_IPALLOCATIONPOLICY
        case "FIXED_POLICY":
            result = FIXED_POLICY_VIRTUALMACHINESPOSTREQUESTBODY_NETWORKCONFIG_IPALLOCATIONPOLICY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy(values []VirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesPostRequestBody_networkConfig_ipAllocationPolicy) isMultiValue() bool {
    return false
}
