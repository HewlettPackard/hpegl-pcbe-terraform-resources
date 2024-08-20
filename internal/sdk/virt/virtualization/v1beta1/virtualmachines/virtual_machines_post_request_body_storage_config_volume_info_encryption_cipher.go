package virtualmachines
// Encryption cipher for the volumes.
type VirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher int

const (
    AES_256_XTS_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_VOLUMEINFO_ENCRYPTION_CIPHER VirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher = iota
    NONE_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_VOLUMEINFO_ENCRYPTION_CIPHER
)

func (i VirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher) String() string {
    return []string{"AES_256_XTS", "None"}[i]
}
func ParseVirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher(v string) (any, error) {
    result := AES_256_XTS_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_VOLUMEINFO_ENCRYPTION_CIPHER
    switch v {
        case "AES_256_XTS":
            result = AES_256_XTS_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_VOLUMEINFO_ENCRYPTION_CIPHER
        case "None":
            result = NONE_VIRTUALMACHINESPOSTREQUESTBODY_STORAGECONFIG_VOLUMEINFO_ENCRYPTION_CIPHER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher(values []VirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualMachinesPostRequestBody_storageConfig_volumeInfo_encryption_cipher) isMultiValue() bool {
    return false
}
