package item
// An enum indicating the type of identity that is used to authorize published events.The provided identity must have permission to the events.- `CLIENT_CREDENTIALS`: A GLCP client credential, *client id* and *client secret* pair.
type WebhooksGetResponse_authorization_type int

const (
    CLIENT_CREDENTIALS_WEBHOOKSGETRESPONSE_AUTHORIZATION_TYPE WebhooksGetResponse_authorization_type = iota
)

func (i WebhooksGetResponse_authorization_type) String() string {
    return []string{"CLIENT_CREDENTIALS"}[i]
}
func ParseWebhooksGetResponse_authorization_type(v string) (any, error) {
    result := CLIENT_CREDENTIALS_WEBHOOKSGETRESPONSE_AUTHORIZATION_TYPE
    switch v {
        case "CLIENT_CREDENTIALS":
            result = CLIENT_CREDENTIALS_WEBHOOKSGETRESPONSE_AUTHORIZATION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhooksGetResponse_authorization_type(values []WebhooksGetResponse_authorization_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhooksGetResponse_authorization_type) isMultiValue() bool {
    return false
}
