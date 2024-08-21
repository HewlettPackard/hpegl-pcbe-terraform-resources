package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorClustersClusterItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-clusters\{cluster-id}
type V1beta1HypervisorClustersClusterItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorClustersClusterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorClustersClusterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorClustersClusterItemRequestBuilderInternal instantiates a new V1beta1HypervisorClustersClusterItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorClustersClusterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorClustersClusterItemRequestBuilder) {
    m := &V1beta1HypervisorClustersClusterItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-clusters/{cluster%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorClustersClusterItemRequestBuilder instantiates a new V1beta1HypervisorClustersClusterItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorClustersClusterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorClustersClusterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorClustersClusterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a hypervisors cluster.
// Deprecated: This method is obsolete. Use GetAsClusterGetResponse instead.
// returns a V1beta1HypervisorClustersItemClusterResponseable when successful
// returns a V1beta1HypervisorClustersItemCluster401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorClustersItemCluster403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorClustersItemCluster404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorClustersItemCluster500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorClustersClusterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorClustersClusterItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorClustersItemClusterResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorClustersItemCluster401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorClustersItemCluster403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorClustersItemCluster404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorClustersItemCluster500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorClustersItemClusterResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorClustersItemClusterResponseable), nil
}
// GetAsClusterGetResponse details of a hypervisors cluster.
// returns a V1beta1HypervisorClustersItemClusterGetResponseable when successful
// returns a V1beta1HypervisorClustersItemCluster401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorClustersItemCluster403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorClustersItemCluster404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorClustersItemCluster500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorClustersClusterItemRequestBuilder) GetAsClusterGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorClustersClusterItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorClustersItemClusterGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorClustersItemCluster401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorClustersItemCluster403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorClustersItemCluster404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorClustersItemCluster500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorClustersItemClusterGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorClustersItemClusterGetResponseable), nil
}
// ToGetRequestInformation details of a hypervisors cluster.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorClustersClusterItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorClustersClusterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorClustersClusterItemRequestBuilder when successful
func (m *V1beta1HypervisorClustersClusterItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorClustersClusterItemRequestBuilder) {
    return NewV1beta1HypervisorClustersClusterItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}