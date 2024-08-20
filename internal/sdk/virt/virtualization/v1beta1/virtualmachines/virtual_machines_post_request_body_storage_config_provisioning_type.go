package virtualmachines
// Specifies whether datastore is THIN or THICK to provision the virtual machine
type VirtualMachinesPostRequestBody_storageConfig_provisioningType int

const (
    THIN_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_PROVISIONINGTYPE VirtualMachinesPostRequestBody_storageConfig_provisioningType = iota
    THICK_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_PROVISIONINGTYPE
)

func (i VirtualMachinesPostRequestBody_storageConfig_provisioningType) String() string {
    return []string{"THIN", "THICK"}[i]
}
func ParseVirtualMachinesPostRequestBody_storageConfig_provisioningType(v string) (any, error) {
    result := THIN_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_PROVISIONINGTYPE
    switch v {
        case "THIN":
            result = THIN_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_PROVISIONINGTYPE
        case "THICK":
            result = THICK_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_PROVISIONINGTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesPostRequestBody_storageConfig_provisioningType(values []VirtualMachinesPostRequestBody_storageConfig_provisioningType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesPostRequestBody_storageConfig_provisioningType) isMultiValue() bool {
    return false
}
