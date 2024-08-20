package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorHostsHostItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-hosts\{host-id}
type V1beta1HypervisorHostsHostItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorHostsHostItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorHostsHostItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorHostsHostItemRequestBuilderInternal instantiates a new V1beta1HypervisorHostsHostItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorHostsHostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorHostsHostItemRequestBuilder) {
    m := &V1beta1HypervisorHostsHostItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-hosts/{host%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorHostsHostItemRequestBuilder instantiates a new V1beta1HypervisorHostsHostItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorHostsHostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorHostsHostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorHostsHostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a hypervisors host.
// Deprecated: This method is obsolete. Use GetAsHostGetResponse instead.
// returns a V1beta1HypervisorHostsItemHostResponseable when successful
// returns a V1beta1HypervisorHostsItemHost401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorHostsItemHost403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorHostsItemHost404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorHostsItemHost500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorHostsHostItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorHostsHostItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorHostsItemHostResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorHostsItemHost401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorHostsItemHost403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorHostsItemHost404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorHostsItemHost500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorHostsItemHostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorHostsItemHostResponseable), nil
}
// GetAsHostGetResponse details of a hypervisors host.
// returns a V1beta1HypervisorHostsItemHostGetResponseable when successful
// returns a V1beta1HypervisorHostsItemHost401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorHostsItemHost403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorHostsItemHost404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorHostsItemHost500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorHostsHostItemRequestBuilder) GetAsHostGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorHostsHostItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorHostsItemHostGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorHostsItemHost401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorHostsItemHost403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorHostsItemHost404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorHostsItemHost500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorHostsItemHostGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorHostsItemHostGetResponseable), nil
}
// ToGetRequestInformation details of a hypervisors host.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorHostsHostItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorHostsHostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorHostsHostItemRequestBuilder when successful
func (m *V1beta1HypervisorHostsHostItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorHostsHostItemRequestBuilder) {
    return NewV1beta1HypervisorHostsHostItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
