package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineImagesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-images
type V1beta1CspMachineImagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineImagesRequestBuilderGetQueryParameters returns a list of cloud service provider (CSP) machine images based on the queryparameters for paging, filtering, and sorting.
type V1beta1CspMachineImagesRequestBuilderGetQueryParameters struct {
    // An expression to filter the results.Filtering is supported with following attributes: * name * cspInfo.id  * cspInfo.region * cspInfo.architecture * cspInfo.state * cspInfo.hypervisor * cspInfo.rootDeviceType * cspInfo.rootDeviceName * cspInfo.virtualizationType * cspType * cspInfo.location * cspInfo.ownerId
    Filter *string `uriparametername:"filter"`
    // The maximum number of items to include in the response.Use offset in conjunction with limit for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to omit from the beginning of the result set.Use offset in conjunction with limit for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc").
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspMachineImagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineImagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspMachineImagesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspMachineImages.item collection
// returns a *V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder when successful
func (m *V1beta1CspMachineImagesRequestBuilder) ById(id string)(*V1beta1CspMachineImagesCspMachineImagesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1CspMachineImagesCspMachineImagesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspMachineImagesRequestBuilderInternal instantiates a new V1beta1CspMachineImagesRequestBuilder and sets the default values.
func NewV1beta1CspMachineImagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineImagesRequestBuilder) {
    m := &V1beta1CspMachineImagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-images{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineImagesRequestBuilder instantiates a new V1beta1CspMachineImagesRequestBuilder and sets the default values.
func NewV1beta1CspMachineImagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineImagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineImagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of cloud service provider (CSP) machine images based on the queryparameters for paging, filtering, and sorting.
// Deprecated: This method is obsolete. Use GetAsCspMachineImagesGetResponse instead.
// returns a V1beta1CspMachineImagesResponseable when successful
// returns a V1beta1CspMachineImages400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineImages401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineImages403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineImages500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineImagesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspMachineImagesRequestBuilderGetRequestConfiguration)(V1beta1CspMachineImagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineImages400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineImages401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineImages403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineImages500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineImagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineImagesResponseable), nil
}
// GetAsCspMachineImagesGetResponse returns a list of cloud service provider (CSP) machine images based on the queryparameters for paging, filtering, and sorting.
// returns a V1beta1CspMachineImagesGetResponseable when successful
// returns a V1beta1CspMachineImages400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineImages401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineImages403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineImages500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineImagesRequestBuilder) GetAsCspMachineImagesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineImagesRequestBuilderGetRequestConfiguration)(V1beta1CspMachineImagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineImages400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineImages401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineImages403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineImages500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineImagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineImagesGetResponseable), nil
}
// ToGetRequestInformation returns a list of cloud service provider (CSP) machine images based on the queryparameters for paging, filtering, and sorting.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineImagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineImagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspMachineImagesRequestBuilder when successful
func (m *V1beta1CspMachineImagesRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineImagesRequestBuilder) {
    return NewV1beta1CspMachineImagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
