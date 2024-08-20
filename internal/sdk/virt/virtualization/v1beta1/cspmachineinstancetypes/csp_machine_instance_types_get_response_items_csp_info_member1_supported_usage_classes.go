package cspmachineinstancetypes
type CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses int

const (
    ONDEMAND_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDUSAGECLASSES CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses = iota
    SPOT_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDUSAGECLASSES
)

func (i CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses) String() string {
    return []string{"on-demand", "spot"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses(v string) (any, error) {
    result := ONDEMAND_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDUSAGECLASSES
    switch v {
        case "on-demand":
            result = ONDEMAND_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDUSAGECLASSES
        case "spot":
            result = SPOT_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDUSAGECLASSES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses(values []CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedUsageClasses) isMultiValue() bool {
    return false
}
