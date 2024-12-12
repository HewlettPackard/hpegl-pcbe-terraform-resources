package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder builds and executes requests for operations under \data-services\v1beta1\data-services-connectors\{id}\generate-support-bundle
type V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderInternal instantiates a new V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) {
    m := &V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/data-services-connectors/{id}/generate-support-bundle", pathParameters),
    }
    return m
}
// NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder instantiates a new V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generate a new support-bundle
// returns a UntypedNodeable when successful
// returns a V1beta1DataServicesConnectorsItemGenerateSupportBundle400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemGenerateSupportBundle401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemGenerateSupportBundle403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemGenerateSupportBundle404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemGenerateSupportBundle500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemGenerateSupportBundle503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) Post(ctx context.Context, body V1beta1DataServicesConnectorsItemGenerateSupportBundlePostRequestBodyable, requestConfiguration *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderPostRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemGenerateSupportBundle400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemGenerateSupportBundle401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemGenerateSupportBundle403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemGenerateSupportBundle404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemGenerateSupportBundle500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemGenerateSupportBundle503ErrorFromDiscriminatorValue,
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
// ToPostRequestInformation generate a new support-bundle
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1DataServicesConnectorsItemGenerateSupportBundlePostRequestBodyable, requestConfiguration *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) WithUrl(rawUrl string)(*V1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder) {
    return NewV1beta1DataServicesConnectorsItemGenerateSupportBundleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
