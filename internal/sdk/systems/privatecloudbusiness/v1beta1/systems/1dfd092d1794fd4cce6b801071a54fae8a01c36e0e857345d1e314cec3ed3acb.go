package systems
// secret state
type SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state int

const (
    UNSPECIFIED_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state = iota
    INPROGRESS_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    INCONSISTENT_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    INSYNC_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
)

func (i SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseSystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state(v string) (any, error) {
    result := UNSPECIFIED_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INSYNC":
            result = INSYNC_SYSTEMSGETRESPONSE_ITEMS_COMPUTECLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state(values []SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_computeClusters_hypervisorManagerAdminPasswordSecret_state) isMultiValue() bool {
    return false
}
