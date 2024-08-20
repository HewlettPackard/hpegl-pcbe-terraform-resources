package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}\folders\{folder-id}
type V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderInternal instantiates a new V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) {
    m := &V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}/folders/{folder%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder instantiates a new V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get details of a hypervisors folder.
// Deprecated: This method is obsolete. Use GetAsFolderGetResponse instead.
// returns a V1beta1HypervisorManagersItemFoldersItemFolderResponseable when successful
// returns a V1beta1HypervisorManagersItemFoldersItemFolder401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemFoldersItemFolder403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemFoldersItemFolder404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemFoldersItemFolder500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemFoldersItemFolderResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemFoldersItemFolder401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemFoldersItemFolder403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemFoldersItemFolder404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemFoldersItemFolder500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemFoldersItemFolderResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemFoldersItemFolderResponseable), nil
}
// GetAsFolderGetResponse details of a hypervisors folder.
// returns a V1beta1HypervisorManagersItemFoldersItemFolderGetResponseable when successful
// returns a V1beta1HypervisorManagersItemFoldersItemFolder401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemFoldersItemFolder403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemFoldersItemFolder404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemFoldersItemFolder500Error error when the service returns a 500 status code
func (m *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) GetAsFolderGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemFoldersItemFolderGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemFoldersItemFolder401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemFoldersItemFolder403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemFoldersItemFolder404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemFoldersItemFolder500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemFoldersItemFolderGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemFoldersItemFolderGetResponseable), nil
}
// ToGetRequestInformation details of a hypervisors folder.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder) {
    return NewV1beta1HypervisorManagersItemFoldersFolderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
