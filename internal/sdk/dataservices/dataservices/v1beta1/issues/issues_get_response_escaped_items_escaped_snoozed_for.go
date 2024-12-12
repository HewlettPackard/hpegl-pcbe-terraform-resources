package issues
type IssuesGetResponse_items_snoozedFor int

const (
    NONE_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR IssuesGetResponse_items_snoozedFor = iota
    DAY_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
    WEEK_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
    MONTH_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
    INFINITE_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
)

func (i IssuesGetResponse_items_snoozedFor) String() string {
    return []string{"NONE", "DAY", "WEEK", "MONTH", "INFINITE"}[i]
}
func ParseIssuesGetResponse_items_snoozedFor(v string) (any, error) {
    result := NONE_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
    switch v {
        case "NONE":
            result = NONE_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
        case "DAY":
            result = DAY_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
        case "WEEK":
            result = WEEK_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
        case "MONTH":
            result = MONTH_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
        case "INFINITE":
            result = INFINITE_ISSUESGETRESPONSE_ITEMS_SNOOZEDFOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIssuesGetResponse_items_snoozedFor(values []IssuesGetResponse_items_snoozedFor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IssuesGetResponse_items_snoozedFor) isMultiValue() bool {
    return false
}
