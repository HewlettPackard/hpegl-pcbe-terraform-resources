package datastores
// Encryption cipher for the volumes.
type DatastoresPostRequestBody_volumeInfo_encryption_cipher int

const (
    AES_256_XTS_DATASTORESPOSTREQUESTBODY_VOLUMEINFO_ENCRYPTION_CIPHER DatastoresPostRequestBody_volumeInfo_encryption_cipher = iota
    NONE_DATASTORESPOSTREQUESTBODY_VOLUMEINFO_ENCRYPTION_CIPHER
)

func (i DatastoresPostRequestBody_volumeInfo_encryption_cipher) String() string {
    return []string{"AES_256_XTS", "None"}[i]
}
func ParseDatastoresPostRequestBody_volumeInfo_encryption_cipher(v string) (any, error) {
    result := AES_256_XTS_DATASTORESPOSTREQUESTBODY_VOLUMEINFO_ENCRYPTION_CIPHER
    switch v {
        case "AES_256_XTS":
            result = AES_256_XTS_DATASTORESPOSTREQUESTBODY_VOLUMEINFO_ENCRYPTION_CIPHER
        case "None":
            result = NONE_DATASTORESPOSTREQUESTBODY_VOLUMEINFO_ENCRYPTION_CIPHER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresPostRequestBody_volumeInfo_encryption_cipher(values []DatastoresPostRequestBody_volumeInfo_encryption_cipher) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresPostRequestBody_volumeInfo_encryption_cipher) isMultiValue() bool {
    return false
}
