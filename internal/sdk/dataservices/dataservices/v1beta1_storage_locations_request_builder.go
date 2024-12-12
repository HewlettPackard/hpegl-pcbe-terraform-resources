package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1StorageLocationsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\storage-locations
type V1beta1StorageLocationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1StorageLocationsRequestBuilderGetQueryParameters returns a list of storage locations. You can filter this list to determine which locations are enabled for a specified service.
type V1beta1StorageLocationsRequestBuilderGetQueryParameters struct {
    // A filter expression to apply. This endpoint only accepts the `in` operator on the `capabilities` property. Grouping expressions to change the evaluation precedence is not supported.
    Filter *string `uriparametername:"filter"`
}
// V1beta1StorageLocationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1StorageLocationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1StorageLocationsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.storageLocations.item collection
// returns a *V1beta1StorageLocationsStorageLocationsItemRequestBuilder when successful
func (m *V1beta1StorageLocationsRequestBuilder) ById(id string)(*V1beta1StorageLocationsStorageLocationsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1StorageLocationsStorageLocationsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1StorageLocationsRequestBuilderInternal instantiates a new V1beta1StorageLocationsRequestBuilder and sets the default values.
func NewV1beta1StorageLocationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1StorageLocationsRequestBuilder) {
    m := &V1beta1StorageLocationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/storage-locations{?filter*}", pathParameters),
    }
    return m
}
// NewV1beta1StorageLocationsRequestBuilder instantiates a new V1beta1StorageLocationsRequestBuilder and sets the default values.
func NewV1beta1StorageLocationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1StorageLocationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1StorageLocationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of storage locations. You can filter this list to determine which locations are enabled for a specified service.
// Deprecated: This method is obsolete. Use GetAsStorageLocationsGetResponse instead.
// returns a V1beta1StorageLocationsResponseable when successful
// returns a V1beta1StorageLocations400Error error when the service returns a 400 status code
// returns a V1beta1StorageLocations401Error error when the service returns a 401 status code
// returns a V1beta1StorageLocations403Error error when the service returns a 403 status code
// returns a V1beta1StorageLocations500Error error when the service returns a 500 status code
// returns a V1beta1StorageLocations503Error error when the service returns a 503 status code
func (m *V1beta1StorageLocationsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1StorageLocationsRequestBuilderGetRequestConfiguration)(V1beta1StorageLocationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1StorageLocations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1StorageLocations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1StorageLocations403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1StorageLocations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1StorageLocations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1StorageLocationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1StorageLocationsResponseable), nil
}
// GetAsStorageLocationsGetResponse returns a list of storage locations. You can filter this list to determine which locations are enabled for a specified service.
// returns a V1beta1StorageLocationsGetResponseable when successful
// returns a V1beta1StorageLocations400Error error when the service returns a 400 status code
// returns a V1beta1StorageLocations401Error error when the service returns a 401 status code
// returns a V1beta1StorageLocations403Error error when the service returns a 403 status code
// returns a V1beta1StorageLocations500Error error when the service returns a 500 status code
// returns a V1beta1StorageLocations503Error error when the service returns a 503 status code
func (m *V1beta1StorageLocationsRequestBuilder) GetAsStorageLocationsGetResponse(ctx context.Context, requestConfiguration *V1beta1StorageLocationsRequestBuilderGetRequestConfiguration)(V1beta1StorageLocationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1StorageLocations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1StorageLocations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1StorageLocations403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1StorageLocations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1StorageLocations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1StorageLocationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1StorageLocationsGetResponseable), nil
}
// ToGetRequestInformation returns a list of storage locations. You can filter this list to determine which locations are enabled for a specified service.
// returns a *RequestInformation when successful
func (m *V1beta1StorageLocationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1StorageLocationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1StorageLocationsRequestBuilder when successful
func (m *V1beta1StorageLocationsRequestBuilder) WithUrl(rawUrl string)(*V1beta1StorageLocationsRequestBuilder) {
    return NewV1beta1StorageLocationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
