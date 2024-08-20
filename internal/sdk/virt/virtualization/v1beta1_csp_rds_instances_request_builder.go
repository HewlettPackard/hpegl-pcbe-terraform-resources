package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspRdsInstancesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-rds-instances
type V1beta1CspRdsInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspRdsInstancesRequestBuilderGetQueryParameters returns a list of cloud service provider (CSP) RDS machine instances according to the given queryparameters for paging, filtering.
type V1beta1CspRdsInstancesRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The comparisons supported are the following:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.Filters are supported on following attributes:* accountInfo/id* name* protectionStatus
    Filter *string `uriparametername:"filter"`
    // The maximum number of items to include in the response.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to omit from the beginning of the result set.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Offset *int32 `uriparametername:"offset"`
}
// V1beta1CspRdsInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspRdsInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspRdsInstancesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspRdsInstances.item collection
// returns a *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder when successful
func (m *V1beta1CspRdsInstancesRequestBuilder) ById(id string)(*V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspRdsInstancesRequestBuilderInternal instantiates a new V1beta1CspRdsInstancesRequestBuilder and sets the default values.
func NewV1beta1CspRdsInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspRdsInstancesRequestBuilder) {
    m := &V1beta1CspRdsInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-rds-instances{?filter*,limit*,offset*}", pathParameters),
    }
    return m
}
// NewV1beta1CspRdsInstancesRequestBuilder instantiates a new V1beta1CspRdsInstancesRequestBuilder and sets the default values.
func NewV1beta1CspRdsInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspRdsInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspRdsInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of cloud service provider (CSP) RDS machine instances according to the given queryparameters for paging, filtering.
// Deprecated: This method is obsolete. Use GetAsCspRdsInstancesGetResponse instead.
// returns a V1beta1CspRdsInstancesResponseable when successful
// returns a V1beta1CspRdsInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspRdsInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspRdsInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspRdsInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspRdsInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspRdsInstancesRequestBuilderGetRequestConfiguration)(V1beta1CspRdsInstancesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspRdsInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspRdsInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspRdsInstances403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspRdsInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspRdsInstancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspRdsInstancesResponseable), nil
}
// GetAsCspRdsInstancesGetResponse returns a list of cloud service provider (CSP) RDS machine instances according to the given queryparameters for paging, filtering.
// returns a V1beta1CspRdsInstancesGetResponseable when successful
// returns a V1beta1CspRdsInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspRdsInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspRdsInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspRdsInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspRdsInstancesRequestBuilder) GetAsCspRdsInstancesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspRdsInstancesRequestBuilderGetRequestConfiguration)(V1beta1CspRdsInstancesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspRdsInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspRdsInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspRdsInstances403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspRdsInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspRdsInstancesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspRdsInstancesGetResponseable), nil
}
// ToGetRequestInformation returns a list of cloud service provider (CSP) RDS machine instances according to the given queryparameters for paging, filtering.
// returns a *RequestInformation when successful
func (m *V1beta1CspRdsInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspRdsInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspRdsInstancesRequestBuilder when successful
func (m *V1beta1CspRdsInstancesRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspRdsInstancesRequestBuilder) {
    return NewV1beta1CspRdsInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
