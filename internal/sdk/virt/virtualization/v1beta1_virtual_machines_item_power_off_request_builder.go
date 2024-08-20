package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemPowerOffRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\power-off
type V1beta1VirtualMachinesItemPowerOffRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemPowerOffRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemPowerOffRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemPowerOffRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemPowerOffRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemPowerOffRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemPowerOffRequestBuilder) {
    m := &V1beta1VirtualMachinesItemPowerOffRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/power-off", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemPowerOffRequestBuilder instantiates a new V1beta1VirtualMachinesItemPowerOffRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemPowerOffRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemPowerOffRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemPowerOffRequestBuilderInternal(urlParams, requestAdapter)
}
// Post power off a virtual machine
// Deprecated: This method is obsolete. Use PostAsPowerOffPostResponse instead.
// returns a V1beta1VirtualMachinesItemPowerOffResponseable when successful
// returns a V1beta1VirtualMachinesItemPowerOff400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemPowerOff401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemPowerOff403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemPowerOff404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemPowerOff409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemPowerOff500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemPowerOff503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemPowerOffRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemPowerOffRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemPowerOffResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemPowerOff400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemPowerOff401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemPowerOff403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemPowerOff404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemPowerOff409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemPowerOff500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemPowerOff503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemPowerOffResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemPowerOffResponseable), nil
}
// PostAsPowerOffPostResponse power off a virtual machine
// returns a V1beta1VirtualMachinesItemPowerOffPostResponseable when successful
// returns a V1beta1VirtualMachinesItemPowerOff400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemPowerOff401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemPowerOff403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemPowerOff404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemPowerOff409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemPowerOff500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemPowerOff503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemPowerOffRequestBuilder) PostAsPowerOffPostResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemPowerOffRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemPowerOffPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemPowerOff400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemPowerOff401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemPowerOff403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemPowerOff404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemPowerOff409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemPowerOff500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemPowerOff503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemPowerOffPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemPowerOffPostResponseable), nil
}
// ToPostRequestInformation power off a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemPowerOffRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemPowerOffRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesItemPowerOffRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemPowerOffRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemPowerOffRequestBuilder) {
    return NewV1beta1VirtualMachinesItemPowerOffRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
