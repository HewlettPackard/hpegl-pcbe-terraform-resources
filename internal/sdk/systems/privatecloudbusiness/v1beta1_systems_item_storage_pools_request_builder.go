package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemStoragePoolsRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\storage-pools
type V1beta1SystemsItemStoragePoolsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemStoragePoolsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemStoragePoolsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemStoragePoolsRequestBuilderInternal instantiates a new V1beta1SystemsItemStoragePoolsRequestBuilder and sets the default values.
func NewV1beta1SystemsItemStoragePoolsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemStoragePoolsRequestBuilder) {
    m := &V1beta1SystemsItemStoragePoolsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/storage-pools", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemStoragePoolsRequestBuilder instantiates a new V1beta1SystemsItemStoragePoolsRequestBuilder and sets the default values.
func NewV1beta1SystemsItemStoragePoolsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemStoragePoolsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemStoragePoolsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists storage pools of system's storage system in system details page view.Does not support Selection, Filtering and Sorting.
// Deprecated: This method is obsolete. Use GetAsStoragePoolsGetResponse instead.
// returns a V1beta1SystemsItemStoragePoolsResponseable when successful
// returns a V1beta1SystemsItemStoragePools400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemStoragePools401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemStoragePools404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemStoragePools500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemStoragePoolsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemStoragePoolsRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemStoragePoolsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemStoragePools400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemStoragePools401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemStoragePools404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemStoragePools500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemStoragePoolsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemStoragePoolsResponseable), nil
}
// GetAsStoragePoolsGetResponse lists storage pools of system's storage system in system details page view.Does not support Selection, Filtering and Sorting.
// returns a V1beta1SystemsItemStoragePoolsGetResponseable when successful
// returns a V1beta1SystemsItemStoragePools400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemStoragePools401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemStoragePools404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemStoragePools500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemStoragePoolsRequestBuilder) GetAsStoragePoolsGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemStoragePoolsRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemStoragePoolsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemStoragePools400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemStoragePools401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemStoragePools404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemStoragePools500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemStoragePoolsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemStoragePoolsGetResponseable), nil
}
// ToGetRequestInformation lists storage pools of system's storage system in system details page view.Does not support Selection, Filtering and Sorting.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemStoragePoolsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemStoragePoolsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemsItemStoragePoolsRequestBuilder when successful
func (m *V1beta1SystemsItemStoragePoolsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemStoragePoolsRequestBuilder) {
    return NewV1beta1SystemsItemStoragePoolsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
