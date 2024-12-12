package item
// An enum indicating current webhook event delivery status- `UNKNOWN`: There have been no events to deliver yet, so delivery status cannot be determined- `OK`: Events are being successfully delivered to the webhook destination endpoint- `WARNING`: There was a recent problem delivering events to the webhook destination endpoint- `ERROR`: There is a current problem delivering events to the webhook destination endpoint
type WebhooksPatchResponse_status int

const (
    UNKNOWN_WEBHOOKSPATCHRESPONSE_STATUS WebhooksPatchResponse_status = iota
    OK_WEBHOOKSPATCHRESPONSE_STATUS
    WARNING_WEBHOOKSPATCHRESPONSE_STATUS
    ERROR_WEBHOOKSPATCHRESPONSE_STATUS
)

func (i WebhooksPatchResponse_status) String() string {
    return []string{"UNKNOWN", "OK", "WARNING", "ERROR"}[i]
}
func ParseWebhooksPatchResponse_status(v string) (any, error) {
    result := UNKNOWN_WEBHOOKSPATCHRESPONSE_STATUS
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_WEBHOOKSPATCHRESPONSE_STATUS
        case "OK":
            result = OK_WEBHOOKSPATCHRESPONSE_STATUS
        case "WARNING":
            result = WARNING_WEBHOOKSPATCHRESPONSE_STATUS
        case "ERROR":
            result = ERROR_WEBHOOKSPATCHRESPONSE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPatchResponse_status(values []WebhooksPatchResponse_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPatchResponse_status) isMultiValue() bool {
    return false
}
