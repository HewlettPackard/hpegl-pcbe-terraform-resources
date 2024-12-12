package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1SoftwareReleasesRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-releases
type V1beta1SoftwareReleasesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SoftwareReleasesRequestBuilderGetQueryParameters list multiple Software Releases with filtering, sorting and pagination.Filtering is supported on the `softwareComponent/id` and `version` properties using the`eq`, `in`, `and` and `or` operators. Sorting is supported on the `id` and `version` properties.
type V1beta1SoftwareReleasesRequestBuilderGetQueryParameters struct {
    // An expression to filter list query results. Query result items that match the givenfilter are returned.Expressions must be in the format `<property> <operator> <value>` or`<value> <operator> <property>`. The available operators are:- `eq`: Test whether a property's value is equal to a literal.- `in`: Test whether a property's value appears in a list of literals.Literals can be:- GUIDs, such as `ae09cc99-57e1-4f82-9d80-e68698da641b`.- Strings, such as `'hello'`, `'world'`.Expressions can also be joined using the `and` and `or` logical operators.
    Filter *string `uriparametername:"filter"`
    // The number of query results to return. Use limit in conjuction with offset for paging. 
    Limit *int32 `uriparametername:"limit"`
    // The offset to use for paging through the result set. Use offset in conjunction with limit for paging.
    Offset *int32 `uriparametername:"offset"`
    // Comma separated properties to return in the result. If omitted, all properties will be returned. This is applied to sub-properties of the objects in the items array. Selecting nested properties of an object is not supported.
    Select []string `uriparametername:"select"`
    // One or more properties and directions to sort query results by. A direction is optional and can be either `asc` or `desc` for ascending and descending order respectively. If the direction is omitted it defaults to `asc`.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SoftwareReleasesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SoftwareReleasesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SoftwareReleasesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.softwareReleases.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder when successful
func (m *V1beta1SoftwareReleasesRequestBuilder) ById(id string)(*V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.softwareReleases.item collection
// returns a *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder when successful
func (m *V1beta1SoftwareReleasesRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SoftwareReleasesRequestBuilderInternal instantiates a new V1beta1SoftwareReleasesRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesRequestBuilder) {
    m := &V1beta1SoftwareReleasesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-releases{?filter*,limit*,offset*,select,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareReleasesRequestBuilder instantiates a new V1beta1SoftwareReleasesRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareReleasesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list multiple Software Releases with filtering, sorting and pagination.Filtering is supported on the `softwareComponent/id` and `version` properties using the`eq`, `in`, `and` and `or` operators. Sorting is supported on the `id` and `version` properties.
// Deprecated: This method is obsolete. Use GetAsSoftwareReleasesGetResponse instead.
// returns a V1beta1SoftwareReleasesResponseable when successful
// returns a V1beta1SoftwareReleases400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleases401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleases500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleases503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SoftwareReleasesRequestBuilderGetRequestConfiguration)(V1beta1SoftwareReleasesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleases400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleases401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleases500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleases503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareReleasesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareReleasesResponseable), nil
}
// GetAsSoftwareReleasesGetResponse list multiple Software Releases with filtering, sorting and pagination.Filtering is supported on the `softwareComponent/id` and `version` properties using the`eq`, `in`, `and` and `or` operators. Sorting is supported on the `id` and `version` properties.
// returns a V1beta1SoftwareReleasesGetResponseable when successful
// returns a V1beta1SoftwareReleases400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleases401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleases500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleases503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesRequestBuilder) GetAsSoftwareReleasesGetResponse(ctx context.Context, requestConfiguration *V1beta1SoftwareReleasesRequestBuilderGetRequestConfiguration)(V1beta1SoftwareReleasesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleases400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleases401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleases500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleases503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareReleasesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareReleasesGetResponseable), nil
}
// ToGetRequestInformation list multiple Software Releases with filtering, sorting and pagination.Filtering is supported on the `softwareComponent/id` and `version` properties using the`eq`, `in`, `and` and `or` operators. Sorting is supported on the `id` and `version` properties.
// returns a *RequestInformation when successful
func (m *V1beta1SoftwareReleasesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SoftwareReleasesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SoftwareReleasesRequestBuilder when successful
func (m *V1beta1SoftwareReleasesRequestBuilder) WithUrl(rawUrl string)(*V1beta1SoftwareReleasesRequestBuilder) {
    return NewV1beta1SoftwareReleasesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
