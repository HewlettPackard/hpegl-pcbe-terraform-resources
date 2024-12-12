package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1WebhooksTestConnectionRequestBuilder builds and executes requests for operations under \data-services\v1beta1\webhooks\test-connection
type V1beta1WebhooksTestConnectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1WebhooksTestConnectionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1WebhooksTestConnectionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1WebhooksTestConnectionRequestBuilderInternal instantiates a new V1beta1WebhooksTestConnectionRequestBuilder and sets the default values.
func NewV1beta1WebhooksTestConnectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksTestConnectionRequestBuilder) {
    m := &V1beta1WebhooksTestConnectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/webhooks/test-connection", pathParameters),
    }
    return m
}
// NewV1beta1WebhooksTestConnectionRequestBuilder instantiates a new V1beta1WebhooksTestConnectionRequestBuilder and sets the default values.
func NewV1beta1WebhooksTestConnectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1WebhooksTestConnectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1WebhooksTestConnectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post tests the connection to a webhook endpoint to ensure that the destinationparameters are set correctly.
// Deprecated: This method is obsolete. Use PostAsTestConnectionPostResponse instead.
// returns a V1beta1WebhooksTestConnectionResponseable when successful
// returns a V1beta1WebhooksTestConnection400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksTestConnection401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksTestConnection403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksTestConnection500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksTestConnection503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksTestConnectionRequestBuilder) Post(ctx context.Context, body V1beta1WebhooksTestConnectionPostRequestBodyable, requestConfiguration *V1beta1WebhooksTestConnectionRequestBuilderPostRequestConfiguration)(V1beta1WebhooksTestConnectionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksTestConnection400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksTestConnection401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksTestConnection403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksTestConnection500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksTestConnection503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksTestConnectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksTestConnectionResponseable), nil
}
// PostAsTestConnectionPostResponse tests the connection to a webhook endpoint to ensure that the destinationparameters are set correctly.
// returns a V1beta1WebhooksTestConnectionPostResponseable when successful
// returns a V1beta1WebhooksTestConnection400Error error when the service returns a 400 status code
// returns a V1beta1WebhooksTestConnection401Error error when the service returns a 401 status code
// returns a V1beta1WebhooksTestConnection403Error error when the service returns a 403 status code
// returns a V1beta1WebhooksTestConnection500Error error when the service returns a 500 status code
// returns a V1beta1WebhooksTestConnection503Error error when the service returns a 503 status code
func (m *V1beta1WebhooksTestConnectionRequestBuilder) PostAsTestConnectionPostResponse(ctx context.Context, body V1beta1WebhooksTestConnectionPostRequestBodyable, requestConfiguration *V1beta1WebhooksTestConnectionRequestBuilderPostRequestConfiguration)(V1beta1WebhooksTestConnectionPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1WebhooksTestConnection400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1WebhooksTestConnection401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1WebhooksTestConnection403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1WebhooksTestConnection500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1WebhooksTestConnection503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1WebhooksTestConnectionPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1WebhooksTestConnectionPostResponseable), nil
}
// ToPostRequestInformation tests the connection to a webhook endpoint to ensure that the destinationparameters are set correctly.
// returns a *RequestInformation when successful
func (m *V1beta1WebhooksTestConnectionRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1WebhooksTestConnectionPostRequestBodyable, requestConfiguration *V1beta1WebhooksTestConnectionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1WebhooksTestConnectionRequestBuilder when successful
func (m *V1beta1WebhooksTestConnectionRequestBuilder) WithUrl(rawUrl string)(*V1beta1WebhooksTestConnectionRequestBuilder) {
    return NewV1beta1WebhooksTestConnectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
