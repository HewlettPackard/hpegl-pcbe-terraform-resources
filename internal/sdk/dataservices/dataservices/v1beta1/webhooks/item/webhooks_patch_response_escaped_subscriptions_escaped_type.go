package item
// An enum indicating the subscription type- `RESOURCE_TYPE`: Matches all events for a resource type
type WebhooksPatchResponse_subscriptions_type int

const (
    RESOURCE_TYPE_WEBHOOKSPATCHRESPONSE_SUBSCRIPTIONS_TYPE WebhooksPatchResponse_subscriptions_type = iota
)

func (i WebhooksPatchResponse_subscriptions_type) String() string {
    return []string{"RESOURCE_TYPE"}[i]
}
func ParseWebhooksPatchResponse_subscriptions_type(v string) (any, error) {
    result := RESOURCE_TYPE_WEBHOOKSPATCHRESPONSE_SUBSCRIPTIONS_TYPE
    switch v {
        case "RESOURCE_TYPE":
            result = RESOURCE_TYPE_WEBHOOKSPATCHRESPONSE_SUBSCRIPTIONS_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPatchResponse_subscriptions_type(values []WebhooksPatchResponse_subscriptions_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPatchResponse_subscriptions_type) isMultiValue() bool {
    return false
}
