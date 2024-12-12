package webhooks
// An enum indicating the subscription type- `RESOURCE_TYPE`: Matches all events for a resource type
type WebhooksPostResponse_subscriptions_type int

const (
    RESOURCE_TYPE_WEBHOOKSPOSTRESPONSE_SUBSCRIPTIONS_TYPE WebhooksPostResponse_subscriptions_type = iota
)

func (i WebhooksPostResponse_subscriptions_type) String() string {
    return []string{"RESOURCE_TYPE"}[i]
}
func ParseWebhooksPostResponse_subscriptions_type(v string) (any, error) {
    result := RESOURCE_TYPE_WEBHOOKSPOSTRESPONSE_SUBSCRIPTIONS_TYPE
    switch v {
        case "RESOURCE_TYPE":
            result = RESOURCE_TYPE_WEBHOOKSPOSTRESPONSE_SUBSCRIPTIONS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPostResponse_subscriptions_type(values []WebhooksPostResponse_subscriptions_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPostResponse_subscriptions_type) isMultiValue() bool {
    return false
}
