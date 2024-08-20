package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instances\{id}
type V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderInternal instantiates a new V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) {
    m := &V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instances/{id}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder instantiates a new V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the cloud service provider (CSP) machine instance.
// Deprecated: This method is obsolete. Use DeleteAsCspMachineInstancesDeleteResponse instead.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesResponseable when successful
// returns a V1beta1CspMachineInstancesItemCspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderDeleteRequestConfiguration)(V1beta1CspMachineInstancesItemCspMachineInstancesResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemCspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemCspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemCspMachineInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemCspMachineInstances404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemCspMachineInstances409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemCspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemCspMachineInstancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemCspMachineInstancesResponseable), nil
}
// DeleteAsCspMachineInstancesDeleteResponse deletes the cloud service provider (CSP) machine instance.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseable when successful
// returns a V1beta1CspMachineInstancesItemCspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) DeleteAsCspMachineInstancesDeleteResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderDeleteRequestConfiguration)(V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemCspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemCspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemCspMachineInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemCspMachineInstances404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemCspMachineInstances409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemCspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemCspMachineInstancesDeleteResponseable), nil
}
// Get returns details of a specified cloud service provider (CSP) machine instance.
// Deprecated: This method is obsolete. Use GetAsCspMachineInstancesGetResponse instead.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesResponseable when successful
// returns a V1beta1CspMachineInstancesItemCspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstancesItemCspMachineInstancesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemCspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemCspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemCspMachineInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemCspMachineInstances404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemCspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemCspMachineInstancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemCspMachineInstancesResponseable), nil
}
// GetAsCspMachineInstancesGetResponse returns details of a specified cloud service provider (CSP) machine instance.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesGetResponseable when successful
// returns a V1beta1CspMachineInstancesItemCspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemCspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) GetAsCspMachineInstancesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemCspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemCspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemCspMachineInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemCspMachineInstances404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemCspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponseable), nil
}
// PowerOff the powerOff property
// returns a *V1beta1CspMachineInstancesItemPowerOffRequestBuilder when successful
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) PowerOff()(*V1beta1CspMachineInstancesItemPowerOffRequestBuilder) {
    return NewV1beta1CspMachineInstancesItemPowerOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PowerOn the powerOn property
// returns a *V1beta1CspMachineInstancesItemPowerOnRequestBuilder when successful
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) PowerOn()(*V1beta1CspMachineInstancesItemPowerOnRequestBuilder) {
    return NewV1beta1CspMachineInstancesItemPowerOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Refresh the refresh property
// returns a *V1beta1CspMachineInstancesItemRefreshRequestBuilder when successful
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) Refresh()(*V1beta1CspMachineInstancesItemRefreshRequestBuilder) {
    return NewV1beta1CspMachineInstancesItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes the cloud service provider (CSP) machine instance.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns details of a specified cloud service provider (CSP) machine instance.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder when successful
func (m *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) {
    return NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
