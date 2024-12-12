package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1WebhooksRequestBuilder builds and executes requests for operations under \data-services\v1beta1\webhooks
type V1beta1WebhooksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1WebhooksRequestBuilderGetQueryParameters returns a list of Webhook resources that are visible to the user.The response can be paged by using the `limit` and `offset` query parameters
type V1beta1WebhooksRequestBuilderGetQueryParameters struct {
    // Limit is the maximum number of items to include in the response.If not specified, The default limit is 10.Limit should be used in conjunction with `offset`.```offset=30&limit=10```
    Limit *int32 `uriparametername:"limit"`
    // Offset is the number of items from the beginning of the result set tothe first item included in the response.If not specified, the default offset is 0.Offset should be used in conjunction with `limit`.```offset=30&limit=10```
    Offset *int32 `uriparametername:"offset"`
}
// V1beta1WebhooksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1WebhooksRequestBuilderGetQueryParameters
}
// V1beta1WebhooksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.webhooks.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1WebhooksWebhooksItemRequestBuilder when successful
func (m *V1beta1WebhooksRequestBuilder) ById(id string)(*V1beta1WebhooksWebhooksItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1WebhooksWebhooksItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.webhooks.item collection
// returns a *V1beta1WebhooksWebhooksItemRequestBuilder when successful
func (m *V1beta1WebhooksRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1WebhooksWebhooksItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1WebhooksWebhooksItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1WebhooksRequestBuilderInternal instantiates a new V1beta1WebhooksRequestBuilder and sets the default values.
func NewV1beta1WebhooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksRequestBuilder) {
    m := &V1beta1WebhooksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/webhooks{?limit*,offset*}", pathParameters),
    }
    return m
}
// NewV1beta1WebhooksRequestBuilder instantiates a new V1beta1WebhooksRequestBuilder and sets the default values.
func NewV1beta1WebhooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1WebhooksRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of Webhook resources that are visible to the user.The response can be paged by using the `limit` and `offset` query parameters
// Deprecated: This method is obsolete. Use GetAsWebhooksGetResponse instead.
// returns a V1beta1WebhooksResponseable when successful
// returns a V1beta1Webhooks400Error error when the service returns a 400 status code
// returns a V1beta1Webhooks401Error error when the service returns a 401 status code
// returns a V1beta1Webhooks403Error error when the service returns a 403 status code
// returns a V1beta1Webhooks500Error error when the service returns a 500 status code
// returns a V1beta1Webhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1WebhooksRequestBuilderGetRequestConfiguration)(V1beta1WebhooksResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Webhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Webhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Webhooks403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Webhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Webhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksResponseable), nil
}
// GetAsWebhooksGetResponse returns a list of Webhook resources that are visible to the user.The response can be paged by using the `limit` and `offset` query parameters
// returns a V1beta1WebhooksGetResponseable when successful
// returns a V1beta1Webhooks400Error error when the service returns a 400 status code
// returns a V1beta1Webhooks401Error error when the service returns a 401 status code
// returns a V1beta1Webhooks403Error error when the service returns a 403 status code
// returns a V1beta1Webhooks500Error error when the service returns a 500 status code
// returns a V1beta1Webhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksRequestBuilder) GetAsWebhooksGetResponse(ctx context.Context, requestConfiguration *V1beta1WebhooksRequestBuilderGetRequestConfiguration)(V1beta1WebhooksGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Webhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Webhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Webhooks403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Webhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Webhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksGetResponseable), nil
}
// Post creates a Webhook.
// Deprecated: This method is obsolete. Use PostAsWebhooksPostResponse instead.
// returns a V1beta1WebhooksResponseable when successful
// returns a V1beta1Webhooks400Error error when the service returns a 400 status code
// returns a V1beta1Webhooks401Error error when the service returns a 401 status code
// returns a V1beta1Webhooks403Error error when the service returns a 403 status code
// returns a V1beta1Webhooks500Error error when the service returns a 500 status code
// returns a V1beta1Webhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksRequestBuilder) Post(ctx context.Context, body V1beta1WebhooksPostRequestBodyable, requestConfiguration *V1beta1WebhooksRequestBuilderPostRequestConfiguration)(V1beta1WebhooksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Webhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Webhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Webhooks403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Webhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Webhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksResponseable), nil
}
// PostAsWebhooksPostResponse creates a Webhook.
// returns a V1beta1WebhooksPostResponseable when successful
// returns a V1beta1Webhooks400Error error when the service returns a 400 status code
// returns a V1beta1Webhooks401Error error when the service returns a 401 status code
// returns a V1beta1Webhooks403Error error when the service returns a 403 status code
// returns a V1beta1Webhooks500Error error when the service returns a 500 status code
// returns a V1beta1Webhooks503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksRequestBuilder) PostAsWebhooksPostResponse(ctx context.Context, body V1beta1WebhooksPostRequestBodyable, requestConfiguration *V1beta1WebhooksRequestBuilderPostRequestConfiguration)(V1beta1WebhooksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Webhooks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Webhooks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Webhooks403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Webhooks500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Webhooks503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksPostResponseable), nil
}
// TestConnection the testConnection property
// returns a *V1beta1WebhooksTestConnectionRequestBuilder when successful
func (m *V1beta1WebhooksRequestBuilder) TestConnection()(*V1beta1WebhooksTestConnectionRequestBuilder) {
    return NewV1beta1WebhooksTestConnectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns a list of Webhook resources that are visible to the user.The response can be paged by using the `limit` and `offset` query parameters
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1WebhooksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a Webhook.
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1WebhooksPostRequestBodyable, requestConfiguration *V1beta1WebhooksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1WebhooksRequestBuilder when successful
func (m *V1beta1WebhooksRequestBuilder) WithUrl(rawUrl string)(*V1beta1WebhooksRequestBuilder) {
    return NewV1beta1WebhooksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
