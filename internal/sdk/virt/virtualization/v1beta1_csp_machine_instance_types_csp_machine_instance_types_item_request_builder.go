package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instance-types\{id}
type V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderInternal instantiates a new V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) {
    m := &V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instance-types/{id}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder instantiates a new V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns details of a specified cloud service provider (CSP) machine instance type.
// Deprecated: This method is obsolete. Use GetAsCspMachineInstanceTypesGetResponse instead.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesResponseable when successful
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesResponseable), nil
}
// GetAsCspMachineInstanceTypesGetResponse returns details of a specified cloud service provider (CSP) machine instance type.
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseable when successful
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) GetAsCspMachineInstanceTypesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseable), nil
}
// ToGetRequestInformation returns details of a specified cloud service provider (CSP) machine instance type.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder when successful
func (m *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) {
    return NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
