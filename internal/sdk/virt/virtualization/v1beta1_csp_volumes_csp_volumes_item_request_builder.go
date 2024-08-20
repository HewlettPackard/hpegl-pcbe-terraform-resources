package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspVolumesCspVolumesItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-volumes\{id}
type V1beta1CspVolumesCspVolumesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspVolumesCspVolumesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspVolumesCspVolumesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspVolumesCspVolumesItemRequestBuilderInternal instantiates a new V1beta1CspVolumesCspVolumesItemRequestBuilder and sets the default values.
func NewV1beta1CspVolumesCspVolumesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesCspVolumesItemRequestBuilder) {
    m := &V1beta1CspVolumesCspVolumesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-volumes/{id}", pathParameters),
    }
    return m
}
// NewV1beta1CspVolumesCspVolumesItemRequestBuilder instantiates a new V1beta1CspVolumesCspVolumesItemRequestBuilder and sets the default values.
func NewV1beta1CspVolumesCspVolumesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesCspVolumesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspVolumesCspVolumesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns details of a specified cloud service provider (CSP) volume.
// Deprecated: This method is obsolete. Use GetAsCspVolumesGetResponse instead.
// returns a V1beta1CspVolumesItemCspVolumesResponseable when successful
// returns a V1beta1CspVolumesItemCspVolumes400Error error when the service returns a 400 status code
// returns a V1beta1CspVolumesItemCspVolumes401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumesItemCspVolumes403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumesItemCspVolumes404Error error when the service returns a 404 status code
// returns a V1beta1CspVolumesItemCspVolumes500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesCspVolumesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspVolumesCspVolumesItemRequestBuilderGetRequestConfiguration)(V1beta1CspVolumesItemCspVolumesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspVolumesItemCspVolumes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspVolumesItemCspVolumes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumesItemCspVolumes403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspVolumesItemCspVolumes404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumesItemCspVolumes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesItemCspVolumesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesItemCspVolumesResponseable), nil
}
// GetAsCspVolumesGetResponse returns details of a specified cloud service provider (CSP) volume.
// returns a V1beta1CspVolumesItemCspVolumesGetResponseable when successful
// returns a V1beta1CspVolumesItemCspVolumes400Error error when the service returns a 400 status code
// returns a V1beta1CspVolumesItemCspVolumes401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumesItemCspVolumes403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumesItemCspVolumes404Error error when the service returns a 404 status code
// returns a V1beta1CspVolumesItemCspVolumes500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesCspVolumesItemRequestBuilder) GetAsCspVolumesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspVolumesCspVolumesItemRequestBuilderGetRequestConfiguration)(V1beta1CspVolumesItemCspVolumesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspVolumesItemCspVolumes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspVolumesItemCspVolumes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumesItemCspVolumes403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspVolumesItemCspVolumes404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumesItemCspVolumes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesItemCspVolumesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesItemCspVolumesGetResponseable), nil
}
// Refresh the refresh property
// returns a *V1beta1CspVolumesItemRefreshRequestBuilder when successful
func (m *V1beta1CspVolumesCspVolumesItemRequestBuilder) Refresh()(*V1beta1CspVolumesItemRefreshRequestBuilder) {
    return NewV1beta1CspVolumesItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns details of a specified cloud service provider (CSP) volume.
// returns a *RequestInformation when successful
func (m *V1beta1CspVolumesCspVolumesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspVolumesCspVolumesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspVolumesCspVolumesItemRequestBuilder when successful
func (m *V1beta1CspVolumesCspVolumesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspVolumesCspVolumesItemRequestBuilder) {
    return NewV1beta1CspVolumesCspVolumesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
