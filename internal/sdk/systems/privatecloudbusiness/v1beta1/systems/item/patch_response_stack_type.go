package item
// Stack Type of the system
type PatchResponse_stackType int

const (
    DHCI_PATCHRESPONSE_STACKTYPE PatchResponse_stackType = iota
    SIMPLIVITY_PATCHRESPONSE_STACKTYPE
)

func (i PatchResponse_stackType) String() string {
    return []string{"DHCI", "SIMPLIVITY"}[i]
}
func ParsePatchResponse_stackType(v string) (any, error) {
    result := DHCI_PATCHRESPONSE_STACKTYPE
    switch v {
        case "DHCI":
            result = DHCI_PATCHRESPONSE_STACKTYPE
        case "SIMPLIVITY":
            result = SIMPLIVITY_PATCHRESPONSE_STACKTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePatchResponse_stackType(values []PatchResponse_stackType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PatchResponse_stackType) isMultiValue() bool {
    return false
}
