package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\networks\{network-id}
type V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) {
    m := &V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/networks/{network%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder instantiates a new V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a hypervisors network resource.
// Deprecated: This method is obsolete. Use GetAsNetworkGetResponse instead.
// returns a V1beta1HypervisorManagersItemNetworksItemNetworkResponseable when successful
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemNetworksItemNetworkResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemNetworksItemNetwork401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemNetworksItemNetwork403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemNetworksItemNetwork404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemNetworksItemNetwork500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemNetworksItemNetworkResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemNetworksItemNetworkResponseable), nil
}
// GetAsNetworkGetResponse details of a hypervisors network resource.
// returns a V1beta1HypervisorManagersItemNetworksItemNetworkGetResponseable when successful
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemNetworksItemNetwork500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) GetAsNetworkGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemNetworksItemNetworkGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemNetworksItemNetwork401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemNetworksItemNetwork403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemNetworksItemNetwork404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemNetworksItemNetwork500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemNetworksItemNetworkGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemNetworksItemNetworkGetResponseable), nil
}
// ToGetRequestInformation details of a hypervisors network resource.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder) {
    return NewV1beta1HypervisorManagersItemNetworksNetworkItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
