package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\dual-auth-operations\{id}
type V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderInternal instantiates a new V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder and sets the default values.
func NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) {
    m := &V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/dual-auth-operations/{id}", pathParameters),
    }
    return m
}
// NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder instantiates a new V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder and sets the default values.
func NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the Dual Authorization operation for the given Id. The operation must belong to the current account and be a resource type the user has read permission for.  The user must also have permission to read pending operations.
// Deprecated: This method is obsolete. Use GetAsDualAuthOperationsGetResponse instead.
// returns a V1beta1DualAuthOperationsItemDualAuthOperationsResponseable when successful
// returns a V1beta1DualAuthOperationsItemDualAuthOperations400Error error when the service returns a 400 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations401Error error when the service returns a 401 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations403Error error when the service returns a 403 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations404Error error when the service returns a 404 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations422Error error when the service returns a 422 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations500Error error when the service returns a 500 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations503Error error when the service returns a 503 status code
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderGetRequestConfiguration)(V1beta1DualAuthOperationsItemDualAuthOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DualAuthOperationsItemDualAuthOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DualAuthOperationsItemDualAuthOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DualAuthOperationsItemDualAuthOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DualAuthOperationsItemDualAuthOperations404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1DualAuthOperationsItemDualAuthOperations422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DualAuthOperationsItemDualAuthOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DualAuthOperationsItemDualAuthOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DualAuthOperationsItemDualAuthOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DualAuthOperationsItemDualAuthOperationsResponseable), nil
}
// GetAsDualAuthOperationsGetResponse returns the Dual Authorization operation for the given Id. The operation must belong to the current account and be a resource type the user has read permission for.  The user must also have permission to read pending operations.
// returns a V1beta1DualAuthOperationsItemDualAuthOperationsGetResponseable when successful
// returns a V1beta1DualAuthOperationsItemDualAuthOperations400Error error when the service returns a 400 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations401Error error when the service returns a 401 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations403Error error when the service returns a 403 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations404Error error when the service returns a 404 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations422Error error when the service returns a 422 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations500Error error when the service returns a 500 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations503Error error when the service returns a 503 status code
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) GetAsDualAuthOperationsGetResponse(ctx context.Context, requestConfiguration *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderGetRequestConfiguration)(V1beta1DualAuthOperationsItemDualAuthOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DualAuthOperationsItemDualAuthOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DualAuthOperationsItemDualAuthOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DualAuthOperationsItemDualAuthOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DualAuthOperationsItemDualAuthOperations404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1DualAuthOperationsItemDualAuthOperations422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DualAuthOperationsItemDualAuthOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DualAuthOperationsItemDualAuthOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DualAuthOperationsItemDualAuthOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DualAuthOperationsItemDualAuthOperationsGetResponseable), nil
}
// Patch approve/Deny the pending operation by changing its state in DB.
// Deprecated: This method is obsolete. Use PatchAsDualAuthOperationsPatchResponse instead.
// returns a V1beta1DualAuthOperationsItemDualAuthOperationsResponseable when successful
// returns a V1beta1DualAuthOperationsItemDualAuthOperations400Error error when the service returns a 400 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations401Error error when the service returns a 401 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations403Error error when the service returns a 403 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations404Error error when the service returns a 404 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations412Error error when the service returns a 412 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations422Error error when the service returns a 422 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations500Error error when the service returns a 500 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations503Error error when the service returns a 503 status code
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) Patch(ctx context.Context, body V1beta1DualAuthOperationsItemDualAuthOperationsPatchRequestBodyable, requestConfiguration *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderPatchRequestConfiguration)(V1beta1DualAuthOperationsItemDualAuthOperationsResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DualAuthOperationsItemDualAuthOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DualAuthOperationsItemDualAuthOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DualAuthOperationsItemDualAuthOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DualAuthOperationsItemDualAuthOperations404ErrorFromDiscriminatorValue,
        "412": CreateV1beta1DualAuthOperationsItemDualAuthOperations412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1DualAuthOperationsItemDualAuthOperations422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DualAuthOperationsItemDualAuthOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DualAuthOperationsItemDualAuthOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DualAuthOperationsItemDualAuthOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DualAuthOperationsItemDualAuthOperationsResponseable), nil
}
// PatchAsDualAuthOperationsPatchResponse approve/Deny the pending operation by changing its state in DB.
// returns a V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponseable when successful
// returns a V1beta1DualAuthOperationsItemDualAuthOperations400Error error when the service returns a 400 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations401Error error when the service returns a 401 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations403Error error when the service returns a 403 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations404Error error when the service returns a 404 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations412Error error when the service returns a 412 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations422Error error when the service returns a 422 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations500Error error when the service returns a 500 status code
// returns a V1beta1DualAuthOperationsItemDualAuthOperations503Error error when the service returns a 503 status code
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) PatchAsDualAuthOperationsPatchResponse(ctx context.Context, body V1beta1DualAuthOperationsItemDualAuthOperationsPatchRequestBodyable, requestConfiguration *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderPatchRequestConfiguration)(V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DualAuthOperationsItemDualAuthOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DualAuthOperationsItemDualAuthOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DualAuthOperationsItemDualAuthOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DualAuthOperationsItemDualAuthOperations404ErrorFromDiscriminatorValue,
        "412": CreateV1beta1DualAuthOperationsItemDualAuthOperations412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1DualAuthOperationsItemDualAuthOperations422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DualAuthOperationsItemDualAuthOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DualAuthOperationsItemDualAuthOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DualAuthOperationsItemDualAuthOperationsPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DualAuthOperationsItemDualAuthOperationsPatchResponseable), nil
}
// ToGetRequestInformation returns the Dual Authorization operation for the given Id. The operation must belong to the current account and be a resource type the user has read permission for.  The user must also have permission to read pending operations.
// returns a *RequestInformation when successful
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation approve/Deny the pending operation by changing its state in DB.
// returns a *RequestInformation when successful
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1DualAuthOperationsItemDualAuthOperationsPatchRequestBodyable, requestConfiguration *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/merge-patch+json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder when successful
func (m *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) {
    return NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
