package item
// secret state
type GetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state int

const (
    UNSPECIFIED_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE GetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state = iota
    INPROGRESS_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    INCONSISTENT_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    INSYNC_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
)

func (i GetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state(v string) (any, error) {
    result := UNSPECIFIED_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        case "INSYNC":
            result = INSYNC_GETRESPONSE_HYPERVISORCLUSTERS_HYPERVISORMANAGERADMINPASSWORDSECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state(values []GetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_hypervisorClusters_hypervisorManagerAdminPasswordSecret_state) isMultiValue() bool {
    return false
}
