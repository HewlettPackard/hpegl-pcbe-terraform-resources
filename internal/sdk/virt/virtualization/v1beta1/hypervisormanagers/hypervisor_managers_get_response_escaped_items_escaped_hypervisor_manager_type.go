package hypervisormanagers
// The type of the hypervisor manager. Currently only vCenter is supported.
type HypervisorManagersGetResponse_items_hypervisorManagerType int

const (
    VMWARE_VCENTER_HYPERVISORMANAGERSGETRESPONSE_ITEMS_HYPERVISORMANAGERTYPE HypervisorManagersGetResponse_items_hypervisorManagerType = iota
)

func (i HypervisorManagersGetResponse_items_hypervisorManagerType) String() string {
    return []string{"VMWARE_VCENTER"}[i]
}
func ParseHypervisorManagersGetResponse_items_hypervisorManagerType(v string) (any, error) {
    result := VMWARE_VCENTER_HYPERVISORMANAGERSGETRESPONSE_ITEMS_HYPERVISORMANAGERTYPE
    switch v {
        case "VMWARE_VCENTER":
            result = VMWARE_VCENTER_HYPERVISORMANAGERSGETRESPONSE_ITEMS_HYPERVISORMANAGERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorManagersGetResponse_items_hypervisorManagerType(values []HypervisorManagersGetResponse_items_hypervisorManagerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorManagersGetResponse_items_hypervisorManagerType) isMultiValue() bool {
    return false
}
