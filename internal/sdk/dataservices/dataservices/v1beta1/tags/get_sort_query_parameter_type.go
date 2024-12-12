package tags
type GetSortQueryParameterType int

const (
    ASC_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    DESC_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"asc", "desc"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := ASC_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "asc":
            result = ASC_GETSORTQUERYPARAMETERTYPE
        case "desc":
            result = DESC_GETSORTQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}
