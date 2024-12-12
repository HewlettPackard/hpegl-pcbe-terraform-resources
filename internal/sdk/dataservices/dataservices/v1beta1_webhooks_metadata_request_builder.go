package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1WebhooksMetadataRequestBuilder builds and executes requests for operations under \data-services\v1beta1\webhooks-metadata
type V1beta1WebhooksMetadataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1WebhooksMetadataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksMetadataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1WebhooksMetadataRequestBuilderInternal instantiates a new V1beta1WebhooksMetadataRequestBuilder and sets the default values.
func NewV1beta1WebhooksMetadataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksMetadataRequestBuilder) {
    m := &V1beta1WebhooksMetadataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/webhooks-metadata", pathParameters),
    }
    return m
}
// NewV1beta1WebhooksMetadataRequestBuilder instantiates a new V1beta1WebhooksMetadataRequestBuilder and sets the default values.
func NewV1beta1WebhooksMetadataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksMetadataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1WebhooksMetadataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all Webhook metadata. This includes the list of values for each of the webhook enum attributes.It also includes the list of resources that have been onboarded with RTM.
// Deprecated: This method is obsolete. Use GetAsWebhooksMetadataGetResponse instead.
// returns a V1beta1WebhooksMetadataResponseable when successful
// returns a V1beta1WebhooksMetadata400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksMetadata401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksMetadata403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksMetadata500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksMetadata503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksMetadataRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1WebhooksMetadataRequestBuilderGetRequestConfiguration)(V1beta1WebhooksMetadataResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksMetadata400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksMetadata401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksMetadata403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksMetadata500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksMetadata503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksMetadataResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksMetadataResponseable), nil
}
// GetAsWebhooksMetadataGetResponse returns all Webhook metadata. This includes the list of values for each of the webhook enum attributes.It also includes the list of resources that have been onboarded with RTM.
// returns a V1beta1WebhooksMetadataGetResponseable when successful
// returns a V1beta1WebhooksMetadata400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksMetadata401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksMetadata403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksMetadata500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksMetadata503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksMetadataRequestBuilder) GetAsWebhooksMetadataGetResponse(ctx context.Context, requestConfiguration *V1beta1WebhooksMetadataRequestBuilderGetRequestConfiguration)(V1beta1WebhooksMetadataGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksMetadata400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksMetadata401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksMetadata403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksMetadata500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksMetadata503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksMetadataGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksMetadataGetResponseable), nil
}
// ToGetRequestInformation returns all Webhook metadata. This includes the list of values for each of the webhook enum attributes.It also includes the list of resources that have been onboarded with RTM.
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksMetadataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1WebhooksMetadataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1WebhooksMetadataRequestBuilder when successful
func (m *V1beta1WebhooksMetadataRequestBuilder) WithUrl(rawUrl string)(*V1beta1WebhooksMetadataRequestBuilder) {
    return NewV1beta1WebhooksMetadataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
