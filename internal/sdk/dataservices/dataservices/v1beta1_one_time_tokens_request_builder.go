package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1OneTimeTokensRequestBuilder builds and executes requests for operations under \data-services\v1beta1\one-time-tokens
type V1beta1OneTimeTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1OneTimeTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1OneTimeTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1OneTimeTokensRequestBuilderInternal instantiates a new V1beta1OneTimeTokensRequestBuilder and sets the default values.
func NewV1beta1OneTimeTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1OneTimeTokensRequestBuilder) {
    m := &V1beta1OneTimeTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/one-time-tokens", pathParameters),
    }
    return m
}
// NewV1beta1OneTimeTokensRequestBuilder instantiates a new V1beta1OneTimeTokensRequestBuilder and sets the default values.
func NewV1beta1OneTimeTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1OneTimeTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1OneTimeTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a JWT that can be used to authenticate to DSCC API endpoints that accept one time use tokens. An API endpoint that accept one time use tokens will mention so in its documentation. API endpoints that do not declare acceptance of one time use tokens will not accept one for authentication.
// Deprecated: This method is obsolete. Use PostAsOneTimeTokensPostResponse instead.
// returns a V1beta1OneTimeTokensResponseable when successful
// returns a V1beta1OneTimeTokens400Error error when the service returns a 400 status code
// returns a V1beta1OneTimeTokens401Error error when the service returns a 401 status code
// returns a V1beta1OneTimeTokens500Error error when the service returns a 500 status code
// returns a V1beta1OneTimeTokens503Error error when the service returns a 503 status code
func (m *V1beta1OneTimeTokensRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1OneTimeTokensRequestBuilderPostRequestConfiguration)(V1beta1OneTimeTokensResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1OneTimeTokens400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1OneTimeTokens401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1OneTimeTokens500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1OneTimeTokens503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1OneTimeTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1OneTimeTokensResponseable), nil
}
// PostAsOneTimeTokensPostResponse returns a JWT that can be used to authenticate to DSCC API endpoints that accept one time use tokens. An API endpoint that accept one time use tokens will mention so in its documentation. API endpoints that do not declare acceptance of one time use tokens will not accept one for authentication.
// returns a V1beta1OneTimeTokensPostResponseable when successful
// returns a V1beta1OneTimeTokens400Error error when the service returns a 400 status code
// returns a V1beta1OneTimeTokens401Error error when the service returns a 401 status code
// returns a V1beta1OneTimeTokens500Error error when the service returns a 500 status code
// returns a V1beta1OneTimeTokens503Error error when the service returns a 503 status code
func (m *V1beta1OneTimeTokensRequestBuilder) PostAsOneTimeTokensPostResponse(ctx context.Context, requestConfiguration *V1beta1OneTimeTokensRequestBuilderPostRequestConfiguration)(V1beta1OneTimeTokensPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1OneTimeTokens400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1OneTimeTokens401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1OneTimeTokens500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1OneTimeTokens503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1OneTimeTokensPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1OneTimeTokensPostResponseable), nil
}
// ToPostRequestInformation returns a JWT that can be used to authenticate to DSCC API endpoints that accept one time use tokens. An API endpoint that accept one time use tokens will mention so in its documentation. API endpoints that do not declare acceptance of one time use tokens will not accept one for authentication.
// returns a *RequestInformation when successful
func (m *V1beta1OneTimeTokensRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1OneTimeTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1OneTimeTokensRequestBuilder when successful
func (m *V1beta1OneTimeTokensRequestBuilder) WithUrl(rawUrl string)(*V1beta1OneTimeTokensRequestBuilder) {
    return NewV1beta1OneTimeTokensRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
