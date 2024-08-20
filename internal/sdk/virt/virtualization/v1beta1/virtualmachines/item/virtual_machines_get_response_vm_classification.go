package item
// Classification of different system/control VM types
type VirtualMachinesGetResponse_vmClassification int

const (
    DATA_ORCHESTRATOR_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION VirtualMachinesGetResponse_vmClassification = iota
    PROTECTION_STORE_GATEWAY_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    VCLS_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    TEMPLATE_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    OMNICUBE_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    DSC_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    INSTANT_RECOVERED_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    PRIVATE_CLOUD_AI_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
)

func (i VirtualMachinesGetResponse_vmClassification) String() string {
    return []string{"DATA_ORCHESTRATOR", "PROTECTION_STORE_GATEWAY", "VCLS_VM", "TEMPLATE_VM", "OMNICUBE_VM", "DSC_VM", "INSTANT_RECOVERED_VM", "PRIVATE_CLOUD_AI_VM"}[i]
}
func ParseVirtualMachinesGetResponse_vmClassification(v string) (any, error) {
    result := DATA_ORCHESTRATOR_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
    switch v {
        case "DATA_ORCHESTRATOR":
            result = DATA_ORCHESTRATOR_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "PROTECTION_STORE_GATEWAY":
            result = PROTECTION_STORE_GATEWAY_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "VCLS_VM":
            result = VCLS_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "TEMPLATE_VM":
            result = TEMPLATE_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "OMNICUBE_VM":
            result = OMNICUBE_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "DSC_VM":
            result = DSC_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "INSTANT_RECOVERED_VM":
            result = INSTANT_RECOVERED_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        case "PRIVATE_CLOUD_AI_VM":
            result = PRIVATE_CLOUD_AI_VM_VIRTUALMACHINESGETRESPONSE_VMCLASSIFICATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_vmClassification(values []VirtualMachinesGetResponse_vmClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_vmClassification) isMultiValue() bool {
    return false
}
