package item
// Data Services Connector deployment type.
type DataServicesConnectorsGetResponse_deploymentType int

const (
    CUSTOMER_DATASERVICESCONNECTORSGETRESPONSE_DEPLOYMENTTYPE DataServicesConnectorsGetResponse_deploymentType = iota
    FACTORY_DATASERVICESCONNECTORSGETRESPONSE_DEPLOYMENTTYPE
)

func (i DataServicesConnectorsGetResponse_deploymentType) String() string {
    return []string{"CUSTOMER", "FACTORY"}[i]
}
func ParseDataServicesConnectorsGetResponse_deploymentType(v string) (any, error) {
    result := CUSTOMER_DATASERVICESCONNECTORSGETRESPONSE_DEPLOYMENTTYPE
    switch v {
        case "CUSTOMER":
            result = CUSTOMER_DATASERVICESCONNECTORSGETRESPONSE_DEPLOYMENTTYPE
        case "FACTORY":
            result = FACTORY_DATASERVICESCONNECTORSGETRESPONSE_DEPLOYMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDataServicesConnectorsGetResponse_deploymentType(values []DataServicesConnectorsGetResponse_deploymentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataServicesConnectorsGetResponse_deploymentType) isMultiValue() bool {
    return false
}
