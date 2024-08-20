package item
// Classification of datastore types for which protection will be disabled
type DatastoresGetResponse_datastoreClassification int

const (
    PROTECTION_STORE_GATEWAY_DATASTORESGETRESPONSE_DATASTORECLASSIFICATION DatastoresGetResponse_datastoreClassification = iota
    INSTANT_RECOVERED_DATASTORESGETRESPONSE_DATASTORECLASSIFICATION
)

func (i DatastoresGetResponse_datastoreClassification) String() string {
    return []string{"PROTECTION_STORE_GATEWAY", "INSTANT_RECOVERED"}[i]
}
func ParseDatastoresGetResponse_datastoreClassification(v string) (any, error) {
    result := PROTECTION_STORE_GATEWAY_DATASTORESGETRESPONSE_DATASTORECLASSIFICATION
    switch v {
        case "PROTECTION_STORE_GATEWAY":
            result = PROTECTION_STORE_GATEWAY_DATASTORESGETRESPONSE_DATASTORECLASSIFICATION
        case "INSTANT_RECOVERED":
            result = INSTANT_RECOVERED_DATASTORESGETRESPONSE_DATASTORECLASSIFICATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_datastoreClassification(values []DatastoresGetResponse_datastoreClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_datastoreClassification) isMultiValue() bool {
    return false
}
