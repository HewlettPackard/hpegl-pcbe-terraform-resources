package item
type CspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes int

const (
    EBS_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_SUPPORTEDROOTDEVICETYPES CspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes = iota
    INSTANCESTORE_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_SUPPORTEDROOTDEVICETYPES
)

func (i CspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes) String() string {
    return []string{"ebs", "instance-store"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes(v string) (any, error) {
    result := EBS_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_SUPPORTEDROOTDEVICETYPES
    switch v {
        case "ebs":
            result = EBS_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_SUPPORTEDROOTDEVICETYPES
        case "instance-store":
            result = INSTANCESTORE_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_SUPPORTEDROOTDEVICETYPES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes(values []CspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_cspInfoMember1_supportedRootDeviceTypes) isMultiValue() bool {
    return false
}
