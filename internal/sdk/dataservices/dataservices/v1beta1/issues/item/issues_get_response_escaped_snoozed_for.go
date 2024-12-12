package item
type IssuesGetResponse_snoozedFor int

const (
    NONE_ISSUESGETRESPONSE_SNOOZEDFOR IssuesGetResponse_snoozedFor = iota
    DAY_ISSUESGETRESPONSE_SNOOZEDFOR
    WEEK_ISSUESGETRESPONSE_SNOOZEDFOR
    MONTH_ISSUESGETRESPONSE_SNOOZEDFOR
    INFINITE_ISSUESGETRESPONSE_SNOOZEDFOR
)

func (i IssuesGetResponse_snoozedFor) String() string {
    return []string{"NONE", "DAY", "WEEK", "MONTH", "INFINITE"}[i]
}
func ParseIssuesGetResponse_snoozedFor(v string) (any, error) {
    result := NONE_ISSUESGETRESPONSE_SNOOZEDFOR
    switch v {
        case "NONE":
            result = NONE_ISSUESGETRESPONSE_SNOOZEDFOR
        case "DAY":
            result = DAY_ISSUESGETRESPONSE_SNOOZEDFOR
        case "WEEK":
            result = WEEK_ISSUESGETRESPONSE_SNOOZEDFOR
        case "MONTH":
            result = MONTH_ISSUESGETRESPONSE_SNOOZEDFOR
        case "INFINITE":
            result = INFINITE_ISSUESGETRESPONSE_SNOOZEDFOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIssuesGetResponse_snoozedFor(values []IssuesGetResponse_snoozedFor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IssuesGetResponse_snoozedFor) isMultiValue() bool {
    return false
}
