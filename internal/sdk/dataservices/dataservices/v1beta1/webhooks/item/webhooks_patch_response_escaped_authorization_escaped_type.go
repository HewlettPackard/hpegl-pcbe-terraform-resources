package item
// An enum indicating the type of identity that is used to authorize published events.The provided identity must have permission to the events.- `CLIENT_CREDENTIALS`: A GLCP client credential, *client id* and *client secret* pair.
type WebhooksPatchResponse_authorization_type int

const (
    CLIENT_CREDENTIALS_WEBHOOKSPATCHRESPONSE_AUTHORIZATION_TYPE WebhooksPatchResponse_authorization_type = iota
)

func (i WebhooksPatchResponse_authorization_type) String() string {
    return []string{"CLIENT_CREDENTIALS"}[i]
}
func ParseWebhooksPatchResponse_authorization_type(v string) (any, error) {
    result := CLIENT_CREDENTIALS_WEBHOOKSPATCHRESPONSE_AUTHORIZATION_TYPE
    switch v {
        case "CLIENT_CREDENTIALS":
            result = CLIENT_CREDENTIALS_WEBHOOKSPATCHRESPONSE_AUTHORIZATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksPatchResponse_authorization_type(values []WebhooksPatchResponse_authorization_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksPatchResponse_authorization_type) isMultiValue() bool {
    return false
}
