package download
// The files within a Software Release.
type DownloadPostRequestBody_file int

const (
    SOFTWARE_DOWNLOADPOSTREQUESTBODY_FILE DownloadPostRequestBody_file = iota
    SIGNATURE_DOWNLOADPOSTREQUESTBODY_FILE
)

func (i DownloadPostRequestBody_file) String() string {
    return []string{"SOFTWARE", "SIGNATURE"}[i]
}
func ParseDownloadPostRequestBody_file(v string) (any, error) {
    result := SOFTWARE_DOWNLOADPOSTREQUESTBODY_FILE
    switch v {
        case "SOFTWARE":
            result = SOFTWARE_DOWNLOADPOSTREQUESTBODY_FILE
        case "SIGNATURE":
            result = SIGNATURE_DOWNLOADPOSTREQUESTBODY_FILE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDownloadPostRequestBody_file(values []DownloadPostRequestBody_file) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DownloadPostRequestBody_file) isMultiValue() bool {
    return false
}
