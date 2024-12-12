package item
// An enum indicating the type of authentication to be used when publishing events to the destination endpoint.This allows the endpoint to authenticate that HPE DSCC is the caller.- `BEARER_TOKEN`: A customer provided token that is added to the `Authorization` HTTP header with the bearer authentication scheme.
type WebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type int

const (
    BEARER_TOKEN_WEBHOOKSPATCHRESPONSE_DESTINATIONS_CLOUDEVENTENDPOINT_AUTHENTICATION_TYPE WebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type = iota
)

func (i WebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type) String() string {
    return []string{"BEARER_TOKEN"}[i]
}
func ParseWebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type(v string) (any, error) {
    result := BEARER_TOKEN_WEBHOOKSPATCHRESPONSE_DESTINATIONS_CLOUDEVENTENDPOINT_AUTHENTICATION_TYPE
    switch v {
        case "BEARER_TOKEN":
            result = BEARER_TOKEN_WEBHOOKSPATCHRESPONSE_DESTINATIONS_CLOUDEVENTENDPOINT_AUTHENTICATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type(values []WebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPatchResponse_destinations_cloudEventEndpoint_authentication_type) isMultiValue() bool {
    return false
}
