package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\update-hardware
type V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemUpdateHardwareRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemUpdateHardwareRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) {
    m := &V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/update-hardware", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilder instantiates a new V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates CPU, memory, network adapters, and disks of a virtual machine. This operation can be performed when the virtual machine is powered off.
// Deprecated: This method is obsolete. Use PostAsUpdateHardwarePostResponse instead.
// returns a V1beta1VirtualMachinesItemUpdateHardwareResponseable when successful
// returns a V1beta1VirtualMachinesItemUpdateHardware400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) Post(ctx context.Context, body V1beta1VirtualMachinesItemUpdateHardwarePostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemUpdateHardwareResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemUpdateHardware400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemUpdateHardware401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemUpdateHardware403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemUpdateHardware404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemUpdateHardware409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemUpdateHardware500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemUpdateHardware503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemUpdateHardwareResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemUpdateHardwareResponseable), nil
}
// PostAsUpdateHardwarePostResponse updates CPU, memory, network adapters, and disks of a virtual machine. This operation can be performed when the virtual machine is powered off.
// returns a V1beta1VirtualMachinesItemUpdateHardwarePostResponseable when successful
// returns a V1beta1VirtualMachinesItemUpdateHardware400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemUpdateHardware503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) PostAsUpdateHardwarePostResponse(ctx context.Context, body V1beta1VirtualMachinesItemUpdateHardwarePostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemUpdateHardwarePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemUpdateHardware400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemUpdateHardware401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemUpdateHardware403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemUpdateHardware404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemUpdateHardware409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemUpdateHardware500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemUpdateHardware503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemUpdateHardwarePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemUpdateHardwarePostResponseable), nil
}
// ToPostRequestInformation updates CPU, memory, network adapters, and disks of a virtual machine. This operation can be performed when the virtual machine is powered off.
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1VirtualMachinesItemUpdateHardwarePostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemUpdateHardwareRequestBuilder) {
    return NewV1beta1VirtualMachinesItemUpdateHardwareRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
