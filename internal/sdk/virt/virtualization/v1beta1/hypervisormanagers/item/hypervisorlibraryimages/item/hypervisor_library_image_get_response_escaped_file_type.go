package item
// The type of the virtual machine image from the hypervisor manager image library
type HypervisorLibraryImageGetResponse_fileType int

const (
    OVF_HYPERVISORLIBRARYIMAGEGETRESPONSE_FILETYPE HypervisorLibraryImageGetResponse_fileType = iota
)

func (i HypervisorLibraryImageGetResponse_fileType) String() string {
    return []string{"OVF"}[i]
}
func ParseHypervisorLibraryImageGetResponse_fileType(v string) (any, error) {
    result := OVF_HYPERVISORLIBRARYIMAGEGETRESPONSE_FILETYPE
    switch v {
        case "OVF":
            result = OVF_HYPERVISORLIBRARYIMAGEGETRESPONSE_FILETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorLibraryImageGetResponse_fileType(values []HypervisorLibraryImageGetResponse_fileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorLibraryImageGetResponse_fileType) isMultiValue() bool {
    return false
}
