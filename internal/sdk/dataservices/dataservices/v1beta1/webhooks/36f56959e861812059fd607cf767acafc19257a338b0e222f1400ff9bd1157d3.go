package webhooks
// An enum indicating the type of authentication to be used when publishing events to the destination endpoint.This allows the endpoint to authenticate that HPE DSCC is the caller.- `BEARER_TOKEN`: A customer provided token that is added to the `Authorization` HTTP header with the bearer authentication scheme.
type WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type int

const (
    BEARER_TOKEN_WEBHOOKSGETRESPONSE_ITEMS_DESTINATIONS_CLOUDEVENTENDPOINT_AUTHENTICATION_TYPE WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type = iota
)

func (i WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type) String() string {
    return []string{"BEARER_TOKEN"}[i]
}
func ParseWebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type(v string) (any, error) {
    result := BEARER_TOKEN_WEBHOOKSGETRESPONSE_ITEMS_DESTINATIONS_CLOUDEVENTENDPOINT_AUTHENTICATION_TYPE
    switch v {
        case "BEARER_TOKEN":
            result = BEARER_TOKEN_WEBHOOKSGETRESPONSE_ITEMS_DESTINATIONS_CLOUDEVENTENDPOINT_AUTHENTICATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type(values []WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksGetResponse_items_destinations_cloudEventEndpoint_authentication_type) isMultiValue() bool {
    return false
}
