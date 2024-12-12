package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\async-operations\{id}
type V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetQueryParameters returns the async-operation with the given id.
type V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetQueryParameters struct {
    // A list of properties to include in the response.
    Select *string `uriparametername:"select"`
}
// V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetQueryParameters
}
// NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilderInternal instantiates a new V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder and sets the default values.
func NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) {
    m := &V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/async-operations/{id}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilder instantiates a new V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder and sets the default values.
func NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the async-operation with the given id.
// Deprecated: This method is obsolete. Use GetAsAsyncOperationsGetResponse instead.
// returns a V1beta1AsyncOperationsItemAsyncOperationsResponseable when successful
// returns a V1beta1AsyncOperationsItemAsyncOperations400Error error when the service returns a 400 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations401Error error when the service returns a 401 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations403Error error when the service returns a 403 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations404Error error when the service returns a 404 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations500Error error when the service returns a 500 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations503Error error when the service returns a 503 status code
func (m *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration)(V1beta1AsyncOperationsItemAsyncOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1AsyncOperationsItemAsyncOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1AsyncOperationsItemAsyncOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1AsyncOperationsItemAsyncOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1AsyncOperationsItemAsyncOperations404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1AsyncOperationsItemAsyncOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1AsyncOperationsItemAsyncOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AsyncOperationsItemAsyncOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AsyncOperationsItemAsyncOperationsResponseable), nil
}
// GetAsAsyncOperationsGetResponse returns the async-operation with the given id.
// returns a V1beta1AsyncOperationsItemAsyncOperationsGetResponseable when successful
// returns a V1beta1AsyncOperationsItemAsyncOperations400Error error when the service returns a 400 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations401Error error when the service returns a 401 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations403Error error when the service returns a 403 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations404Error error when the service returns a 404 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations500Error error when the service returns a 500 status code
// returns a V1beta1AsyncOperationsItemAsyncOperations503Error error when the service returns a 503 status code
func (m *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) GetAsAsyncOperationsGetResponse(ctx context.Context, requestConfiguration *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration)(V1beta1AsyncOperationsItemAsyncOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1AsyncOperationsItemAsyncOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1AsyncOperationsItemAsyncOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1AsyncOperationsItemAsyncOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1AsyncOperationsItemAsyncOperations404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1AsyncOperationsItemAsyncOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1AsyncOperationsItemAsyncOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AsyncOperationsItemAsyncOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AsyncOperationsItemAsyncOperationsGetResponseable), nil
}
// ToGetRequestInformation returns the async-operation with the given id.
// returns a *RequestInformation when successful
func (m *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder when successful
func (m *V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1AsyncOperationsAsyncOperationsItemRequestBuilder) {
    return NewV1beta1AsyncOperationsAsyncOperationsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
