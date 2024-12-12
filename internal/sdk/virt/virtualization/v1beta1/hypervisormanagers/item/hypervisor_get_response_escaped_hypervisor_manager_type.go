package item
// The type of the hypervisor manager. Currently only vCenter is supported.
type HypervisorGetResponse_hypervisorManagerType int

const (
    VMWARE_VCENTER_HYPERVISORGETRESPONSE_HYPERVISORMANAGERTYPE HypervisorGetResponse_hypervisorManagerType = iota
)

func (i HypervisorGetResponse_hypervisorManagerType) String() string {
    return []string{"VMWARE_VCENTER"}[i]
}
func ParseHypervisorGetResponse_hypervisorManagerType(v string) (any, error) {
    result := VMWARE_VCENTER_HYPERVISORGETRESPONSE_HYPERVISORMANAGERTYPE
    switch v {
        case "VMWARE_VCENTER":
            result = VMWARE_VCENTER_HYPERVISORGETRESPONSE_HYPERVISORMANAGERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorGetResponse_hypervisorManagerType(values []HypervisorGetResponse_hypervisorManagerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorGetResponse_hypervisorManagerType) isMultiValue() bool {
    return false
}
