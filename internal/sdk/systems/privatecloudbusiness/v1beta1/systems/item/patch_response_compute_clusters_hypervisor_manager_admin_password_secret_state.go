package item
// secret state
type PatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state int

const (
    UNSPECIFIED_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE PatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state = iota
    INPROGRESS_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    INCONSISTENT_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    INSYNC_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
)

func (i PatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParsePatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state(v string) (any, error) {
    result := UNSPECIFIED_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INSYNC":
            result = INSYNC_PATCHRESPONSE_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state(values []PatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_computeClusters_hypervisorManagerAdminPasswordSecret_state) isMultiValue() bool {
    return false
}
