package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1SoftwareUpgradesRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-upgrades
type V1beta1SoftwareUpgradesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SoftwareUpgradesRequestBuilderGetQueryParameters list the available upgrades for an installed Software Release, identified by a version andSoftware Component ID. Pagination is not supported on the returned items.The returned items are upgrades that can be applied to an existing installation of a SoftwareRelease. Some releases can be applied immediately, while others may require correctiveactions to be completed before the release is allowed.Upgrades beyond those initially returned can be found by recursively calling this API withthe new Software Release version. This can be useful for presenting the series of updatesrequired to bring an installation to the latest version.
type V1beta1SoftwareUpgradesRequestBuilderGetQueryParameters struct {
    // The identifier for the virtual machine being upgraded. Either `agent-id` or `serial-number`must be supplied.
    AgentId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID `uriparametername:"agent%2Did"`
    // The serial number of the hardware being upgraded. Either `agent-id` or `serial-number` mustbe supplied.
    SerialNumber *string `uriparametername:"serial%2Dnumber"`
    // The ID of a Software Component.
    SoftwareComponentId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID `uriparametername:"software%2Dcomponent%2Did"`
    // The version of an installed Software Release that is being upgraded.
    Version *string `uriparametername:"version"`
}
// V1beta1SoftwareUpgradesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SoftwareUpgradesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SoftwareUpgradesRequestBuilderGetQueryParameters
}
// NewV1beta1SoftwareUpgradesRequestBuilderInternal instantiates a new V1beta1SoftwareUpgradesRequestBuilder and sets the default values.
func NewV1beta1SoftwareUpgradesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareUpgradesRequestBuilder) {
    m := &V1beta1SoftwareUpgradesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-upgrades?software-component-id={software%2Dcomponent%2Did}&version={version}{&agent%2Did*,serial%2Dnumber*}", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareUpgradesRequestBuilder instantiates a new V1beta1SoftwareUpgradesRequestBuilder and sets the default values.
func NewV1beta1SoftwareUpgradesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareUpgradesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareUpgradesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the available upgrades for an installed Software Release, identified by a version andSoftware Component ID. Pagination is not supported on the returned items.The returned items are upgrades that can be applied to an existing installation of a SoftwareRelease. Some releases can be applied immediately, while others may require correctiveactions to be completed before the release is allowed.Upgrades beyond those initially returned can be found by recursively calling this API withthe new Software Release version. This can be useful for presenting the series of updatesrequired to bring an installation to the latest version.
// Deprecated: This method is obsolete. Use GetAsSoftwareUpgradesGetResponse instead.
// returns a V1beta1SoftwareUpgradesResponseable when successful
// returns a V1beta1SoftwareUpgrades400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareUpgrades401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareUpgrades404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareUpgrades500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareUpgrades503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareUpgradesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SoftwareUpgradesRequestBuilderGetRequestConfiguration)(V1beta1SoftwareUpgradesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareUpgrades400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareUpgrades401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareUpgrades404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareUpgrades500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareUpgrades503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareUpgradesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareUpgradesResponseable), nil
}
// GetAsSoftwareUpgradesGetResponse list the available upgrades for an installed Software Release, identified by a version andSoftware Component ID. Pagination is not supported on the returned items.The returned items are upgrades that can be applied to an existing installation of a SoftwareRelease. Some releases can be applied immediately, while others may require correctiveactions to be completed before the release is allowed.Upgrades beyond those initially returned can be found by recursively calling this API withthe new Software Release version. This can be useful for presenting the series of updatesrequired to bring an installation to the latest version.
// returns a V1beta1SoftwareUpgradesGetResponseable when successful
// returns a V1beta1SoftwareUpgrades400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareUpgrades401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareUpgrades404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareUpgrades500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareUpgrades503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareUpgradesRequestBuilder) GetAsSoftwareUpgradesGetResponse(ctx context.Context, requestConfiguration *V1beta1SoftwareUpgradesRequestBuilderGetRequestConfiguration)(V1beta1SoftwareUpgradesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareUpgrades400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareUpgrades401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareUpgrades404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareUpgrades500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareUpgrades503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareUpgradesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareUpgradesGetResponseable), nil
}
// ToGetRequestInformation list the available upgrades for an installed Software Release, identified by a version andSoftware Component ID. Pagination is not supported on the returned items.The returned items are upgrades that can be applied to an existing installation of a SoftwareRelease. Some releases can be applied immediately, while others may require correctiveactions to be completed before the release is allowed.Upgrades beyond those initially returned can be found by recursively calling this API withthe new Software Release version. This can be useful for presenting the series of updatesrequired to bring an installation to the latest version.
// returns a *RequestInformation when successful
func (m *V1beta1SoftwareUpgradesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SoftwareUpgradesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SoftwareUpgradesRequestBuilder when successful
func (m *V1beta1SoftwareUpgradesRequestBuilder) WithUrl(rawUrl string)(*V1beta1SoftwareUpgradesRequestBuilder) {
    return NewV1beta1SoftwareUpgradesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
