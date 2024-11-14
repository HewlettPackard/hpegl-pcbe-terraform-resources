package servers
// Type of OS
type ServersGetResponse_items_operatingSystemInfo_operatingSystemType int

const (
    RHEL_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE ServersGetResponse_items_operatingSystemInfo_operatingSystemType = iota
    LINUX_UBUNTU_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
    LINUX_COREOS_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
    HYPERVISOR_VMWARE_ESXI_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
)

func (i ServersGetResponse_items_operatingSystemInfo_operatingSystemType) String() string {
    return []string{"RHEL", "LINUX_UBUNTU", "LINUX_COREOS", "HYPERVISOR_VMWARE_ESXI"}[i]
}
func ParseServersGetResponse_items_operatingSystemInfo_operatingSystemType(v string) (any, error) {
    result := RHEL_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
    switch v {
        case "RHEL":
            result = RHEL_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        case "LINUX_UBUNTU":
            result = LINUX_UBUNTU_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        case "LINUX_COREOS":
            result = LINUX_COREOS_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        case "HYPERVISOR_VMWARE_ESXI":
            result = HYPERVISOR_VMWARE_ESXI_SERVERSGETRESPONSE_ITEMS_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeServersGetResponse_items_operatingSystemInfo_operatingSystemType(values []ServersGetResponse_items_operatingSystemInfo_operatingSystemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ServersGetResponse_items_operatingSystemInfo_operatingSystemType) isMultiValue() bool {
    return false
}
