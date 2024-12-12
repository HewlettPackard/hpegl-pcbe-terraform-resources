package item
// The type of the hypervisor folder.
type FolderGetResponse_folderType int

const (
    DATASTORE_FOLDERGETRESPONSE_FOLDERTYPE FolderGetResponse_folderType = iota
    VM_FOLDERGETRESPONSE_FOLDERTYPE
    HOST_FOLDERGETRESPONSE_FOLDERTYPE
)

func (i FolderGetResponse_folderType) String() string {
    return []string{"DATASTORE", "VM", "HOST"}[i]
}
func ParseFolderGetResponse_folderType(v string) (any, error) {
    result := DATASTORE_FOLDERGETRESPONSE_FOLDERTYPE
    switch v {
        case "DATASTORE":
            result = DATASTORE_FOLDERGETRESPONSE_FOLDERTYPE
        case "VM":
            result = VM_FOLDERGETRESPONSE_FOLDERTYPE
        case "HOST":
            result = HOST_FOLDERGETRESPONSE_FOLDERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeFolderGetResponse_folderType(values []FolderGetResponse_folderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FolderGetResponse_folderType) isMultiValue() bool {
    return false
}
