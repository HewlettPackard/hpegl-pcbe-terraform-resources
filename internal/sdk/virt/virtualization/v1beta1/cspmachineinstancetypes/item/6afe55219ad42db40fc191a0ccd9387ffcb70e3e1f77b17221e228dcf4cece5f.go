package item
type CspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures int

const (
    X86_64_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES CspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures = iota
    ARM64_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES
    I386_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES
)

func (i CspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures) String() string {
    return []string{"x86_64", "arm64", "i386"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures(v string) (any, error) {
    result := X86_64_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES
    switch v {
        case "x86_64":
            result = X86_64_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES
        case "arm64":
            result = ARM64_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES
        case "i386":
            result = I386_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PROCESSORINFO_SUPPORTEDARCHITECTURES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures(values []CspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_cspInfoMember1_processorInfo_supportedArchitectures) isMultiValue() bool {
    return false
}
