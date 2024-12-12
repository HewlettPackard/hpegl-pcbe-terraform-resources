package systems
// Stack Type of the system
type SystemsGetResponse_items_stackType int

const (
    DHCI_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE SystemsGetResponse_items_stackType = iota
    DHCI_GENAI_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
    PCE_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
    SIMPLIVITY_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
)

func (i SystemsGetResponse_items_stackType) String() string {
    return []string{"DHCI", "DHCI_GENAI", "PCE", "SIMPLIVITY"}[i]
}
func ParseSystemsGetResponse_items_stackType(v string) (any, error) {
    result := DHCI_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
    switch v {
        case "DHCI":
            result = DHCI_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
        case "DHCI_GENAI":
            result = DHCI_GENAI_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
        case "PCE":
            result = PCE_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
        case "SIMPLIVITY":
            result = SIMPLIVITY_SYSTEMSGETRESPONSE_ITEMS_STACKTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemsGetResponse_items_stackType(values []SystemsGetResponse_items_stackType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemsGetResponse_items_stackType) isMultiValue() bool {
    return false
}
