package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-rds-instances\{id}
type V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderInternal instantiates a new V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder and sets the default values.
func NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) {
    m := &V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-rds-instances/{id}", pathParameters),
    }
    return m
}
// NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder instantiates a new V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder and sets the default values.
func NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns details of a specified cloud service provider (CSP) RDS machine instance.
// Deprecated: This method is obsolete. Use GetAsCspRdsInstancesGetResponse instead.
// returns a V1beta1CspRdsInstancesItemCspRdsInstancesResponseable when successful
// returns a V1beta1CspRdsInstancesItemCspRdsInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderGetRequestConfiguration)(V1beta1CspRdsInstancesItemCspRdsInstancesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspRdsInstancesItemCspRdsInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspRdsInstancesItemCspRdsInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspRdsInstancesItemCspRdsInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspRdsInstancesItemCspRdsInstances404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspRdsInstancesItemCspRdsInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspRdsInstancesItemCspRdsInstancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspRdsInstancesItemCspRdsInstancesResponseable), nil
}
// GetAsCspRdsInstancesGetResponse returns details of a specified cloud service provider (CSP) RDS machine instance.
// returns a V1beta1CspRdsInstancesItemCspRdsInstancesGetResponseable when successful
// returns a V1beta1CspRdsInstancesItemCspRdsInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspRdsInstancesItemCspRdsInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) GetAsCspRdsInstancesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderGetRequestConfiguration)(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspRdsInstancesItemCspRdsInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspRdsInstancesItemCspRdsInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspRdsInstancesItemCspRdsInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspRdsInstancesItemCspRdsInstances404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspRdsInstancesItemCspRdsInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponseable), nil
}
// Refresh the refresh property
// returns a *V1beta1CspRdsInstancesItemRefreshRequestBuilder when successful
func (m *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) Refresh()(*V1beta1CspRdsInstancesItemRefreshRequestBuilder) {
    return NewV1beta1CspRdsInstancesItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns details of a specified cloud service provider (CSP) RDS machine instance.
// returns a *RequestInformation when successful
func (m *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder when successful
func (m *V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder) {
    return NewV1beta1CspRdsInstancesCspRdsInstancesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
