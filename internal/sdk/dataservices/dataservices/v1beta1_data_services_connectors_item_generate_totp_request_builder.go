package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder builds and executes requests for operations under \data-services\v1beta1\data-services-connectors\{id}\generate-totp
type V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderInternal instantiates a new V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) {
    m := &V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/data-services-connectors/{id}/generate-totp", pathParameters),
    }
    return m
}
// NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder instantiates a new V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generate a TOTP to access DSC on-prem UI
// Deprecated: This method is obsolete. Use PostAsGenerateTotpPostResponse instead.
// returns a V1beta1DataServicesConnectorsItemGenerateTotpResponseable when successful
// returns a V1beta1DataServicesConnectorsItemGenerateTotp400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderPostRequestConfiguration)(V1beta1DataServicesConnectorsItemGenerateTotpResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemGenerateTotp400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemGenerateTotp401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemGenerateTotp403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemGenerateTotp404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemGenerateTotp500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemGenerateTotp503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DataServicesConnectorsItemGenerateTotpResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DataServicesConnectorsItemGenerateTotpResponseable), nil
}
// PostAsGenerateTotpPostResponse generate a TOTP to access DSC on-prem UI
// returns a V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable when successful
// returns a V1beta1DataServicesConnectorsItemGenerateTotp400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp404Error error when the service returns a 404 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectorsItemGenerateTotp503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) PostAsGenerateTotpPostResponse(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderPostRequestConfiguration)(V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectorsItemGenerateTotp400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectorsItemGenerateTotp401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectorsItemGenerateTotp403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DataServicesConnectorsItemGenerateTotp404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectorsItemGenerateTotp500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectorsItemGenerateTotp503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DataServicesConnectorsItemGenerateTotpPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DataServicesConnectorsItemGenerateTotpPostResponseable), nil
}
// ToPostRequestInformation generate a TOTP to access DSC on-prem UI
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) WithUrl(rawUrl string)(*V1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder) {
    return NewV1beta1DataServicesConnectorsItemGenerateTotpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
