package item
// An enum indicating current webhook event delivery status- `UNKNOWN`: There have been no events to deliver yet, so delivery status cannot be determined- `OK`: Events are being successfully delivered to the webhook destination endpoint- `WARNING`: There was a recent problem delivering events to the webhook destination endpoint- `ERROR`: There is a current problem delivering events to the webhook destination endpoint
type WebhooksGetResponse_status int

const (
    UNKNOWN_WEBHOOKSGETRESPONSE_STATUS WebhooksGetResponse_status = iota
    OK_WEBHOOKSGETRESPONSE_STATUS
    WARNING_WEBHOOKSGETRESPONSE_STATUS
    ERROR_WEBHOOKSGETRESPONSE_STATUS
)

func (i WebhooksGetResponse_status) String() string {
    return []string{"UNKNOWN", "OK", "WARNING", "ERROR"}[i]
}
func ParseWebhooksGetResponse_status(v string) (any, error) {
    result := UNKNOWN_WEBHOOKSGETRESPONSE_STATUS
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_WEBHOOKSGETRESPONSE_STATUS
        case "OK":
            result = OK_WEBHOOKSGETRESPONSE_STATUS
        case "WARNING":
            result = WARNING_WEBHOOKSGETRESPONSE_STATUS
        case "ERROR":
            result = ERROR_WEBHOOKSGETRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksGetResponse_status(values []WebhooksGetResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksGetResponse_status) isMultiValue() bool {
    return false
}
