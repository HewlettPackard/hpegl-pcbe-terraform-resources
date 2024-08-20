package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\virtual-machines\{id}\customize-guest-os
type V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderInternal instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/virtual-machines/{id}/customize-guest-os", pathParameters),
    }
    return m
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post customize guest operating system attributes.
// Deprecated: This method is obsolete. Use PostAsCustomizeGuestOsPostResponse instead.
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsResponseable when successful
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) Post(ctx context.Context, body V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemCustomizeGuestOsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemCustomizeGuestOs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemCustomizeGuestOs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemCustomizeGuestOs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemCustomizeGuestOs404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemCustomizeGuestOs409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemCustomizeGuestOs500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemCustomizeGuestOs503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemCustomizeGuestOsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemCustomizeGuestOsResponseable), nil
}
// PostAsCustomizeGuestOsPostResponse customize guest operating system attributes.
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostResponseable when successful
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs400Error error when the service returns a 400 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs401Error error when the service returns a 401 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs403Error error when the service returns a 403 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs404Error error when the service returns a 404 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs409Error error when the service returns a 409 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs500Error error when the service returns a 500 status code
// returns a V1beta1VirtualMachinesItemCustomizeGuestOs503Error error when the service returns a 503 status code
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) PostAsCustomizeGuestOsPostResponse(ctx context.Context, body V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderPostRequestConfiguration)(V1beta1VirtualMachinesItemCustomizeGuestOsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1VirtualMachinesItemCustomizeGuestOs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1VirtualMachinesItemCustomizeGuestOs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1VirtualMachinesItemCustomizeGuestOs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1VirtualMachinesItemCustomizeGuestOs404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1VirtualMachinesItemCustomizeGuestOs409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1VirtualMachinesItemCustomizeGuestOs500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1VirtualMachinesItemCustomizeGuestOs503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1VirtualMachinesItemCustomizeGuestOsPostResponseable), nil
}
// ToPostRequestInformation customize guest operating system attributes.
// returns a *RequestInformation when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBodyable, requestConfiguration *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) WithUrl(rawUrl string)(*V1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
