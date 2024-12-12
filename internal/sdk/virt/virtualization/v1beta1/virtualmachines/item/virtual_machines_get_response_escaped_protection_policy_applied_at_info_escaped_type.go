package item
// Type of the protection policy.
type VirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type int

const (
    DATASTORE_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE VirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type = iota
    VIRTUAL_MACHINE_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
    VM_PROTECTION_GROUP_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
)

func (i VirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type) String() string {
    return []string{"DATASTORE", "VIRTUAL_MACHINE", "VM_PROTECTION_GROUP"}[i]
}
func ParseVirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type(v string) (any, error) {
    result := DATASTORE_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
    switch v {
        case "DATASTORE":
            result = DATASTORE_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
        case "VIRTUAL_MACHINE":
            result = VIRTUAL_MACHINE_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
        case "VM_PROTECTION_GROUP":
            result = VM_PROTECTION_GROUP_VIRTUALMACHINESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type(values []VirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesGetResponse_protectionPolicyAppliedAtInfo_type) isMultiValue() bool {
    return false
}
