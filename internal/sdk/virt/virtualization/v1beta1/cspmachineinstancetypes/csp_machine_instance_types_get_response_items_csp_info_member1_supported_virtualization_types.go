package cspmachineinstancetypes
type CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes int

const (
    HVM_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDVIRTUALIZATIONTYPES CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes = iota
)

func (i CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes) String() string {
    return []string{"hvm"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes(v string) (any, error) {
    result := HVM_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDVIRTUALIZATIONTYPES
    switch v {
        case "hvm":
            result = HVM_CSPMACHINEINSTANCETYPESGETRESPONSE_ITEMS_CSPINFOMEMBER1_SUPPORTEDVIRTUALIZATIONTYPES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes(values []CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_items_cspInfoMember1_supportedVirtualizationTypes) isMultiValue() bool {
    return false
}
