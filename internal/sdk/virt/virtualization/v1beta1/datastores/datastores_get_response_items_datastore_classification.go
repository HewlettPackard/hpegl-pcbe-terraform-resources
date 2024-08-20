package datastores
// Classification of datastore types for which protection will be disabled
type DatastoresGetResponse_items_datastoreClassification int

const (
    PROTECTION_STORE_GATEWAY_DATASTORESGETRESPONSE_ITEMS_DATASTORECLASSIFICATION DatastoresGetResponse_items_datastoreClassification = iota
    INSTANT_RECOVERED_DATASTORESGETRESPONSE_ITEMS_DATASTORECLASSIFICATION
)

func (i DatastoresGetResponse_items_datastoreClassification) String() string {
    return []string{"PROTECTION_STORE_GATEWAY", "INSTANT_RECOVERED"}[i]
}
func ParseDatastoresGetResponse_items_datastoreClassification(v string) (any, error) {
    result := PROTECTION_STORE_GATEWAY_DATASTORESGETRESPONSE_ITEMS_DATASTORECLASSIFICATION
    switch v {
        case "PROTECTION_STORE_GATEWAY":
            result = PROTECTION_STORE_GATEWAY_DATASTORESGETRESPONSE_ITEMS_DATASTORECLASSIFICATION
        case "INSTANT_RECOVERED":
            result = INSTANT_RECOVERED_DATASTORESGETRESPONSE_ITEMS_DATASTORECLASSIFICATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_datastoreClassification(values []DatastoresGetResponse_items_datastoreClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_datastoreClassification) isMultiValue() bool {
    return false
}
