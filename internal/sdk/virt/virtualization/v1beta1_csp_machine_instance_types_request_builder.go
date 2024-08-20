package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstanceTypesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instance-types
type V1beta1CspMachineInstanceTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstanceTypesRequestBuilderGetQueryParameters returns a list of cloud service provider (CSP) machine instance types based on the queryparameters for paging, filtering, and sorting.
type V1beta1CspMachineInstanceTypesRequestBuilderGetQueryParameters struct {
    // An expression to filter the results.Filtering is supported with following attributes: * cspInfo.region * cspInfo.freeTierEligible * cspInfo.processorInfo.supportedArchitectures * cspInfo.supportedVirtualizationTypes * cspInfo.cspInfo.supportedRootDeviceTypes * cspInfo.hypervisor * cspInfo.ebsInfo.encryptionSupport * cspInfo.networkInfo.enaSupport * cspType
    Filter *string `uriparametername:"filter"`
    // The maximum number of items to include in the response.Use offset in conjunction with limit for pagination, for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to omit from the beginning of the result set.Use offset in conjunction with limit for pagination, for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc").
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspMachineInstanceTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstanceTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspMachineInstanceTypesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspMachineInstanceTypes.item collection
// returns a *V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder when successful
func (m *V1beta1CspMachineInstanceTypesRequestBuilder) ById(id string)(*V1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1CspMachineInstanceTypesCspMachineInstanceTypesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspMachineInstanceTypesRequestBuilderInternal instantiates a new V1beta1CspMachineInstanceTypesRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstanceTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstanceTypesRequestBuilder) {
    m := &V1beta1CspMachineInstanceTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instance-types{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstanceTypesRequestBuilder instantiates a new V1beta1CspMachineInstanceTypesRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstanceTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstanceTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstanceTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of cloud service provider (CSP) machine instance types based on the queryparameters for paging, filtering, and sorting.
// Deprecated: This method is obsolete. Use GetAsCspMachineInstanceTypesGetResponse instead.
// returns a V1beta1CspMachineInstanceTypesResponseable when successful
// returns a V1beta1CspMachineInstanceTypes400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstanceTypes401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstanceTypes403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstanceTypes500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstanceTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspMachineInstanceTypesRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstanceTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstanceTypes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstanceTypes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstanceTypes403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstanceTypes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstanceTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstanceTypesResponseable), nil
}
// GetAsCspMachineInstanceTypesGetResponse returns a list of cloud service provider (CSP) machine instance types based on the queryparameters for paging, filtering, and sorting.
// returns a V1beta1CspMachineInstanceTypesGetResponseable when successful
// returns a V1beta1CspMachineInstanceTypes400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstanceTypes401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstanceTypes403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstanceTypes500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstanceTypesRequestBuilder) GetAsCspMachineInstanceTypesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstanceTypesRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstanceTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstanceTypes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstanceTypes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstanceTypes403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstanceTypes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstanceTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstanceTypesGetResponseable), nil
}
// ToGetRequestInformation returns a list of cloud service provider (CSP) machine instance types based on the queryparameters for paging, filtering, and sorting.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstanceTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstanceTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstanceTypesRequestBuilder when successful
func (m *V1beta1CspMachineInstanceTypesRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstanceTypesRequestBuilder) {
    return NewV1beta1CspMachineInstanceTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
