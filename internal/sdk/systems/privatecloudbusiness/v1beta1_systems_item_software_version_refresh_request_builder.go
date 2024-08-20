package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\software-version-refresh
type V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSoftwareVersionRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSoftwareVersionRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilderInternal instantiates a new V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) {
    m := &V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/software-version-refresh", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilder instantiates a new V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post refresh the software catalogs and software versions on the specified system.The current and available software catalog versions are recomputed based on the component versions running on the system and software catalogs available for update.
// Deprecated: This method is obsolete. Use PostAsSoftwareVersionRefreshPostResponse instead.
// returns a V1beta1SystemsItemSoftwareVersionRefreshResponseable when successful
// returns a V1beta1SystemsItemSoftwareVersionRefresh400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareVersionRefreshResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareVersionRefresh400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareVersionRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareVersionRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareVersionRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareVersionRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareVersionRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareVersionRefreshResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareVersionRefreshResponseable), nil
}
// PostAsSoftwareVersionRefreshPostResponse refresh the software catalogs and software versions on the specified system.The current and available software catalog versions are recomputed based on the component versions running on the system and software catalogs available for update.
// returns a V1beta1SystemsItemSoftwareVersionRefreshPostResponseable when successful
// returns a V1beta1SystemsItemSoftwareVersionRefresh400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareVersionRefresh500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) PostAsSoftwareVersionRefreshPostResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareVersionRefreshPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareVersionRefresh400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareVersionRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareVersionRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareVersionRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareVersionRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareVersionRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareVersionRefreshPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareVersionRefreshPostResponseable), nil
}
// ToPostRequestInformation refresh the software catalogs and software versions on the specified system.The current and available software catalog versions are recomputed based on the component versions running on the system and software catalogs available for update.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder when successful
func (m *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
