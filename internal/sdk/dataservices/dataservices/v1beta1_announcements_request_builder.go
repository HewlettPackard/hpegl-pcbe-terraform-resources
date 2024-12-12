package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1AnnouncementsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\announcements
type V1beta1AnnouncementsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1AnnouncementsRequestBuilderGetQueryParameters returns a collection of all published announcements in the system. The following parameters are supported toreduce the collection according to the specified criteria:- filtering announcements by title, content, services and updated timeframe;- sorting announcements by title, created and updated timeframe;- pagination controls (`limit` and `offset`);- selecting specific announcement properties to view.
type V1beta1AnnouncementsRequestBuilderGetQueryParameters struct {
    // The expression to use for filtering responses.The OData filter query parameter supports the following:Operators:- and- or- eq- ge- inProperties:- title- content- services- updatedAtFunctions:- containsFor example:- `contains(title, 'HPE GreenLake') or contains(content, 'HPE GreenLake')`;- `backup-and-recovery in services`;- `updatedAt ge 2023-03-03T07:20:00.00Z and updatedAt lt 2023-03-04T07:20:00.00Z`.        Grouping of expressions to change the evaluation precedence is NOT supported.
    Filter *string `uriparametername:"filter"`
    // The `limit` query parameter should be used in conjunction with `offset` for paging. The`limit` is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // The `offset` query parameter should be used in conjunction with `limit` for paging. The`offset` is the number of items from the beginning of the result set to the first itemincluded in the response.
    Offset *int32 `uriparametername:"offset"`
    // A list of announcement properties to include in the response separated by a `,`,e.g. `id,title`. If this is omitted, all properties shall be returned.
    Select []string `uriparametername:"select"`
    // A list of properties defining the sort order. This takes a single property name followedby the direction to sort in, separated by a space.The supported properties are `title`, `createdAt` and `updatedAt`.If not specified, the default property is `updatedAt`. The supported directions are `asc` and `desc`for ascending and descending respectively. If not specified, the direction defaults to `desc`.Simply specifying `asc` or `desc` without a preceding property is considered invalidand will result in an error being returned.
    Sort *string `uriparametername:"sort"`
}
// V1beta1AnnouncementsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1AnnouncementsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1AnnouncementsRequestBuilderGetQueryParameters
}
// ByAnnouncementId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.announcements.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder when successful
func (m *V1beta1AnnouncementsRequestBuilder) ByAnnouncementId(announcementId string)(*V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if announcementId != "" {
        urlTplParams["announcementId"] = announcementId
    }
    return NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAnnouncementIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.announcements.item collection
// returns a *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder when successful
func (m *V1beta1AnnouncementsRequestBuilder) ByAnnouncementIdGuid(announcementId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["announcementId"] = announcementId.String()
    return NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1AnnouncementsRequestBuilderInternal instantiates a new V1beta1AnnouncementsRequestBuilder and sets the default values.
func NewV1beta1AnnouncementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AnnouncementsRequestBuilder) {
    m := &V1beta1AnnouncementsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/announcements{?filter*,limit*,offset*,select,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1AnnouncementsRequestBuilder instantiates a new V1beta1AnnouncementsRequestBuilder and sets the default values.
func NewV1beta1AnnouncementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AnnouncementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1AnnouncementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a collection of all published announcements in the system. The following parameters are supported toreduce the collection according to the specified criteria:- filtering announcements by title, content, services and updated timeframe;- sorting announcements by title, created and updated timeframe;- pagination controls (`limit` and `offset`);- selecting specific announcement properties to view.
// Deprecated: This method is obsolete. Use GetAsAnnouncementsGetResponse instead.
// returns a V1beta1AnnouncementsResponseable when successful
// returns a V1beta1Announcements400Error error when the service returns a 400 status code
// returns a V1beta1Announcements401Error error when the service returns a 401 status code
// returns a V1beta1Announcements403Error error when the service returns a 403 status code
// returns a V1beta1Announcements500Error error when the service returns a 500 status code
// returns a V1beta1Announcements503Error error when the service returns a 503 status code
func (m *V1beta1AnnouncementsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1AnnouncementsRequestBuilderGetRequestConfiguration)(V1beta1AnnouncementsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Announcements400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Announcements401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Announcements403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Announcements500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Announcements503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AnnouncementsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AnnouncementsResponseable), nil
}
// GetAsAnnouncementsGetResponse returns a collection of all published announcements in the system. The following parameters are supported toreduce the collection according to the specified criteria:- filtering announcements by title, content, services and updated timeframe;- sorting announcements by title, created and updated timeframe;- pagination controls (`limit` and `offset`);- selecting specific announcement properties to view.
// returns a V1beta1AnnouncementsGetResponseable when successful
// returns a V1beta1Announcements400Error error when the service returns a 400 status code
// returns a V1beta1Announcements401Error error when the service returns a 401 status code
// returns a V1beta1Announcements403Error error when the service returns a 403 status code
// returns a V1beta1Announcements500Error error when the service returns a 500 status code
// returns a V1beta1Announcements503Error error when the service returns a 503 status code
func (m *V1beta1AnnouncementsRequestBuilder) GetAsAnnouncementsGetResponse(ctx context.Context, requestConfiguration *V1beta1AnnouncementsRequestBuilderGetRequestConfiguration)(V1beta1AnnouncementsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Announcements400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Announcements401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Announcements403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Announcements500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Announcements503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AnnouncementsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AnnouncementsGetResponseable), nil
}
// ToGetRequestInformation returns a collection of all published announcements in the system. The following parameters are supported toreduce the collection according to the specified criteria:- filtering announcements by title, content, services and updated timeframe;- sorting announcements by title, created and updated timeframe;- pagination controls (`limit` and `offset`);- selecting specific announcement properties to view.
// returns a *RequestInformation when successful
func (m *V1beta1AnnouncementsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1AnnouncementsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1AnnouncementsRequestBuilder when successful
func (m *V1beta1AnnouncementsRequestBuilder) WithUrl(rawUrl string)(*V1beta1AnnouncementsRequestBuilder) {
    return NewV1beta1AnnouncementsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
