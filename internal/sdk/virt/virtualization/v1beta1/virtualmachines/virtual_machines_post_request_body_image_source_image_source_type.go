package virtualmachines
// The source of the image where it was hosted
type VirtualMachinesPostRequestBody_imageSource_imageSourceType int

const (
    HYPERVISOR_IMAGE_LIBRARY_VIRTUALMACHINESPOSTREQUESTBODY_IMAGESOURCE_IMAGESOURCETYPE VirtualMachinesPostRequestBody_imageSource_imageSourceType = iota
)

func (i VirtualMachinesPostRequestBody_imageSource_imageSourceType) String() string {
    return []string{"HYPERVISOR_IMAGE_LIBRARY"}[i]
}
func ParseVirtualMachinesPostRequestBody_imageSource_imageSourceType(v string) (any, error) {
    result := HYPERVISOR_IMAGE_LIBRARY_VIRTUALMACHINESPOSTREQUESTBODY_IMAGESOURCE_IMAGESOURCETYPE
    switch v {
        case "HYPERVISOR_IMAGE_LIBRARY":
            result = HYPERVISOR_IMAGE_LIBRARY_VIRTUALMACHINESPOSTREQUESTBODY_IMAGESOURCE_IMAGESOURCETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesPostRequestBody_imageSource_imageSourceType(values []VirtualMachinesPostRequestBody_imageSource_imageSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesPostRequestBody_imageSource_imageSourceType) isMultiValue() bool {
    return false
}
