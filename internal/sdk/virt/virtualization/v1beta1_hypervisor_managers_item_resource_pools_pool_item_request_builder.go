package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\resource-pools\{pool-id}
type V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) {
    m := &V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/resource-pools/{pool%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder instantiates a new V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a hypervisors resource pool.
// Deprecated: This method is obsolete. Use GetAsPoolGetResponse instead.
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPoolResponseable when successful
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemResourcePoolsItemPoolResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemResourcePoolsItemPoolResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemResourcePoolsItemPoolResponseable), nil
}
// GetAsPoolGetResponse details of a hypervisors resource pool.
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseable when successful
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemResourcePoolsItemPool500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) GetAsPoolGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemResourcePoolsItemPool500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemResourcePoolsItemPoolGetResponseable), nil
}
// ToGetRequestInformation details of a hypervisors resource pool.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder) {
    return NewV1beta1HypervisorManagersItemResourcePoolsPoolItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
