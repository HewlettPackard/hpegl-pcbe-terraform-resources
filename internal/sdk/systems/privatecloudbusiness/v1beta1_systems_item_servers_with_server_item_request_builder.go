package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemServersWithServerItemRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\servers\{serverId}
type V1beta1SystemsItemServersWithServerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemServersWithServerItemRequestBuilderGetQueryParameters get server details by system id and server id Returns information about a specific server on a specific system. Retrieving all of the properties for a server can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want to retrieve. For example, to get details of the server's id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
type V1beta1SystemsItemServersWithServerItemRequestBuilderGetQueryParameters struct {
    // Query parameter listing the properties of Server information to fetch.
    Select *string `uriparametername:"select"`
}
// V1beta1SystemsItemServersWithServerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemServersWithServerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemsItemServersWithServerItemRequestBuilderGetQueryParameters
}
// NewV1beta1SystemsItemServersWithServerItemRequestBuilderInternal instantiates a new V1beta1SystemsItemServersWithServerItemRequestBuilder and sets the default values.
func NewV1beta1SystemsItemServersWithServerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemServersWithServerItemRequestBuilder) {
    m := &V1beta1SystemsItemServersWithServerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/servers/{serverId}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemServersWithServerItemRequestBuilder instantiates a new V1beta1SystemsItemServersWithServerItemRequestBuilder and sets the default values.
func NewV1beta1SystemsItemServersWithServerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemServersWithServerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemServersWithServerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get server details by system id and server id Returns information about a specific server on a specific system. Retrieving all of the properties for a server can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want to retrieve. For example, to get details of the server's id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
// Deprecated: This method is obsolete. Use GetAsWithServerGetResponse instead.
// returns a V1beta1SystemsItemServersItemWithServerResponseable when successful
// returns a V1beta1SystemsItemServersItemWithServer400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemServersItemWithServer401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemServersItemWithServer404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemServersItemWithServer500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemServersWithServerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemServersWithServerItemRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemServersItemWithServerResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemServersItemWithServer400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemServersItemWithServer401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemServersItemWithServer404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemServersItemWithServer500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemServersItemWithServerResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemServersItemWithServerResponseable), nil
}
// GetAsWithServerGetResponse get server details by system id and server id Returns information about a specific server on a specific system. Retrieving all of the properties for a server can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want to retrieve. For example, to get details of the server's id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
// returns a V1beta1SystemsItemServersItemWithServerGetResponseable when successful
// returns a V1beta1SystemsItemServersItemWithServer400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemServersItemWithServer401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemServersItemWithServer404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemServersItemWithServer500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemServersWithServerItemRequestBuilder) GetAsWithServerGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemServersWithServerItemRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemServersItemWithServerGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemServersItemWithServer400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemServersItemWithServer401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemServersItemWithServer404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemServersItemWithServer500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemServersItemWithServerGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemServersItemWithServerGetResponseable), nil
}
// ToGetRequestInformation get server details by system id and server id Returns information about a specific server on a specific system. Retrieving all of the properties for a server can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want to retrieve. For example, to get details of the server's id, name, serial number and hypervisor host, use ?select=id,name,serialNumber,hypervisorHost
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemServersWithServerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemServersWithServerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemServersWithServerItemRequestBuilder when successful
func (m *V1beta1SystemsItemServersWithServerItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemServersWithServerItemRequestBuilder) {
    return NewV1beta1SystemsItemServersWithServerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
