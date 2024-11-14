package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems
type V1beta1SystemsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsRequestBuilderGetQueryParameters returns the systems and their properties defined by the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the ‘select’ query parameter to choose only the properties you want to retrieve for a system. For example, to get details of the id, name and software information for each system, use ‘?select=id,name,softwareInfo’
type V1beta1SystemsRequestBuilderGetQueryParameters struct {
    // The expression to filter responses.This API doesn't support filter by computerClusters, hypervisorClusters, softwareInfo/hypervisorClusters and systemVms collection properties.
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging, e.g.: offset=30&&limit=10. Limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // Use offset in conjunction with limit for paging, e.g.: offset=30&&limit=10. Offset is the number of items from the beginning of the result set to the first item included in the response.
    Offset *int32 `uriparametername:"offset"`
    // Query parameter listing the properties of system information to fetch.Although hypervisorClusters, computeClusters, systemVms and softwareInfo/hypervisorClusters can be selected, selecting elements of these collections is not supported.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a order indicator ("asc" or "desc"). Default order is ascending.This API doesn't support sorting by computerClusters, hypervisorClusters, softwareInfo/hypervisorClusters and systemVms collection properties.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SystemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems.privateCloudBusiness.v1beta1.systems.item collection
// returns a *V1beta1SystemsItemRequestBuilder when successful
func (m *V1beta1SystemsRequestBuilder) ById(id string)(*V1beta1SystemsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["%2Did"] = id
    }
    return NewV1beta1SystemsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SystemsRequestBuilderInternal instantiates a new V1beta1SystemsRequestBuilder and sets the default values.
func NewV1beta1SystemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsRequestBuilder) {
    m := &V1beta1SystemsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemsRequestBuilder instantiates a new V1beta1SystemsRequestBuilder and sets the default values.
func NewV1beta1SystemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the systems and their properties defined by the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the ‘select’ query parameter to choose only the properties you want to retrieve for a system. For example, to get details of the id, name and software information for each system, use ‘?select=id,name,softwareInfo’
// Deprecated: This method is obsolete. Use GetAsSystemsGetResponse instead.
// returns a V1beta1SystemsResponseable when successful
// returns a V1beta1Systems400Error error when the service returns a 400 status code
// returns a V1beta1Systems401Error error when the service returns a 401 status code
// returns a V1beta1Systems404Error error when the service returns a 404 status code
// returns a V1beta1Systems500Error error when the service returns a 500 status code
func (m *V1beta1SystemsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsRequestBuilderGetRequestConfiguration)(V1beta1SystemsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Systems400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Systems401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Systems404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Systems500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsResponseable), nil
}
// GetAsSystemsGetResponse returns the systems and their properties defined by the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the ‘select’ query parameter to choose only the properties you want to retrieve for a system. For example, to get details of the id, name and software information for each system, use ‘?select=id,name,softwareInfo’
// returns a V1beta1SystemsGetResponseable when successful
// returns a V1beta1Systems400Error error when the service returns a 400 status code
// returns a V1beta1Systems401Error error when the service returns a 401 status code
// returns a V1beta1Systems404Error error when the service returns a 404 status code
// returns a V1beta1Systems500Error error when the service returns a 500 status code
func (m *V1beta1SystemsRequestBuilder) GetAsSystemsGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsRequestBuilderGetRequestConfiguration)(V1beta1SystemsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Systems400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Systems401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Systems404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Systems500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsGetResponseable), nil
}
// ToGetRequestInformation returns the systems and their properties defined by the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the ‘select’ query parameter to choose only the properties you want to retrieve for a system. For example, to get details of the id, name and software information for each system, use ‘?select=id,name,softwareInfo’
// returns a *RequestInformation when successful
func (m *V1beta1SystemsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsRequestBuilder when successful
func (m *V1beta1SystemsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsRequestBuilder) {
    return NewV1beta1SystemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
