package folders
// The type of the hypervisor folder.
type FoldersGetResponse_items_folderType int

const (
    DATASTORE_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE FoldersGetResponse_items_folderType = iota
    VM_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE
    HOST_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE
)

func (i FoldersGetResponse_items_folderType) String() string {
    return []string{"DATASTORE", "VM", "HOST"}[i]
}
func ParseFoldersGetResponse_items_folderType(v string) (any, error) {
    result := DATASTORE_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE
    switch v {
        case "DATASTORE":
            result = DATASTORE_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE
        case "VM":
            result = VM_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE
        case "HOST":
            result = HOST_FOLDERSGETRESPONSE_ITEMS_FOLDERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeFoldersGetResponse_items_folderType(values []FoldersGetResponse_items_folderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FoldersGetResponse_items_folderType) isMultiValue() bool {
    return false
}
