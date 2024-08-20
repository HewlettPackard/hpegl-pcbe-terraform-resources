package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1AsyncOperationsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\async-operations
type V1beta1AsyncOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1AsyncOperationsRequestBuilderGetQueryParameters returns a list of async-operations that are visible to the user. The response canbe paged by using the limit and offset query parameters and filtered andsorted by using the filter and sort query parameters.
type V1beta1AsyncOperationsRequestBuilderGetQueryParameters struct {
    // The expression to filter responses.
    Filter *string `uriparametername:"filter"`
    // The limit query parameter should be used in conjunction with offsetfor paging, e.g.: offset=30&&limit=10. The limit is the maximumnumber of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // The offset query parameter should be used in conjunction with limitfor paging, e.g.: offset=30&&limit=10. The offset is the number ofitems from the beginning of the result set to the first itemincluded in the response.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc"). If no direction indicator is specified the default order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1AsyncOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1AsyncOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1AsyncOperationsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async.dataServices.v1beta1.asyncOperations.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder when successful
func (m *V1beta1AsyncOperationsRequestBuilder) ById(id string)(*V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async.dataServices.v1beta1.asyncOperations.item collection
// returns a *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder when successful
func (m *V1beta1AsyncOperationsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1AsyncOperationsRequestBuilderInternal instantiates a new V1beta1AsyncOperationsRequestBuilder and sets the default values.
func NewV1beta1AsyncOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AsyncOperationsRequestBuilder) {
    m := &V1beta1AsyncOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/async-operations{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1AsyncOperationsRequestBuilder instantiates a new V1beta1AsyncOperationsRequestBuilder and sets the default values.
func NewV1beta1AsyncOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AsyncOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1AsyncOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of async-operations that are visible to the user. The response canbe paged by using the limit and offset query parameters and filtered andsorted by using the filter and sort query parameters.
// Deprecated: This method is obsolete. Use GetAsAsyncOperationsGetResponse instead.
// returns a V1beta1AsyncOperationsResponseable when successful
// returns a V1beta1AsyncOperations400Error error when the service returns a 400 status code
// returns a V1beta1AsyncOperations401Error error when the service returns a 401 status code
// returns a V1beta1AsyncOperations403Error error when the service returns a 403 status code
// returns a V1beta1AsyncOperations500Error error when the service returns a 500 status code
// returns a V1beta1AsyncOperations503Error error when the service returns a 503 status code
func (m *V1beta1AsyncOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1AsyncOperationsRequestBuilderGetRequestConfiguration)(V1beta1AsyncOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1AsyncOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1AsyncOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1AsyncOperations403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1AsyncOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1AsyncOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AsyncOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AsyncOperationsResponseable), nil
}
// GetAsAsyncOperationsGetResponse returns a list of async-operations that are visible to the user. The response canbe paged by using the limit and offset query parameters and filtered andsorted by using the filter and sort query parameters.
// returns a V1beta1AsyncOperationsGetResponseable when successful
// returns a V1beta1AsyncOperations400Error error when the service returns a 400 status code
// returns a V1beta1AsyncOperations401Error error when the service returns a 401 status code
// returns a V1beta1AsyncOperations403Error error when the service returns a 403 status code
// returns a V1beta1AsyncOperations500Error error when the service returns a 500 status code
// returns a V1beta1AsyncOperations503Error error when the service returns a 503 status code
func (m *V1beta1AsyncOperationsRequestBuilder) GetAsAsyncOperationsGetResponse(ctx context.Context, requestConfiguration *V1beta1AsyncOperationsRequestBuilderGetRequestConfiguration)(V1beta1AsyncOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1AsyncOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1AsyncOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1AsyncOperations403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1AsyncOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1AsyncOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AsyncOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AsyncOperationsGetResponseable), nil
}
// ToGetRequestInformation returns a list of async-operations that are visible to the user. The response canbe paged by using the limit and offset query parameters and filtered andsorted by using the filter and sort query parameters.
// returns a *RequestInformation when successful
func (m *V1beta1AsyncOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1AsyncOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1AsyncOperationsRequestBuilder when successful
func (m *V1beta1AsyncOperationsRequestBuilder) WithUrl(rawUrl string)(*V1beta1AsyncOperationsRequestBuilder) {
    return NewV1beta1AsyncOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
