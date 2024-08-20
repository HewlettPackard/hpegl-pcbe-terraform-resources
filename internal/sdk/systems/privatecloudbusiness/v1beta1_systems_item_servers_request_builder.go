package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemServersRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\servers
type V1beta1SystemsItemServersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemServersRequestBuilderGetQueryParameters get server information by system id Returns details about the servers for the specified system id. Retrieving all of the properties for all servers in a system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get the server id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
type V1beta1SystemsItemServersRequestBuilderGetQueryParameters struct {
    // The expression to filter responses.
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging, e.g.: offset=30&&limit=10. Limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // Use offset in conjunction with limit for paging, e.g.: offset=30&&limit=10. Offset is the number of items from the beginning of the result set to the first item included in the response.
    Offset *int32 `uriparametername:"offset"`
    // Query parameter listing the properties of Server information to fetch.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a directionindicator ("asc" or "desc"). If no direction indicator is specified thedefault order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SystemsItemServersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemServersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemsItemServersRequestBuilderGetQueryParameters
}
// ByServerId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems.privateCloudBusiness.v1beta1.systems.item.servers.item collection
// returns a *V1beta1SystemsItemServersWithServerItemRequestBuilder when successful
func (m *V1beta1SystemsItemServersRequestBuilder) ByServerId(serverId string)(*V1beta1SystemsItemServersWithServerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if serverId != "" {
        urlTplParams["serverId"] = serverId
    }
    return NewV1beta1SystemsItemServersWithServerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SystemsItemServersRequestBuilderInternal instantiates a new V1beta1SystemsItemServersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemServersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemServersRequestBuilder) {
    m := &V1beta1SystemsItemServersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/servers{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemServersRequestBuilder instantiates a new V1beta1SystemsItemServersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemServersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemServersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemServersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get server information by system id Returns details about the servers for the specified system id. Retrieving all of the properties for all servers in a system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get the server id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
// Deprecated: This method is obsolete. Use GetAsServersGetResponse instead.
// returns a V1beta1SystemsItemServersResponseable when successful
// returns a V1beta1SystemsItemServers400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemServers401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemServers404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemServers500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemServersRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemServersRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemServersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemServers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemServers401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemServers404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemServers500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemServersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemServersResponseable), nil
}
// GetAsServersGetResponse get server information by system id Returns details about the servers for the specified system id. Retrieving all of the properties for all servers in a system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get the server id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
// returns a V1beta1SystemsItemServersGetResponseable when successful
// returns a V1beta1SystemsItemServers400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemServers401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemServers404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemServers500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemServersRequestBuilder) GetAsServersGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemServersRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemServersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemServers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemServers401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemServers404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemServers500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemServersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemServersGetResponseable), nil
}
// ToGetRequestInformation get server information by system id Returns details about the servers for the specified system id. Retrieving all of the properties for all servers in a system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get the server id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemServersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemServersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemServersRequestBuilder when successful
func (m *V1beta1SystemsItemServersRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemServersRequestBuilder) {
    return NewV1beta1SystemsItemServersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
