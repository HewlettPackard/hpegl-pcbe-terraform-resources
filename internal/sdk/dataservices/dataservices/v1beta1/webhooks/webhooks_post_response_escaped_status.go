package webhooks
// An enum indicating current webhook event delivery status- `UNKNOWN`: There have been no events to deliver yet, so delivery status cannot be determined- `OK`: Events are being successfully delivered to the webhook destination endpoint- `WARNING`: There was a recent problem delivering events to the webhook destination endpoint- `ERROR`: There is a current problem delivering events to the webhook destination endpoint
type WebhooksPostResponse_status int

const (
    UNKNOWN_WEBHOOKSPOSTRESPONSE_STATUS WebhooksPostResponse_status = iota
    OK_WEBHOOKSPOSTRESPONSE_STATUS
    WARNING_WEBHOOKSPOSTRESPONSE_STATUS
    ERROR_WEBHOOKSPOSTRESPONSE_STATUS
)

func (i WebhooksPostResponse_status) String() string {
    return []string{"UNKNOWN", "OK", "WARNING", "ERROR"}[i]
}
func ParseWebhooksPostResponse_status(v string) (any, error) {
    result := UNKNOWN_WEBHOOKSPOSTRESPONSE_STATUS
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_WEBHOOKSPOSTRESPONSE_STATUS
        case "OK":
            result = OK_WEBHOOKSPOSTRESPONSE_STATUS
        case "WARNING":
            result = WARNING_WEBHOOKSPOSTRESPONSE_STATUS
        case "ERROR":
            result = ERROR_WEBHOOKSPOSTRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPostResponse_status(values []WebhooksPostResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPostResponse_status) isMultiValue() bool {
    return false
}
