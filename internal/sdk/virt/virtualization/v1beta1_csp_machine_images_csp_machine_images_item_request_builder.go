package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-images\{id}
type V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineImagesCspMachineImagesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineImagesCspMachineImagesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilderInternal instantiates a new V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder and sets the default values.
func NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) {
    m := &V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-images/{id}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilder instantiates a new V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder and sets the default values.
func NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns details of a specified cloud service provider (CSP) machine image.
// Deprecated: This method is obsolete. Use GetAsCspMachineImagesGetResponse instead.
// returns a V1beta1CspMachineImagesItemCspMachineImagesResponseable when successful
// returns a V1beta1CspMachineImagesItemCspMachineImages400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineImagesItemCspMachineImages401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineImagesItemCspMachineImages403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineImagesItemCspMachineImages500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilderGetRequestConfiguration)(V1beta1CspMachineImagesItemCspMachineImagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineImagesItemCspMachineImages400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineImagesItemCspMachineImages401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineImagesItemCspMachineImages403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineImagesItemCspMachineImages500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineImagesItemCspMachineImagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineImagesItemCspMachineImagesResponseable), nil
}
// GetAsCspMachineImagesGetResponse returns details of a specified cloud service provider (CSP) machine image.
// returns a V1beta1CspMachineImagesItemCspMachineImagesGetResponseable when successful
// returns a V1beta1CspMachineImagesItemCspMachineImages400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineImagesItemCspMachineImages401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineImagesItemCspMachineImages403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineImagesItemCspMachineImages500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) GetAsCspMachineImagesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilderGetRequestConfiguration)(V1beta1CspMachineImagesItemCspMachineImagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineImagesItemCspMachineImages400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineImagesItemCspMachineImages401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineImagesItemCspMachineImages403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineImagesItemCspMachineImages500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineImagesItemCspMachineImagesGetResponseable), nil
}
// ToGetRequestInformation returns details of a specified cloud service provider (CSP) machine image.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder when successful
func (m *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) {
    return NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
