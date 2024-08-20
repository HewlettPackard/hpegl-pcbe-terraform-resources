package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemPowerOnRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\power-on
type V1beta1VirtualMachinesItemPowerOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemPowerOnRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemPowerOnRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemPowerOnRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemPowerOnRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemPowerOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemPowerOnRequestBuilder) {
    m := &V1beta1VirtualMachinesItemPowerOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/power-on", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemPowerOnRequestBuilder instantiates a new V1beta1VirtualMachinesItemPowerOnRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemPowerOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemPowerOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemPowerOnRequestBuilderInternal(urlParams, requestAdapter)
}
// Post power on a virtual machine
// Deprecated: This method is obsolete. Use PostAsPowerOnPostResponse instead.
// returns a V1beta1VirtualMachinesItemPowerOnResponseable when successful
// returns a V1beta1VirtualMachinesItemPowerOn400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemPowerOn401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemPowerOn403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemPowerOn404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemPowerOn409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemPowerOn500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemPowerOn503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemPowerOnRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemPowerOnRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemPowerOnResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemPowerOn400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemPowerOn401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemPowerOn403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemPowerOn404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemPowerOn409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemPowerOn500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemPowerOn503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemPowerOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemPowerOnResponseable), nil
}
// PostAsPowerOnPostResponse power on a virtual machine
// returns a V1beta1VirtualMachinesItemPowerOnPostResponseable when successful
// returns a V1beta1VirtualMachinesItemPowerOn400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemPowerOn401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemPowerOn403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemPowerOn404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemPowerOn409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemPowerOn500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemPowerOn503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemPowerOnRequestBuilder) PostAsPowerOnPostResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemPowerOnRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemPowerOnPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemPowerOn400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemPowerOn401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemPowerOn403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemPowerOn404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemPowerOn409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemPowerOn500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemPowerOn503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemPowerOnPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemPowerOnPostResponseable), nil
}
// ToPostRequestInformation power on a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemPowerOnRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemPowerOnRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesItemPowerOnRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemPowerOnRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemPowerOnRequestBuilder) {
    return NewV1beta1VirtualMachinesItemPowerOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
