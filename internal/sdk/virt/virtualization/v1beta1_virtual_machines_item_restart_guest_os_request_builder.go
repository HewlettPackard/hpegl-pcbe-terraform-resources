package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\restart-guest-os
type V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemRestartGuestOsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemRestartGuestOsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) {
    m := &V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/restart-guest-os", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilder instantiates a new V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restart guest OS of a virtual machine
// Deprecated: This method is obsolete. Use PostAsRestartGuestOsPostResponse instead.
// returns a V1beta1VirtualMachinesItemRestartGuestOsResponseable when successful
// returns a V1beta1VirtualMachinesItemRestartGuestOs400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemRestartGuestOsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemRestartGuestOs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemRestartGuestOs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemRestartGuestOs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemRestartGuestOs404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemRestartGuestOs409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemRestartGuestOs500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemRestartGuestOs503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemRestartGuestOsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemRestartGuestOsResponseable), nil
}
// PostAsRestartGuestOsPostResponse restart guest OS of a virtual machine
// returns a V1beta1VirtualMachinesItemRestartGuestOsPostResponseable when successful
// returns a V1beta1VirtualMachinesItemRestartGuestOs400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemRestartGuestOs503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) PostAsRestartGuestOsPostResponse(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemRestartGuestOsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemRestartGuestOs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemRestartGuestOs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemRestartGuestOs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemRestartGuestOs404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemRestartGuestOs409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemRestartGuestOs500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemRestartGuestOs503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemRestartGuestOsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemRestartGuestOsPostResponseable), nil
}
// ToPostRequestInformation restart guest OS of a virtual machine
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemRestartGuestOsRequestBuilder) {
    return NewV1beta1VirtualMachinesItemRestartGuestOsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
