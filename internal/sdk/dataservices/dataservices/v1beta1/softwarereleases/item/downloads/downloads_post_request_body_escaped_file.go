package downloads
// The files within a Software Release.
type DownloadsPostRequestBody_file int

const (
    SOFTWARE_DOWNLOADSPOSTREQUESTBODY_FILE DownloadsPostRequestBody_file = iota
    SIGNATURE_DOWNLOADSPOSTREQUESTBODY_FILE
)

func (i DownloadsPostRequestBody_file) String() string {
    return []string{"SOFTWARE", "SIGNATURE"}[i]
}
func ParseDownloadsPostRequestBody_file(v string) (any, error) {
    result := SOFTWARE_DOWNLOADSPOSTREQUESTBODY_FILE
    switch v {
        case "SOFTWARE":
            result = SOFTWARE_DOWNLOADSPOSTREQUESTBODY_FILE
        case "SIGNATURE":
            result = SIGNATURE_DOWNLOADSPOSTREQUESTBODY_FILE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDownloadsPostRequestBody_file(values []DownloadsPostRequestBody_file) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DownloadsPostRequestBody_file) isMultiValue() bool {
    return false
}
