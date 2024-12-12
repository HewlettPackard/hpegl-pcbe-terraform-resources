package dataservicesconnectors
// Data Services Connector deployment type.
type DataServicesConnectorsGetResponse_items_deploymentType int

const (
    CUSTOMER_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DEPLOYMENTTYPE DataServicesConnectorsGetResponse_items_deploymentType = iota
    FACTORY_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DEPLOYMENTTYPE
)

func (i DataServicesConnectorsGetResponse_items_deploymentType) String() string {
    return []string{"CUSTOMER", "FACTORY"}[i]
}
func ParseDataServicesConnectorsGetResponse_items_deploymentType(v string) (any, error) {
    result := CUSTOMER_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DEPLOYMENTTYPE
    switch v {
        case "CUSTOMER":
            result = CUSTOMER_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DEPLOYMENTTYPE
        case "FACTORY":
            result = FACTORY_DATASERVICESCONNECTORSGETRESPONSE_ITEMS_DEPLOYMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_items_deploymentType(values []DataServicesConnectorsGetResponse_items_deploymentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_items_deploymentType) isMultiValue() bool {
    return false
}
