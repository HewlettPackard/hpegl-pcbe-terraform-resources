package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\hypervisor-library-images\{hypervisor-library-image-id}
type V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) {
    m := &V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/hypervisor-library-images/{hypervisor%2Dlibrary%2Dimage%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder instantiates a new V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a virtual machine image
// Deprecated: This method is obsolete. Use GetAsHypervisorLibraryImageGetResponse instead.
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageResponseable when successful
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageResponseable), nil
}
// GetAsHypervisorLibraryImageGetResponse details of a virtual machine image
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseable when successful
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) GetAsHypervisorLibraryImageGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImage500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseable), nil
}
// ToGetRequestInformation details of a virtual machine image
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder) {
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesHypervisorLibraryImageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
