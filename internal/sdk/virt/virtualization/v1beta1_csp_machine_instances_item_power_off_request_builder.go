package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstancesItemPowerOffRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instances\{id}\power-off
type V1beta1CspMachineInstancesItemPowerOffRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstancesItemPowerOffRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesItemPowerOffRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspMachineInstancesItemPowerOffRequestBuilderInternal instantiates a new V1beta1CspMachineInstancesItemPowerOffRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesItemPowerOffRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesItemPowerOffRequestBuilder) {
    m := &V1beta1CspMachineInstancesItemPowerOffRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instances/{id}/power-off", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstancesItemPowerOffRequestBuilder instantiates a new V1beta1CspMachineInstancesItemPowerOffRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesItemPowerOffRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesItemPowerOffRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstancesItemPowerOffRequestBuilderInternal(urlParams, requestAdapter)
}
// Post powers off the cloud service provider (CSP) machine instance.
// Deprecated: This method is obsolete. Use PostAsPowerOffPostResponse instead.
// returns a V1beta1CspMachineInstancesItemPowerOffResponseable when successful
// returns a V1beta1CspMachineInstancesItemPowerOff400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemPowerOff401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemPowerOff403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemPowerOff404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemPowerOff409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemPowerOff500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesItemPowerOffRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemPowerOffRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesItemPowerOffResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemPowerOff400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemPowerOff401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemPowerOff403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemPowerOff404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemPowerOff409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemPowerOff500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemPowerOffResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemPowerOffResponseable), nil
}
// PostAsPowerOffPostResponse powers off the cloud service provider (CSP) machine instance.
// returns a V1beta1CspMachineInstancesItemPowerOffPostResponseable when successful
// returns a V1beta1CspMachineInstancesItemPowerOff400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesItemPowerOff401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemPowerOff403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemPowerOff404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemPowerOff409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemPowerOff500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesItemPowerOffRequestBuilder) PostAsPowerOffPostResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemPowerOffRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesItemPowerOffPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesItemPowerOff400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesItemPowerOff401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemPowerOff403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemPowerOff404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemPowerOff409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemPowerOff500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemPowerOffPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemPowerOffPostResponseable), nil
}
// ToPostRequestInformation powers off the cloud service provider (CSP) machine instance.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesItemPowerOffRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemPowerOffRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstancesItemPowerOffRequestBuilder when successful
func (m *V1beta1CspMachineInstancesItemPowerOffRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstancesItemPowerOffRequestBuilder) {
    return NewV1beta1CspMachineInstancesItemPowerOffRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
