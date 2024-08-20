package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstancesItemPowerOnRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instances\{id}\power-on
type V1beta1CspMachineInstancesItemPowerOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstancesItemPowerOnRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesItemPowerOnRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspMachineInstancesItemPowerOnRequestBuilderInternal instantiates a new V1beta1CspMachineInstancesItemPowerOnRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesItemPowerOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesItemPowerOnRequestBuilder) {
    m := &V1beta1CspMachineInstancesItemPowerOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instances/{id}/power-on", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstancesItemPowerOnRequestBuilder instantiates a new V1beta1CspMachineInstancesItemPowerOnRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesItemPowerOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesItemPowerOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstancesItemPowerOnRequestBuilderInternal(urlParams, requestAdapter)
}
// Post powers on cloud service provider (CSP) machine instance.
// Deprecated: This method is obsolete. Use PostAsPowerOnPostResponse instead.
// returns a V1beta1CspMachineInstancesItemPowerOnResponseable when successful
// returns a V1beta1CspMachineInstancesItemPowerOn400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemPowerOn401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemPowerOn403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemPowerOn404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemPowerOn409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemPowerOn500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesItemPowerOnRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemPowerOnRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesItemPowerOnResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemPowerOn400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemPowerOn401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemPowerOn403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemPowerOn404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemPowerOn409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemPowerOn500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemPowerOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemPowerOnResponseable), nil
}
// PostAsPowerOnPostResponse powers on cloud service provider (CSP) machine instance.
// returns a V1beta1CspMachineInstancesItemPowerOnPostResponseable when successful
// returns a V1beta1CspMachineInstancesItemPowerOn400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemPowerOn401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemPowerOn403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemPowerOn404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemPowerOn409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemPowerOn500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesItemPowerOnRequestBuilder) PostAsPowerOnPostResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemPowerOnRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesItemPowerOnPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemPowerOn400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemPowerOn401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemPowerOn403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemPowerOn404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemPowerOn409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemPowerOn500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemPowerOnPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemPowerOnPostResponseable), nil
}
// ToPostRequestInformation powers on cloud service provider (CSP) machine instance.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesItemPowerOnRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemPowerOnRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstancesItemPowerOnRequestBuilder when successful
func (m *V1beta1CspMachineInstancesItemPowerOnRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstancesItemPowerOnRequestBuilder) {
    return NewV1beta1CspMachineInstancesItemPowerOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
