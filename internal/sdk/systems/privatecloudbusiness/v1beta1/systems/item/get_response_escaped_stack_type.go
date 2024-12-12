package item
// Stack Type of the system
type GetResponse_stackType int

const (
    DHCI_GETRESPONSE_STACKTYPE GetResponse_stackType = iota
    DHCI_GENAI_GETRESPONSE_STACKTYPE
    PCE_GETRESPONSE_STACKTYPE
    SIMPLIVITY_GETRESPONSE_STACKTYPE
)

func (i GetResponse_stackType) String() string {
    return []string{"DHCI", "DHCI_GENAI", "PCE", "SIMPLIVITY"}[i]
}
func ParseGetResponse_stackType(v string) (any, error) {
    result := DHCI_GETRESPONSE_STACKTYPE
    switch v {
        case "DHCI":
            result = DHCI_GETRESPONSE_STACKTYPE
        case "DHCI_GENAI":
            result = DHCI_GENAI_GETRESPONSE_STACKTYPE
        case "PCE":
            result = PCE_GETRESPONSE_STACKTYPE
        case "SIMPLIVITY":
            result = SIMPLIVITY_GETRESPONSE_STACKTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetResponse_stackType(values []GetResponse_stackType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetResponse_stackType) isMultiValue() bool {
    return false
}
