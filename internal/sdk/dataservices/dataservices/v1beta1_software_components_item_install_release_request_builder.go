package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-components\{id}\install-release
type V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SoftwareComponentsItemInstallReleaseRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SoftwareComponentsItemInstallReleaseRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilderInternal instantiates a new V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder and sets the default values.
func NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) {
    m := &V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-components/{id}/install-release", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilder instantiates a new V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder and sets the default values.
func NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilderInternal(urlParams, requestAdapter)
}
// Get find the latest stable Software Release that can be installed for the specified Software Component. A Software Release is considered stable if the stability property is set to `GENERAL_AVAILABILITY`.
// Deprecated: This method is obsolete. Use GetAsInstallReleaseGetResponse instead.
// returns a V1beta1SoftwareComponentsItemInstallReleaseResponseable when successful
// returns a V1beta1SoftwareComponentsItemInstallRelease400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilderGetRequestConfiguration)(V1beta1SoftwareComponentsItemInstallReleaseResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareComponentsItemInstallRelease400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareComponentsItemInstallRelease401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareComponentsItemInstallRelease404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareComponentsItemInstallRelease500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareComponentsItemInstallRelease503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareComponentsItemInstallReleaseResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareComponentsItemInstallReleaseResponseable), nil
}
// GetAsInstallReleaseGetResponse find the latest stable Software Release that can be installed for the specified Software Component. A Software Release is considered stable if the stability property is set to `GENERAL_AVAILABILITY`.
// returns a V1beta1SoftwareComponentsItemInstallReleaseGetResponseable when successful
// returns a V1beta1SoftwareComponentsItemInstallRelease400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareComponentsItemInstallRelease503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) GetAsInstallReleaseGetResponse(ctx context.Context, requestConfiguration *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilderGetRequestConfiguration)(V1beta1SoftwareComponentsItemInstallReleaseGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareComponentsItemInstallRelease400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareComponentsItemInstallRelease401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareComponentsItemInstallRelease404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareComponentsItemInstallRelease500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareComponentsItemInstallRelease503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareComponentsItemInstallReleaseGetResponseable), nil
}
// ToGetRequestInformation find the latest stable Software Release that can be installed for the specified Software Component. A Software Release is considered stable if the stability property is set to `GENERAL_AVAILABILITY`.
// returns a *RequestInformation when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) WithUrl(rawUrl string)(*V1beta1SoftwareComponentsItemInstallReleaseRequestBuilder) {
    return NewV1beta1SoftwareComponentsItemInstallReleaseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
