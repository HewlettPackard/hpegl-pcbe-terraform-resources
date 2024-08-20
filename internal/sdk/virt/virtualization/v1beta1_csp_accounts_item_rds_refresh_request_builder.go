package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspAccountsItemRdsRefreshRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\rds-refresh
type V1beta1CspAccountsItemRdsRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemRdsRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemRdsRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemRdsRefreshRequestBuilderInternal instantiates a new V1beta1CspAccountsItemRdsRefreshRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemRdsRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemRdsRefreshRequestBuilder) {
    m := &V1beta1CspAccountsItemRdsRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/rds-refresh", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemRdsRefreshRequestBuilder instantiates a new V1beta1CspAccountsItemRdsRefreshRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemRdsRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemRdsRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemRdsRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates the resource inventory with the latest information in a cloud account.
// returns a UntypedNodeable when successful
// returns a V1beta1CspAccountsItemRdsRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemRdsRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemRdsRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemRdsRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemRdsRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemRdsRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemRdsRefreshRequestBuilderPostRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspAccountsItemRdsRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemRdsRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemRdsRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemRdsRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemRdsRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable), nil
}
// ToPostRequestInformation updates the resource inventory with the latest information in a cloud account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemRdsRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemRdsRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemRdsRefreshRequestBuilder when successful
func (m *V1beta1CspAccountsItemRdsRefreshRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemRdsRefreshRequestBuilder) {
    return NewV1beta1CspAccountsItemRdsRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
