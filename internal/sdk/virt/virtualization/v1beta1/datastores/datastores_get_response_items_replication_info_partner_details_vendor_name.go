package datastores
// Vendor name
type DatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName int

const (
    NIMBLE_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_VENDORNAME DatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName = iota
    PRIMERA_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_VENDORNAME
)

func (i DatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName) String() string {
    return []string{"NIMBLE", "PRIMERA"}[i]
}
func ParseDatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName(v string) (any, error) {
    result := NIMBLE_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_VENDORNAME
    switch v {
        case "NIMBLE":
            result = NIMBLE_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_VENDORNAME
        case "PRIMERA":
            result = PRIMERA_DATASTORESGETRESPONSE_ITEMS_REPLICATIONINFO_PARTNERDETAILS_VENDORNAME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName(values []DatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DatastoresGetResponse_items_replicationInfo_partnerDetails_vendorName) isMultiValue() bool {
    return false
}
