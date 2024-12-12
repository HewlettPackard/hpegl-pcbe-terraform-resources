package item
// Status of VMware Tools running in the guest operating system. Values are inline with the vCenter provided values
type VirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status int

const (
    NOT_INSTALLED_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS VirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status = iota
    NOT_RUNNING_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
    OK_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
    OLD_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
)

func (i VirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status) String() string {
    return []string{"NOT_INSTALLED", "NOT_RUNNING", "OK", "OLD"}[i]
}
func ParseVirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status(v string) (any, error) {
    result := NOT_INSTALLED_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
    switch v {
        case "NOT_INSTALLED":
            result = NOT_INSTALLED_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
        case "NOT_RUNNING":
            result = NOT_RUNNING_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
        case "OK":
            result = OK_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
        case "OLD":
            result = OLD_VIRTUALMACHINESGETRESPONSE_APPINFO_VMWARE_TOOLSINFO_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status(values []VirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_appInfo_vmware_toolsInfo_status) isMultiValue() bool {
    return false
}
