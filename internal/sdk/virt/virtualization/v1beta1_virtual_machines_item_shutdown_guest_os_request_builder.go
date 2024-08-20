package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\shutdown-guest-os
type V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) {
    m := &V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/shutdown-guest-os", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder instantiates a new V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post shutdown guest OS of a virtual machine
// Deprecated: This method is obsolete. Use PostAsShutdownGuestOsPostResponse instead.
// returns a V1beta1VirtualMachinesItemShutdownGuestOsResponseable when successful
// returns a V1beta1VirtualMachinesItemShutdownGuestOs400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemShutdownGuestOsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemShutdownGuestOs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemShutdownGuestOs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemShutdownGuestOs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemShutdownGuestOs404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemShutdownGuestOs409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemShutdownGuestOs500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemShutdownGuestOs503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemShutdownGuestOsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemShutdownGuestOsResponseable), nil
}
// PostAsShutdownGuestOsPostResponse shutdown guest OS of a virtual machine
// returns a V1beta1VirtualMachinesItemShutdownGuestOsPostResponseable when successful
// returns a V1beta1VirtualMachinesItemShutdownGuestOs400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemShutdownGuestOs503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) PostAsShutdownGuestOsPostResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemShutdownGuestOsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemShutdownGuestOs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemShutdownGuestOs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemShutdownGuestOs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemShutdownGuestOs404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemShutdownGuestOs409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemShutdownGuestOs500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemShutdownGuestOs503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemShutdownGuestOsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemShutdownGuestOsPostResponseable), nil
}
// ToPostRequestInformation shutdown guest OS of a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder) {
    return NewV1beta1VirtualMachinesItemShutdownGuestOsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
