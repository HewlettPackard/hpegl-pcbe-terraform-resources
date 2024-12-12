package webhooks
// An enum indicating the type of endpoint that events are published to.- `CLOUD_EVENT`: A customer controlled HTTPS endpoint to which DSCC cloud events are posted.
type WebhooksPostResponse_destinations_type int

const (
    CLOUD_EVENT_WEBHOOKSPOSTRESPONSE_DESTINATIONS_TYPE WebhooksPostResponse_destinations_type = iota
)

func (i WebhooksPostResponse_destinations_type) String() string {
    return []string{"CLOUD_EVENT"}[i]
}
func ParseWebhooksPostResponse_destinations_type(v string) (any, error) {
    result := CLOUD_EVENT_WEBHOOKSPOSTRESPONSE_DESTINATIONS_TYPE
    switch v {
        case "CLOUD_EVENT":
            result = CLOUD_EVENT_WEBHOOKSPOSTRESPONSE_DESTINATIONS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPostResponse_destinations_type(values []WebhooksPostResponse_destinations_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPostResponse_destinations_type) isMultiValue() bool {
    return false
}
