package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemStorageArraysRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\storage-arrays
type V1beta1SystemsItemStorageArraysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemStorageArraysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemStorageArraysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemStorageArraysRequestBuilderInternal instantiates a new V1beta1SystemsItemStorageArraysRequestBuilder and sets the default values.
func NewV1beta1SystemsItemStorageArraysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemStorageArraysRequestBuilder) {
    m := &V1beta1SystemsItemStorageArraysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/storage-arrays", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemStorageArraysRequestBuilder instantiates a new V1beta1SystemsItemStorageArraysRequestBuilder and sets the default values.
func NewV1beta1SystemsItemStorageArraysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemStorageArraysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemStorageArraysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists storage arrays of system's storage system in system details page view.
// Deprecated: This method is obsolete. Use GetAsStorageArraysGetResponse instead.
// returns a V1beta1SystemsItemStorageArraysResponseable when successful
// returns a V1beta1SystemsItemStorageArrays400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemStorageArrays401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemStorageArrays404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemStorageArrays500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemStorageArraysRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemStorageArraysRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemStorageArraysResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemStorageArrays400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemStorageArrays401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemStorageArrays404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemStorageArrays500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemStorageArraysResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemStorageArraysResponseable), nil
}
// GetAsStorageArraysGetResponse lists storage arrays of system's storage system in system details page view.
// returns a V1beta1SystemsItemStorageArraysGetResponseable when successful
// returns a V1beta1SystemsItemStorageArrays400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemStorageArrays401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemStorageArrays404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemStorageArrays500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemStorageArraysRequestBuilder) GetAsStorageArraysGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemStorageArraysRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemStorageArraysGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemStorageArrays400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemStorageArrays401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemStorageArrays404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemStorageArrays500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemStorageArraysGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemStorageArraysGetResponseable), nil
}
// ToGetRequestInformation lists storage arrays of system's storage system in system details page view.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemStorageArraysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemStorageArraysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemsItemStorageArraysRequestBuilder when successful
func (m *V1beta1SystemsItemStorageArraysRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemStorageArraysRequestBuilder) {
    return NewV1beta1SystemsItemStorageArraysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
