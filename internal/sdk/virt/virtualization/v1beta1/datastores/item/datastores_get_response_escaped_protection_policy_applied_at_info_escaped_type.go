package item
// Type of the protection policy.
type DatastoresGetResponse_protectionPolicyAppliedAtInfo_type int

const (
    DATASTORE_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE DatastoresGetResponse_protectionPolicyAppliedAtInfo_type = iota
    VIRTUAL_MACHINE_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
    VM_PROTECTION_GROUP_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
)

func (i DatastoresGetResponse_protectionPolicyAppliedAtInfo_type) String() string {
    return []string{"DATASTORE", "VIRTUAL_MACHINE", "VM_PROTECTION_GROUP"}[i]
}
func ParseDatastoresGetResponse_protectionPolicyAppliedAtInfo_type(v string) (any, error) {
    result := DATASTORE_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
    switch v {
        case "DATASTORE":
            result = DATASTORE_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
        case "VIRTUAL_MACHINE":
            result = VIRTUAL_MACHINE_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
        case "VM_PROTECTION_GROUP":
            result = VM_PROTECTION_GROUP_DATASTORESGETRESPONSE_PROTECTIONPOLICYAPPLIEDATINFO_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_protectionPolicyAppliedAtInfo_type(values []DatastoresGetResponse_protectionPolicyAppliedAtInfo_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_protectionPolicyAppliedAtInfo_type) isMultiValue() bool {
    return false
}
