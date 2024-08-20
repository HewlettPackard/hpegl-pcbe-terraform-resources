package hypervisorlibraryimages
// The type of the virtual machine image from the hypervisor manager image library
type HypervisorLibraryImagesGetResponse_items_fileType int

const (
    OVF_HYPERVISORLIBRARYIMAGESGETRESPONSE_ITEMS_FILETYPE HypervisorLibraryImagesGetResponse_items_fileType = iota
)

func (i HypervisorLibraryImagesGetResponse_items_fileType) String() string {
    return []string{"OVF"}[i]
}
func ParseHypervisorLibraryImagesGetResponse_items_fileType(v string) (any, error) {
    result := OVF_HYPERVISORLIBRARYIMAGESGETRESPONSE_ITEMS_FILETYPE
    switch v {
        case "OVF":
            result = OVF_HYPERVISORLIBRARYIMAGESGETRESPONSE_ITEMS_FILETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHypervisorLibraryImagesGetResponse_items_fileType(values []HypervisorLibraryImagesGetResponse_items_fileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HypervisorLibraryImagesGetResponse_items_fileType) isMultiValue() bool {
    return false
}
