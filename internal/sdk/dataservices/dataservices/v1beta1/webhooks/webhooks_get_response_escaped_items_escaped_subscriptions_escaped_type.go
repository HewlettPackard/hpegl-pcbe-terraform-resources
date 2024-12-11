package webhooks
// An enum indicating the subscription type- `RESOURCE_TYPE`: Matches all events for a resource type
type WebhooksGetResponse_items_subscriptions_type int

const (
    RESOURCE_TYPE_WEBHOOKSGETRESPONSE_ITEMS_SUBSCRIPTIONS_TYPE WebhooksGetResponse_items_subscriptions_type = iota
)

func (i WebhooksGetResponse_items_subscriptions_type) String() string {
    return []string{"RESOURCE_TYPE"}[i]
}
func ParseWebhooksGetResponse_items_subscriptions_type(v string) (any, error) {
    result := RESOURCE_TYPE_WEBHOOKSGETRESPONSE_ITEMS_SUBSCRIPTIONS_TYPE
    switch v {
        case "RESOURCE_TYPE":
            result = RESOURCE_TYPE_WEBHOOKSGETRESPONSE_ITEMS_SUBSCRIPTIONS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksGetResponse_items_subscriptions_type(values []WebhooksGetResponse_items_subscriptions_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksGetResponse_items_subscriptions_type) isMultiValue() bool {
    return false
}