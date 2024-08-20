package item
// secret state
type WithServerGetResponse_hypervisorHostRootUserPasswordSecret_state int

const (
    UNSPECIFIED_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE WithServerGetResponse_hypervisorHostRootUserPasswordSecret_state = iota
    INPROGRESS_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
    INCONSISTENT_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
    INSYNC_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
)

func (i WithServerGetResponse_hypervisorHostRootUserPasswordSecret_state) String() string {
    return []string{"UNSPECIFIED", "INPROGRESS", "INCONSISTENT", "INSYNC"}[i]
}
func ParseWithServerGetResponse_hypervisorHostRootUserPasswordSecret_state(v string) (any, error) {
    result := UNSPECIFIED_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
    switch v {
        case "UNSPECIFIED":
            result = UNSPECIFIED_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
        case "INPROGRESS":
            result = INPROGRESS_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
        case "INCONSISTENT":
            result = INCONSISTENT_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
        case "INSYNC":
            result = INSYNC_WITHSERVERGETRESPONSE_HYPERVISORHOSTROOTUSERPASSWORDSECRET_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithServerGetResponse_hypervisorHostRootUserPasswordSecret_state(values []WithServerGetResponse_hypervisorHostRootUserPasswordSecret_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithServerGetResponse_hypervisorHostRootUserPasswordSecret_state) isMultiValue() bool {
    return false
}
