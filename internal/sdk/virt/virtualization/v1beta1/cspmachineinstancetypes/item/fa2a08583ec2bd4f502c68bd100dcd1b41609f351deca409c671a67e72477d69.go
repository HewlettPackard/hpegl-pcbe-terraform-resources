package item
type CspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies int

const (
    CLUSTER_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES CspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies = iota
    PARTITION_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES
    SPREAD_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES
)

func (i CspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies) String() string {
    return []string{"cluster", "partition", "spread"}[i]
}
func ParseCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies(v string) (any, error) {
    result := CLUSTER_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES
    switch v {
        case "cluster":
            result = CLUSTER_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES
        case "partition":
            result = PARTITION_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES
        case "spread":
            result = SPREAD_CSPMACHINEINSTANCETYPESGETRESPONSE_CSPINFOMEMBER1_PLACEMENTGROUPINFO_SUPPORTEDSTRATEGIES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies(values []CspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CspMachineInstanceTypesGetResponse_cspInfoMember1_placementGroupInfo_supportedStrategies) isMultiValue() bool {
    return false
}
