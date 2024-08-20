package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemRefreshRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\refresh
type V1beta1CspAccountsItemRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemRefreshRequestBuilderInternal instantiates a new V1beta1CspAccountsItemRefreshRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemRefreshRequestBuilder) {
    m := &V1beta1CspAccountsItemRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/refresh", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemRefreshRequestBuilder instantiates a new V1beta1CspAccountsItemRefreshRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates the resource inventory with the latest information in a cloud account.
// Deprecated: This method is obsolete. Use PostAsRefreshPostResponse instead.
// returns a V1beta1CspAccountsItemRefreshResponseable when successful
// returns a V1beta1CspAccountsItemRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemRefreshRequestBuilderPostRequestConfiguration)(V1beta1CspAccountsItemRefreshResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspAccountsItemRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemRefreshResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemRefreshResponseable), nil
}
// PostAsRefreshPostResponse updates the resource inventory with the latest information in a cloud account.
// returns a V1beta1CspAccountsItemRefreshPostResponseable when successful
// returns a V1beta1CspAccountsItemRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemRefreshRequestBuilder) PostAsRefreshPostResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemRefreshRequestBuilderPostRequestConfiguration)(V1beta1CspAccountsItemRefreshPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspAccountsItemRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemRefreshPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemRefreshPostResponseable), nil
}
// ToPostRequestInformation updates the resource inventory with the latest information in a cloud account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemRefreshRequestBuilder when successful
func (m *V1beta1CspAccountsItemRefreshRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemRefreshRequestBuilder) {
    return NewV1beta1CspAccountsItemRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
