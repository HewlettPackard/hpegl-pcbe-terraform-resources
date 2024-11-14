package item
// Type of OS
type WithServerGetResponse_operatingSystemInfo_operatingSystemType int

const (
    RHEL_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE WithServerGetResponse_operatingSystemInfo_operatingSystemType = iota
    LINUX_UBUNTU_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
    LINUX_COREOS_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
    HYPERVISOR_VMWARE_ESXI_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
)

func (i WithServerGetResponse_operatingSystemInfo_operatingSystemType) String() string {
    return []string{"RHEL", "LINUX_UBUNTU", "LINUX_COREOS", "HYPERVISOR_VMWARE_ESXI"}[i]
}
func ParseWithServerGetResponse_operatingSystemInfo_operatingSystemType(v string) (any, error) {
    result := RHEL_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
    switch v {
        case "RHEL":
            result = RHEL_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        case "LINUX_UBUNTU":
            result = LINUX_UBUNTU_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        case "LINUX_COREOS":
            result = LINUX_COREOS_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        case "HYPERVISOR_VMWARE_ESXI":
            result = HYPERVISOR_VMWARE_ESXI_WITHSERVERGETRESPONSE_OPERATINGSYSTEMINFO_OPERATINGSYSTEMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_operatingSystemInfo_operatingSystemType(values []WithServerGetResponse_operatingSystemInfo_operatingSystemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_operatingSystemInfo_operatingSystemType) isMultiValue() bool {
    return false
}
