package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1WebhooksWebhooksItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\webhooks\{id}
type V1beta1WebhooksWebhooksItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1WebhooksWebhooksItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksWebhooksItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1WebhooksWebhooksItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksWebhooksItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1WebhooksWebhooksItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksWebhooksItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1WebhooksWebhooksItemRequestBuilderInternal instantiates a new V1beta1WebhooksWebhooksItemRequestBuilder and sets the default values.
func NewV1beta1WebhooksWebhooksItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksWebhooksItemRequestBuilder) {
    m := &V1beta1WebhooksWebhooksItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/webhooks/{id}", pathParameters),
    }
    return m
}
// NewV1beta1WebhooksWebhooksItemRequestBuilder instantiates a new V1beta1WebhooksWebhooksItemRequestBuilder and sets the default values.
func NewV1beta1WebhooksWebhooksItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksWebhooksItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1WebhooksWebhooksItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the webhook having the given unique identifier.
// returns a V1beta1WebhooksItemWebhooks400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksItemWebhooks401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksItemWebhooks403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksItemWebhooks404Error error when the service returns a 404 status code
// returns a V1beta1WebhooksItemWebhooks500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksItemWebhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksItemWebhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksItemWebhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksItemWebhooks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1WebhooksItemWebhooks404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksItemWebhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksItemWebhooks503ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns details of a webhook having the given unique identifier.
// Deprecated: This method is obsolete. Use GetAsWebhooksGetResponse instead.
// returns a V1beta1WebhooksItemWebhooksResponseable when successful
// returns a V1beta1WebhooksItemWebhooks400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksItemWebhooks401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksItemWebhooks403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksItemWebhooks404Error error when the service returns a 404 status code
// returns a V1beta1WebhooksItemWebhooks500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksItemWebhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderGetRequestConfiguration)(V1beta1WebhooksItemWebhooksResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksItemWebhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksItemWebhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksItemWebhooks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1WebhooksItemWebhooks404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksItemWebhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksItemWebhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksItemWebhooksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksItemWebhooksResponseable), nil
}
// GetAsWebhooksGetResponse returns details of a webhook having the given unique identifier.
// returns a V1beta1WebhooksItemWebhooksGetResponseable when successful
// returns a V1beta1WebhooksItemWebhooks400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksItemWebhooks401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksItemWebhooks403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksItemWebhooks404Error error when the service returns a 404 status code
// returns a V1beta1WebhooksItemWebhooks500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksItemWebhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) GetAsWebhooksGetResponse(ctx context.Context, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderGetRequestConfiguration)(V1beta1WebhooksItemWebhooksGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksItemWebhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksItemWebhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksItemWebhooks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1WebhooksItemWebhooks404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksItemWebhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksItemWebhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksItemWebhooksGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksItemWebhooksGetResponseable), nil
}
// Patch updates the webhook having the given unique identifier.Request body format is [RFC-7396 JSON Merge Patch](https://datatracker.ietf.org/doc/html/rfc7396).
// Deprecated: This method is obsolete. Use PatchAsWebhooksPatchResponse instead.
// returns a V1beta1WebhooksItemWebhooksResponseable when successful
// returns a V1beta1WebhooksItemWebhooks400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksItemWebhooks401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksItemWebhooks403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksItemWebhooks404Error error when the service returns a 404 status code
// returns a V1beta1WebhooksItemWebhooks500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksItemWebhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) Patch(ctx context.Context, body V1beta1WebhooksItemWebhooksPatchRequestBodyable, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderPatchRequestConfiguration)(V1beta1WebhooksItemWebhooksResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksItemWebhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksItemWebhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksItemWebhooks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1WebhooksItemWebhooks404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksItemWebhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksItemWebhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksItemWebhooksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksItemWebhooksResponseable), nil
}
// PatchAsWebhooksPatchResponse updates the webhook having the given unique identifier.Request body format is [RFC-7396 JSON Merge Patch](https://datatracker.ietf.org/doc/html/rfc7396).
// returns a V1beta1WebhooksItemWebhooksPatchResponseable when successful
// returns a V1beta1WebhooksItemWebhooks400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksItemWebhooks401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksItemWebhooks403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksItemWebhooks404Error error when the service returns a 404 status code
// returns a V1beta1WebhooksItemWebhooks500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksItemWebhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) PatchAsWebhooksPatchResponse(ctx context.Context, body V1beta1WebhooksItemWebhooksPatchRequestBodyable, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderPatchRequestConfiguration)(V1beta1WebhooksItemWebhooksPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksItemWebhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksItemWebhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksItemWebhooks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1WebhooksItemWebhooks404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksItemWebhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksItemWebhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksItemWebhooksPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksItemWebhooksPatchResponseable), nil
}
// ToDeleteRequestInformation deletes the webhook having the given unique identifier.
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns details of a webhook having the given unique identifier.
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates the webhook having the given unique identifier.Request body format is [RFC-7396 JSON Merge Patch](https://datatracker.ietf.org/doc/html/rfc7396).
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1WebhooksItemWebhooksPatchRequestBodyable, requestConfiguration *V1beta1WebhooksWebhooksItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/merge-patch+json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1WebhooksWebhooksItemRequestBuilder when successful
func (m *V1beta1WebhooksWebhooksItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1WebhooksWebhooksItemRequestBuilder) {
    return NewV1beta1WebhooksWebhooksItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
