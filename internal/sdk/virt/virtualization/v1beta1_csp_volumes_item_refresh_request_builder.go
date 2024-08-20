package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspVolumesItemRefreshRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-volumes\{id}\refresh
type V1beta1CspVolumesItemRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspVolumesItemRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspVolumesItemRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspVolumesItemRefreshRequestBuilderInternal instantiates a new V1beta1CspVolumesItemRefreshRequestBuilder and sets the default values.
func NewV1beta1CspVolumesItemRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesItemRefreshRequestBuilder) {
    m := &V1beta1CspVolumesItemRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-volumes/{id}/refresh", pathParameters),
    }
    return m
}
// NewV1beta1CspVolumesItemRefreshRequestBuilder instantiates a new V1beta1CspVolumesItemRefreshRequestBuilder and sets the default values.
func NewV1beta1CspVolumesItemRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesItemRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspVolumesItemRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates the properties of a specified volume to match the settings in the cloud account. Also updates the CSP properties of any CSP machine instances to which it is attached. Protection groups with dynamic membership are updated to include or exclude the volume according to its updated settings.
// Deprecated: This method is obsolete. Use PostAsRefreshPostResponse instead.
// returns a V1beta1CspVolumesItemRefreshResponseable when successful
// returns a V1beta1CspVolumesItemRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumesItemRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumesItemRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspVolumesItemRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspVolumesItemRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesItemRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspVolumesItemRefreshRequestBuilderPostRequestConfiguration)(V1beta1CspVolumesItemRefreshResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspVolumesItemRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumesItemRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspVolumesItemRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspVolumesItemRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumesItemRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesItemRefreshResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesItemRefreshResponseable), nil
}
// PostAsRefreshPostResponse updates the properties of a specified volume to match the settings in the cloud account. Also updates the CSP properties of any CSP machine instances to which it is attached. Protection groups with dynamic membership are updated to include or exclude the volume according to its updated settings.
// returns a V1beta1CspVolumesItemRefreshPostResponseable when successful
// returns a V1beta1CspVolumesItemRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumesItemRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumesItemRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspVolumesItemRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspVolumesItemRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesItemRefreshRequestBuilder) PostAsRefreshPostResponse(ctx context.Context, requestConfiguration *V1beta1CspVolumesItemRefreshRequestBuilderPostRequestConfiguration)(V1beta1CspVolumesItemRefreshPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspVolumesItemRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumesItemRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspVolumesItemRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspVolumesItemRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumesItemRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesItemRefreshPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesItemRefreshPostResponseable), nil
}
// ToPostRequestInformation updates the properties of a specified volume to match the settings in the cloud account. Also updates the CSP properties of any CSP machine instances to which it is attached. Protection groups with dynamic membership are updated to include or exclude the volume according to its updated settings.
// returns a *RequestInformation when successful
func (m *V1beta1CspVolumesItemRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspVolumesItemRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspVolumesItemRefreshRequestBuilder when successful
func (m *V1beta1CspVolumesItemRefreshRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspVolumesItemRefreshRequestBuilder) {
    return NewV1beta1CspVolumesItemRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
