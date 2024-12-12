package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\data-services-connectors\{id}
type V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteQueryParameters this API will unregister the Data Services Connector from the DSCC customer account.
type V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteQueryParameters struct {
    // Force option can be used to forcefully unregister Data Services Connector which internally unregisters associated hypervisor managers
    Force *bool `uriparametername:"force"`
}
// V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteQueryParameters
}
// V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderInternal instantiates a new V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) {
    m := &V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/data-services-connectors/{id}{?force*}", pathParameters),
    }
    return m
}
// NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder instantiates a new V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete this API will unregister the Data Services Connector from the DSCC customer account.
// returns a UntypedNodeable when successful
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors503ErrorFromDiscriminatorValue,
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
// GenerateSupportBundle the generateSupportBundle property
// returns a *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) GenerateSupportBundle()(*V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) {
    return NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateTotp the generateTotp property
// returns a *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) GenerateTotp()(*V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) {
    return NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns details for a specified Data Services Connector
// Deprecated: This method is obsolete. Use GetAsDataServicesConnectorsGetResponse instead.
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsResponseable when successful
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderGetRequestConfiguration)(V1beta1DataServicesConnectorsItemDataServicesConnectorsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DataServicesConnectorsItemDataServicesConnectorsResponseable), nil
}
// GetAsDataServicesConnectorsGetResponse returns details for a specified Data Services Connector
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable when successful
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) GetAsDataServicesConnectorsGetResponse(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderGetRequestConfiguration)(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponseable), nil
}
// Patch update attributes of a specified Data Services Connector
// returns a UntypedNodeable when successful
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemDataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) Patch(ctx context.Context, body V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBodyable, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderPatchRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemDataServicesConnectors503ErrorFromDiscriminatorValue,
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
// SetRemoteSupport the setRemoteSupport property
// returns a *V1beta1DataServicesConnectorsItemSetRemoteSupportRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) SetRemoteSupport()(*V1beta1DataServicesConnectorsItemSetRemoteSupportRequestBuilder) {
    return NewV1beta1DataServicesConnectorsItemSetRemoteSupportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation this API will unregister the Data Services Connector from the DSCC customer account.
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns details for a specified Data Services Connector
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update attributes of a specified Data Services Connector
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1DataServicesConnectorsItemDataServicesConnectorsPatchRequestBodyable, requestConfiguration *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) {
    return NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
