package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}
type V1beta1SystemsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemRequestBuilderGetQueryParameters returns the system properties specified in the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get details of the system's id, name and software information, use '?select=id,name,softwareInfo'
type V1beta1SystemsItemRequestBuilderGetQueryParameters struct {
    // Query parameter listing the properties of system information to fetch.Although Hypervisor Clusters collection (property hypervisorClusters) can be selected, selecting elements of the collection is not supported.Similarly, hypervisor clusters update status collection (property softwareInfo.hypervisorClusters) can be selected, but, selecting elements of the collection is not supported in the select query parameter.
    Select *string `uriparametername:"select"`
}
// V1beta1SystemsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemsItemRequestBuilderGetQueryParameters
}
// V1beta1SystemsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddHypervisorCluster the addHypervisorCluster property
// returns a *V1beta1SystemsItemAddHypervisorClusterRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) AddHypervisorCluster()(*V1beta1SystemsItemAddHypervisorClusterRequestBuilder) {
    return NewV1beta1SystemsItemAddHypervisorClusterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SystemsItemRequestBuilderInternal instantiates a new V1beta1SystemsItemRequestBuilder and sets the default values.
func NewV1beta1SystemsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemRequestBuilder) {
    m := &V1beta1SystemsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemRequestBuilder instantiates a new V1beta1SystemsItemRequestBuilder and sets the default values.
func NewV1beta1SystemsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the system properties specified in the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get details of the system's id, name and software information, use '?select=id,name,softwareInfo'
// Deprecated: This method is obsolete. Use GetAsGetResponse instead.
// returns a V1beta1SystemsItemResponseable when successful
// returns a V1beta1SystemsItemFourZeroZeroError error when the service returns a 400 status code
// returns a V1beta1SystemsItemFourZeroOneError error when the service returns a 401 status code
// returns a V1beta1SystemsItemFourZeroFourError error when the service returns a 404 status code
// returns a V1beta1SystemsItemFiveZeroZeroError error when the service returns a 500 status code
func (m *V1beta1SystemsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemFourZeroZeroErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemFourZeroOneErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemFourZeroFourErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemFiveZeroZeroErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemResponseable), nil
}
// GetAsGetResponse returns the system properties specified in the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get details of the system's id, name and software information, use '?select=id,name,softwareInfo'
// returns a V1beta1SystemsItemGetResponseable when successful
// returns a V1beta1SystemsItemFourZeroZeroError error when the service returns a 400 status code
// returns a V1beta1SystemsItemFourZeroOneError error when the service returns a 401 status code
// returns a V1beta1SystemsItemFourZeroFourError error when the service returns a 404 status code
// returns a V1beta1SystemsItemFiveZeroZeroError error when the service returns a 500 status code
func (m *V1beta1SystemsItemRequestBuilder) GetAsGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemFourZeroZeroErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemFourZeroOneErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemFourZeroFourErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemFiveZeroZeroErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemGetResponseable), nil
}
// Patch modify specified system.
// Deprecated: This method is obsolete. Use PatchAsPatchResponse instead.
// returns a V1beta1SystemsItemResponseable when successful
// returns a V1beta1SystemsItemFourZeroZeroError error when the service returns a 400 status code
// returns a V1beta1SystemsItemFourZeroOneError error when the service returns a 401 status code
// returns a V1beta1SystemsItemFourZeroThreeError error when the service returns a 403 status code
// returns a V1beta1SystemsItemFourZeroFourError error when the service returns a 404 status code
// returns a V1beta1SystemsItemFiveZeroZeroError error when the service returns a 500 status code
// returns a V1beta1SystemsItemFiveZeroThreeError error when the service returns a 503 status code
func (m *V1beta1SystemsItemRequestBuilder) Patch(ctx context.Context, body V1beta1SystemsItemPatchRequestBodyable, requestConfiguration *V1beta1SystemsItemRequestBuilderPatchRequestConfiguration)(V1beta1SystemsItemResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemFourZeroZeroErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemFourZeroOneErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemFourZeroThreeErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemFourZeroFourErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemFiveZeroZeroErrorFromDiscriminatorValue,
        "503": CreateV1beta1SystemsItemFiveZeroThreeErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemResponseable), nil
}
// PatchAsPatchResponse modify specified system.
// returns a V1beta1SystemsItemPatchResponseable when successful
// returns a V1beta1SystemsItemFourZeroZeroError error when the service returns a 400 status code
// returns a V1beta1SystemsItemFourZeroOneError error when the service returns a 401 status code
// returns a V1beta1SystemsItemFourZeroThreeError error when the service returns a 403 status code
// returns a V1beta1SystemsItemFourZeroFourError error when the service returns a 404 status code
// returns a V1beta1SystemsItemFiveZeroZeroError error when the service returns a 500 status code
// returns a V1beta1SystemsItemFiveZeroThreeError error when the service returns a 503 status code
func (m *V1beta1SystemsItemRequestBuilder) PatchAsPatchResponse(ctx context.Context, body V1beta1SystemsItemPatchRequestBodyable, requestConfiguration *V1beta1SystemsItemRequestBuilderPatchRequestConfiguration)(V1beta1SystemsItemPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemFourZeroZeroErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemFourZeroOneErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemFourZeroThreeErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemFourZeroFourErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemFiveZeroZeroErrorFromDiscriminatorValue,
        "503": CreateV1beta1SystemsItemFiveZeroThreeErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemPatchResponseable), nil
}
// Servers the servers property
// returns a *V1beta1SystemsItemServersRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) Servers()(*V1beta1SystemsItemServersRequestBuilder) {
    return NewV1beta1SystemsItemServersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwarePrechecks the softwarePrechecks property
// returns a *V1beta1SystemsItemSoftwarePrechecksRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) SoftwarePrechecks()(*V1beta1SystemsItemSoftwarePrechecksRequestBuilder) {
    return NewV1beta1SystemsItemSoftwarePrechecksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareUpdate the softwareUpdate property
// returns a *V1beta1SystemsItemSoftwareUpdateRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) SoftwareUpdate()(*V1beta1SystemsItemSoftwareUpdateRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareUpdateResume the softwareUpdateResume property
// returns a *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) SoftwareUpdateResume()(*V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareVersionRefresh the softwareVersionRefresh property
// returns a *V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) SoftwareVersionRefresh()(*V1beta1SystemsItemSoftwareVersionRefreshRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareVersionRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StorageArrays the storageArrays property
// returns a *V1beta1SystemsItemStorageArraysRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) StorageArrays()(*V1beta1SystemsItemStorageArraysRequestBuilder) {
    return NewV1beta1SystemsItemStorageArraysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StoragePools the storagePools property
// returns a *V1beta1SystemsItemStoragePoolsRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) StoragePools()(*V1beta1SystemsItemStoragePoolsRequestBuilder) {
    return NewV1beta1SystemsItemStoragePoolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StorageReplicationPartners the storageReplicationPartners property
// returns a *V1beta1SystemsItemStorageReplicationPartnersRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) StorageReplicationPartners()(*V1beta1SystemsItemStorageReplicationPartnersRequestBuilder) {
    return NewV1beta1SystemsItemStorageReplicationPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Switches the switches property
// returns a *V1beta1SystemsItemSwitchesRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) Switches()(*V1beta1SystemsItemSwitchesRequestBuilder) {
    return NewV1beta1SystemsItemSwitchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation returns the system properties specified in the query parameters. Retrieving all of the properties for the system can take a long time because the amount of data is large. Use the  'select' query parameter to choose only the properties you want to retrieve.For example, to get details of the system's id, name and software information, use '?select=id,name,softwareInfo'
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation modify specified system.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1SystemsItemPatchRequestBodyable, requestConfiguration *V1beta1SystemsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemRequestBuilder when successful
func (m *V1beta1SystemsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemRequestBuilder) {
    return NewV1beta1SystemsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
