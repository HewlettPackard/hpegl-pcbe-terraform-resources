package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemResetRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\reset
type V1beta1VirtualMachinesItemResetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemResetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemResetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemResetRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemResetRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemResetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemResetRequestBuilder) {
    m := &V1beta1VirtualMachinesItemResetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/reset", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemResetRequestBuilder instantiates a new V1beta1VirtualMachinesItemResetRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemResetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemResetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemResetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset a virtual machine
// Deprecated: This method is obsolete. Use PostAsResetPostResponse instead.
// returns a V1beta1VirtualMachinesItemResetResponseable when successful
// returns a V1beta1VirtualMachinesItemReset400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemReset401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemReset403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemReset404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemReset409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemReset500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemReset503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemResetRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemResetRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemResetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemReset400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemReset401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemReset403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemReset404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemReset409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemReset500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemReset503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemResetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemResetResponseable), nil
}
// PostAsResetPostResponse reset a virtual machine
// returns a V1beta1VirtualMachinesItemResetPostResponseable when successful
// returns a V1beta1VirtualMachinesItemReset400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemReset401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemReset403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemReset404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemReset409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemReset500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemReset503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemResetRequestBuilder) PostAsResetPostResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemResetRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemResetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemReset400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemReset401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemReset403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemReset404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemReset409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemReset500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemReset503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemResetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemResetPostResponseable), nil
}
// ToPostRequestInformation reset a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemResetRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemResetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesItemResetRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemResetRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemResetRequestBuilder) {
    return NewV1beta1VirtualMachinesItemResetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
