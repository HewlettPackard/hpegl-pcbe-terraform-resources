package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresDatastoresItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\datastores\{id}
type V1beta1DatastoresDatastoresItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DatastoresDatastoresItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresDatastoresItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1DatastoresDatastoresItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresDatastoresItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1DatastoresDatastoresItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresDatastoresItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DatastoresDatastoresItemRequestBuilderInternal instantiates a new V1beta1DatastoresDatastoresItemRequestBuilder and sets the default values.
func NewV1beta1DatastoresDatastoresItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresDatastoresItemRequestBuilder) {
    m := &V1beta1DatastoresDatastoresItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/datastores/{id}", pathParameters),
    }
    return m
}
// NewV1beta1DatastoresDatastoresItemRequestBuilder instantiates a new V1beta1DatastoresDatastoresItemRequestBuilder and sets the default values.
func NewV1beta1DatastoresDatastoresItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresDatastoresItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DatastoresDatastoresItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the datastore.
// returns a UntypedNodeable when successful
// returns a V1beta1DatastoresItemDatastores400Error error when the service returns a 400 status code
// returns a V1beta1DatastoresItemDatastores401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemDatastores403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemDatastores404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemDatastores500Error error when the service returns a 500 status code
// returns a V1beta1DatastoresItemDatastores503Error error when the service returns a 503 status code
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderDeleteRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DatastoresItemDatastores400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DatastoresItemDatastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemDatastores403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemDatastores404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemDatastores500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DatastoresItemDatastores503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable), nil
}
// Get returns details for the specified datastore.
// Deprecated: This method is obsolete. Use GetAsDatastoresGetResponse instead.
// returns a V1beta1DatastoresItemDatastoresResponseable when successful
// returns a V1beta1DatastoresItemDatastores401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemDatastores403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemDatastores404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemDatastores500Error error when the service returns a 500 status code
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderGetRequestConfiguration)(V1beta1DatastoresItemDatastoresResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1DatastoresItemDatastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemDatastores403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemDatastores404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemDatastores500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresItemDatastoresResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresItemDatastoresResponseable), nil
}
// GetAsDatastoresGetResponse returns details for the specified datastore.
// returns a V1beta1DatastoresItemDatastoresGetResponseable when successful
// returns a V1beta1DatastoresItemDatastores401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemDatastores403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemDatastores404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemDatastores500Error error when the service returns a 500 status code
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) GetAsDatastoresGetResponse(ctx context.Context, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderGetRequestConfiguration)(V1beta1DatastoresItemDatastoresGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1DatastoresItemDatastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemDatastores403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemDatastores404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemDatastores500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresItemDatastoresGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresItemDatastoresGetResponseable), nil
}
// Mount the mount property
// returns a *V1beta1DatastoresItemMountRequestBuilder when successful
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) Mount()(*V1beta1DatastoresItemMountRequestBuilder) {
    return NewV1beta1DatastoresItemMountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update attributes of a specified datastore
// returns a UntypedNodeable when successful
// returns a V1beta1DatastoresItemDatastores400Error error when the service returns a 400 status code
// returns a V1beta1DatastoresItemDatastores401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemDatastores403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemDatastores404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemDatastores500Error error when the service returns a 500 status code
// returns a V1beta1DatastoresItemDatastores503Error error when the service returns a 503 status code
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) Patch(ctx context.Context, body V1beta1DatastoresItemDatastoresPatchRequestBodyable, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderPatchRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DatastoresItemDatastores400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DatastoresItemDatastores401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemDatastores403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemDatastores404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemDatastores500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DatastoresItemDatastores503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable), nil
}
// ToDeleteRequestInformation deletes the datastore.
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns details for the specified datastore.
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update attributes of a specified datastore
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1DatastoresItemDatastoresPatchRequestBodyable, requestConfiguration *V1beta1DatastoresDatastoresItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/merge-patch+json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Unmount the unmount property
// returns a *V1beta1DatastoresItemUnmountRequestBuilder when successful
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) Unmount()(*V1beta1DatastoresItemUnmountRequestBuilder) {
    return NewV1beta1DatastoresItemUnmountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1DatastoresDatastoresItemRequestBuilder when successful
func (m *V1beta1DatastoresDatastoresItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1DatastoresDatastoresItemRequestBuilder) {
    return NewV1beta1DatastoresDatastoresItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
