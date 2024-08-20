package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemTagsTagItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\tags\{tag-id}
type V1beta1HypervisorManagersItemTagsTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemTagsTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemTagsTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemTagsTagItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) {
    m := &V1beta1HypervisorManagersItemTagsTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/tags/{tag%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilder instantiates a new V1beta1HypervisorManagersItemTagsTagItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a hypervisors tag resource.
// Deprecated: This method is obsolete. Use GetAsTagGetResponse instead.
// returns a V1beta1HypervisorManagersItemTagsItemTagResponseable when successful
// returns a V1beta1HypervisorManagersItemTagsItemTag401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemTagsItemTag403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemTagsItemTag404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemTagsItemTag500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemTagsTagItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemTagsItemTagResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemTagsItemTag401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemTagsItemTag403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemTagsItemTag404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemTagsItemTag500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemTagsItemTagResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemTagsItemTagResponseable), nil
}
// GetAsTagGetResponse details of a hypervisors tag resource.
// returns a V1beta1HypervisorManagersItemTagsItemTagGetResponseable when successful
// returns a V1beta1HypervisorManagersItemTagsItemTag401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemTagsItemTag403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemTagsItemTag404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemTagsItemTag500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) GetAsTagGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemTagsTagItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemTagsItemTagGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemTagsItemTag401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemTagsItemTag403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemTagsItemTag404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemTagsItemTag500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemTagsItemTagGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemTagsItemTagGetResponseable), nil
}
// ToGetRequestInformation details of a hypervisors tag resource.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemTagsTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorManagersItemTagsTagItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemTagsTagItemRequestBuilder) {
    return NewV1beta1HypervisorManagersItemTagsTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
