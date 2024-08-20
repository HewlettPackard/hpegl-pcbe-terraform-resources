package datastores
// Replication Mode
type DatastoresGetResponse_items_replicationInfo_partnerDetails_mode int

const (
    SYNCHRONOUS_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_MODE DatastoresGetResponse_items_replicationInfo_partnerDetails_mode = iota
    PERIODIC_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_MODE
)

func (i DatastoresGetResponse_items_replicationInfo_partnerDetails_mode) String() string {
    return []string{"SYNCHRONOUS", "PERIODIC"}[i]
}
func ParseDatastoresGetResponse_items_replicationInfo_partnerDetails_mode(v string) (any, error) {
    result := SYNCHRONOUS_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_MODE
    switch v {
        case "SYNCHRONOUS":
            result = SYNCHRONOUS_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_MODE
        case "PERIODIC":
            result = PERIODIC_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_MODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_replicationInfo_partnerDetails_mode(values []DatastoresGetResponse_items_replicationInfo_partnerDetails_mode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_replicationInfo_partnerDetails_mode) isMultiValue() bool {
    return false
}
