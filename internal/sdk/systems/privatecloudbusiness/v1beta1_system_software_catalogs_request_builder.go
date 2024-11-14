package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemSoftwareCatalogsRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\system-software-catalogs
type V1beta1SystemSoftwareCatalogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemSoftwareCatalogsRequestBuilderGetQueryParameters returns a list of all system software catalogs.Use 'select' and 'filter' query parameters to customize the response returned by this API.For example, to get the End User License Agreement (EULA) for catalog version 1.2.3.4, use '?select=eula&filter=version eq 1.2.3.4'. To get recommended systems for precheck to a given catalog version, use '?select=systemsWithUpdatePath&filter=version eq 1.2.3.4'.
type V1beta1SystemSoftwareCatalogsRequestBuilderGetQueryParameters struct {
    // The expression to filter responses.The property names which are of type string should be passed in quotes('') and nested property names should be passed with "/" as the separator.Filtering is supported with following properties: * eula * createdAt * customerId * generation * id * name * resourceUri * type * updatedAt * hypervisor/name * hypervisor/releaseDate * hypervisor/releaseNotesUrl * hypervisor/version * releaseDate * serverFirmware/name * serverFirmware/releaseDate * serverFirmware/releaseNotesUrl * serverFirmware/version * storageConnectionManager/name * storageConnectionManager/releaseDate * storageConnectionManager/releaseNotesUrl * storageConnectionManager/version * storageSoftware/name * storageSoftware/releaseDate * storageSoftware/releaseNotesUrl * storageSoftware/version * aiEssentials/name * aiEssentials/releaseDate * aiEssentials/releaseNotesUrl * aiEssentials/version * version
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging, e.g.: offset=30&&limit=10. Limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // A list of properties in the items collection to include in the response.
    Select *string `uriparametername:"select"`
}
// V1beta1SystemSoftwareCatalogsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemSoftwareCatalogsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemSoftwareCatalogsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems.privateCloudBusiness.v1beta1.systemSoftwareCatalogs.item collection
// returns a *V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder when successful
func (m *V1beta1SystemSoftwareCatalogsRequestBuilder) ById(id string)(*V1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1SystemSoftwareCatalogsSystemSoftwareCatalogsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SystemSoftwareCatalogsRequestBuilderInternal instantiates a new V1beta1SystemSoftwareCatalogsRequestBuilder and sets the default values.
func NewV1beta1SystemSoftwareCatalogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemSoftwareCatalogsRequestBuilder) {
    m := &V1beta1SystemSoftwareCatalogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/system-software-catalogs{?filter*,limit*,select*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemSoftwareCatalogsRequestBuilder instantiates a new V1beta1SystemSoftwareCatalogsRequestBuilder and sets the default values.
func NewV1beta1SystemSoftwareCatalogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemSoftwareCatalogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemSoftwareCatalogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of all system software catalogs.Use 'select' and 'filter' query parameters to customize the response returned by this API.For example, to get the End User License Agreement (EULA) for catalog version 1.2.3.4, use '?select=eula&filter=version eq 1.2.3.4'. To get recommended systems for precheck to a given catalog version, use '?select=systemsWithUpdatePath&filter=version eq 1.2.3.4'.
// Deprecated: This method is obsolete. Use GetAsSystemSoftwareCatalogsGetResponse instead.
// returns a V1beta1SystemSoftwareCatalogsResponseable when successful
// returns a V1beta1SystemSoftwareCatalogs400Error error when the service returns a 400 status code
// returns a V1beta1SystemSoftwareCatalogs401Error error when the service returns a 401 status code
// returns a V1beta1SystemSoftwareCatalogs404Error error when the service returns a 404 status code
// returns a V1beta1SystemSoftwareCatalogs500Error error when the service returns a 500 status code
func (m *V1beta1SystemSoftwareCatalogsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemSoftwareCatalogsRequestBuilderGetRequestConfiguration)(V1beta1SystemSoftwareCatalogsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemSoftwareCatalogs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemSoftwareCatalogs401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemSoftwareCatalogs404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemSoftwareCatalogs500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemSoftwareCatalogsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemSoftwareCatalogsResponseable), nil
}
// GetAsSystemSoftwareCatalogsGetResponse returns a list of all system software catalogs.Use 'select' and 'filter' query parameters to customize the response returned by this API.For example, to get the End User License Agreement (EULA) for catalog version 1.2.3.4, use '?select=eula&filter=version eq 1.2.3.4'. To get recommended systems for precheck to a given catalog version, use '?select=systemsWithUpdatePath&filter=version eq 1.2.3.4'.
// returns a V1beta1SystemSoftwareCatalogsGetResponseable when successful
// returns a V1beta1SystemSoftwareCatalogs400Error error when the service returns a 400 status code
// returns a V1beta1SystemSoftwareCatalogs401Error error when the service returns a 401 status code
// returns a V1beta1SystemSoftwareCatalogs404Error error when the service returns a 404 status code
// returns a V1beta1SystemSoftwareCatalogs500Error error when the service returns a 500 status code
func (m *V1beta1SystemSoftwareCatalogsRequestBuilder) GetAsSystemSoftwareCatalogsGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemSoftwareCatalogsRequestBuilderGetRequestConfiguration)(V1beta1SystemSoftwareCatalogsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemSoftwareCatalogs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemSoftwareCatalogs401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemSoftwareCatalogs404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemSoftwareCatalogs500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemSoftwareCatalogsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemSoftwareCatalogsGetResponseable), nil
}
// ToGetRequestInformation returns a list of all system software catalogs.Use 'select' and 'filter' query parameters to customize the response returned by this API.For example, to get the End User License Agreement (EULA) for catalog version 1.2.3.4, use '?select=eula&filter=version eq 1.2.3.4'. To get recommended systems for precheck to a given catalog version, use '?select=systemsWithUpdatePath&filter=version eq 1.2.3.4'.
// returns a *RequestInformation when successful
func (m *V1beta1SystemSoftwareCatalogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemSoftwareCatalogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemSoftwareCatalogsRequestBuilder when successful
func (m *V1beta1SystemSoftwareCatalogsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemSoftwareCatalogsRequestBuilder) {
    return NewV1beta1SystemSoftwareCatalogsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
