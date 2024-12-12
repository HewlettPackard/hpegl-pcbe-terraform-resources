package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1StorageLocationsStorageLocationsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\storage-locations\{id}
type V1beta1StorageLocationsStorageLocationsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1StorageLocationsStorageLocationsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1StorageLocationsStorageLocationsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1StorageLocationsStorageLocationsItemRequestBuilderInternal instantiates a new V1beta1StorageLocationsStorageLocationsItemRequestBuilder and sets the default values.
func NewV1beta1StorageLocationsStorageLocationsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1StorageLocationsStorageLocationsItemRequestBuilder) {
    m := &V1beta1StorageLocationsStorageLocationsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/storage-locations/{id}", pathParameters),
    }
    return m
}
// NewV1beta1StorageLocationsStorageLocationsItemRequestBuilder instantiates a new V1beta1StorageLocationsStorageLocationsItemRequestBuilder and sets the default values.
func NewV1beta1StorageLocationsStorageLocationsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1StorageLocationsStorageLocationsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1StorageLocationsStorageLocationsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns details of the storage location that corresponds to the provided ID.
// Deprecated: This method is obsolete. Use GetAsStorageLocationsGetResponse instead.
// returns a V1beta1StorageLocationsItemStorageLocationsResponseable when successful
// returns a V1beta1StorageLocationsItemStorageLocations400Error error when the service returns a 400 status code
// returns a V1beta1StorageLocationsItemStorageLocations401Error error when the service returns a 401 status code
// returns a V1beta1StorageLocationsItemStorageLocations403Error error when the service returns a 403 status code
// returns a V1beta1StorageLocationsItemStorageLocations404Error error when the service returns a 404 status code
// returns a V1beta1StorageLocationsItemStorageLocations500Error error when the service returns a 500 status code
// returns a V1beta1StorageLocationsItemStorageLocations503Error error when the service returns a 503 status code
func (m *V1beta1StorageLocationsStorageLocationsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1StorageLocationsStorageLocationsItemRequestBuilderGetRequestConfiguration)(V1beta1StorageLocationsItemStorageLocationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1StorageLocationsItemStorageLocations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1StorageLocationsItemStorageLocations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1StorageLocationsItemStorageLocations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1StorageLocationsItemStorageLocations404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1StorageLocationsItemStorageLocations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1StorageLocationsItemStorageLocations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1StorageLocationsItemStorageLocationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1StorageLocationsItemStorageLocationsResponseable), nil
}
// GetAsStorageLocationsGetResponse returns details of the storage location that corresponds to the provided ID.
// returns a V1beta1StorageLocationsItemStorageLocationsGetResponseable when successful
// returns a V1beta1StorageLocationsItemStorageLocations400Error error when the service returns a 400 status code
// returns a V1beta1StorageLocationsItemStorageLocations401Error error when the service returns a 401 status code
// returns a V1beta1StorageLocationsItemStorageLocations403Error error when the service returns a 403 status code
// returns a V1beta1StorageLocationsItemStorageLocations404Error error when the service returns a 404 status code
// returns a V1beta1StorageLocationsItemStorageLocations500Error error when the service returns a 500 status code
// returns a V1beta1StorageLocationsItemStorageLocations503Error error when the service returns a 503 status code
func (m *V1beta1StorageLocationsStorageLocationsItemRequestBuilder) GetAsStorageLocationsGetResponse(ctx context.Context, requestConfiguration *V1beta1StorageLocationsStorageLocationsItemRequestBuilderGetRequestConfiguration)(V1beta1StorageLocationsItemStorageLocationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1StorageLocationsItemStorageLocations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1StorageLocationsItemStorageLocations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1StorageLocationsItemStorageLocations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1StorageLocationsItemStorageLocations404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1StorageLocationsItemStorageLocations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1StorageLocationsItemStorageLocations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1StorageLocationsItemStorageLocationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1StorageLocationsItemStorageLocationsGetResponseable), nil
}
// ToGetRequestInformation returns details of the storage location that corresponds to the provided ID.
// returns a *RequestInformation when successful
func (m *V1beta1StorageLocationsStorageLocationsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1StorageLocationsStorageLocationsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1StorageLocationsStorageLocationsItemRequestBuilder when successful
func (m *V1beta1StorageLocationsStorageLocationsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1StorageLocationsStorageLocationsItemRequestBuilder) {
    return NewV1beta1StorageLocationsStorageLocationsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
