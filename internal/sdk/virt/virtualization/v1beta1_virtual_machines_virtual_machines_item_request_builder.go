package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}
type V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilderInternal instantiates a new V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) {
    m := &V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilder instantiates a new V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomizeGuestOs the customizeGuestOs property
// returns a *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) CustomizeGuestOs()(*V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a virtual machine
// Deprecated: This method is obsolete. Use DeleteAsVirtualMachinesDeleteResponse instead.
// returns a V1beta1VirtualMachinesItemVirtualMachinesResponseable when successful
// returns a V1beta1VirtualMachinesItemVirtualMachines400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderDeleteRequestConfiguration)(V1beta1VirtualMachinesItemVirtualMachinesResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemVirtualMachines400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemVirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemVirtualMachines403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemVirtualMachines404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemVirtualMachines409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemVirtualMachines500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemVirtualMachines503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemVirtualMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemVirtualMachinesResponseable), nil
}
// DeleteAsVirtualMachinesDeleteResponse delete a virtual machine
// returns a V1beta1VirtualMachinesItemVirtualMachinesDeleteResponseable when successful
// returns a V1beta1VirtualMachinesItemVirtualMachines400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) DeleteAsVirtualMachinesDeleteResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderDeleteRequestConfiguration)(V1beta1VirtualMachinesItemVirtualMachinesDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemVirtualMachines400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemVirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemVirtualMachines403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemVirtualMachines404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemVirtualMachines409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemVirtualMachines500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemVirtualMachines503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemVirtualMachinesDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemVirtualMachinesDeleteResponseable), nil
}
// Get details of a virtual machine
// Deprecated: This method is obsolete. Use GetAsVirtualMachinesGetResponse instead.
// returns a V1beta1VirtualMachinesItemVirtualMachinesResponseable when successful
// returns a V1beta1VirtualMachinesItemVirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderGetRequestConfiguration)(V1beta1VirtualMachinesItemVirtualMachinesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1VirtualMachinesItemVirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemVirtualMachines403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemVirtualMachines404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemVirtualMachines500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemVirtualMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemVirtualMachinesResponseable), nil
}
// GetAsVirtualMachinesGetResponse details of a virtual machine
// returns a V1beta1VirtualMachinesItemVirtualMachinesGetResponseable when successful
// returns a V1beta1VirtualMachinesItemVirtualMachines401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemVirtualMachines500Error error when the service returns a 500 status code
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) GetAsVirtualMachinesGetResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderGetRequestConfiguration)(V1beta1VirtualMachinesItemVirtualMachinesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1VirtualMachinesItemVirtualMachines401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemVirtualMachines403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemVirtualMachines404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemVirtualMachines500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemVirtualMachinesGetResponseable), nil
}
// PowerOff the powerOff property
// returns a *V1beta1VirtualMachinesItemPowerOffRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) PowerOff()(*V1beta1VirtualMachinesItemPowerOffRequestBuilder) {
    return NewV1beta1VirtualMachinesItemPowerOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PowerOn the powerOn property
// returns a *V1beta1VirtualMachinesItemPowerOnRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) PowerOn()(*V1beta1VirtualMachinesItemPowerOnRequestBuilder) {
    return NewV1beta1VirtualMachinesItemPowerOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Refresh the refresh property
// returns a *V1beta1VirtualMachinesItemRefreshRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) Refresh()(*V1beta1VirtualMachinesItemRefreshRequestBuilder) {
    return NewV1beta1VirtualMachinesItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reset the reset property
// returns a *V1beta1VirtualMachinesItemResetRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) Reset()(*V1beta1VirtualMachinesItemResetRequestBuilder) {
    return NewV1beta1VirtualMachinesItemResetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestartGuestOs the restartGuestOs property
// returns a *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) RestartGuestOs()(*V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) {
    return NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutdownGuestOs the shutdownGuestOs property
// returns a *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) ShutdownGuestOs()(*V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) {
    return NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation details of a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// UpdateHardware the updateHardware property
// returns a *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) UpdateHardware()(*V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) {
    return NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder when successful
func (m *V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesVirtualMachinesItemRequestBuilder) {
    return NewV1beta1VirtualMachinesVirtualMachinesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
