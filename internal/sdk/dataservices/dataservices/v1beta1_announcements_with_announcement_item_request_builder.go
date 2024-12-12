package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1AnnouncementsWithAnnouncementItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\announcements\{announcementId}
type V1beta1AnnouncementsWithAnnouncementItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1AnnouncementsWithAnnouncementItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1AnnouncementsWithAnnouncementItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilderInternal instantiates a new V1beta1AnnouncementsWithAnnouncementItemRequestBuilder and sets the default values.
func NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) {
    m := &V1beta1AnnouncementsWithAnnouncementItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/announcements/{announcementId}", pathParameters),
    }
    return m
}
// NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilder instantiates a new V1beta1AnnouncementsWithAnnouncementItemRequestBuilder and sets the default values.
func NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns details of a specific announcement with the given ID.
// Deprecated: This method is obsolete. Use GetAsWithAnnouncementGetResponse instead.
// returns a V1beta1AnnouncementsItemWithAnnouncementResponseable when successful
// returns a V1beta1AnnouncementsItemWithAnnouncement400Error error when the service returns a 400 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement401Error error when the service returns a 401 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement403Error error when the service returns a 403 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement404Error error when the service returns a 404 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement500Error error when the service returns a 500 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement503Error error when the service returns a 503 status code
func (m *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1AnnouncementsWithAnnouncementItemRequestBuilderGetRequestConfiguration)(V1beta1AnnouncementsItemWithAnnouncementResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1AnnouncementsItemWithAnnouncement400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1AnnouncementsItemWithAnnouncement401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1AnnouncementsItemWithAnnouncement403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1AnnouncementsItemWithAnnouncement404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1AnnouncementsItemWithAnnouncement500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1AnnouncementsItemWithAnnouncement503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AnnouncementsItemWithAnnouncementResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AnnouncementsItemWithAnnouncementResponseable), nil
}
// GetAsWithAnnouncementGetResponse returns details of a specific announcement with the given ID.
// returns a V1beta1AnnouncementsItemWithAnnouncementGetResponseable when successful
// returns a V1beta1AnnouncementsItemWithAnnouncement400Error error when the service returns a 400 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement401Error error when the service returns a 401 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement403Error error when the service returns a 403 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement404Error error when the service returns a 404 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement500Error error when the service returns a 500 status code
// returns a V1beta1AnnouncementsItemWithAnnouncement503Error error when the service returns a 503 status code
func (m *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) GetAsWithAnnouncementGetResponse(ctx context.Context, requestConfiguration *V1beta1AnnouncementsWithAnnouncementItemRequestBuilderGetRequestConfiguration)(V1beta1AnnouncementsItemWithAnnouncementGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1AnnouncementsItemWithAnnouncement400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1AnnouncementsItemWithAnnouncement401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1AnnouncementsItemWithAnnouncement403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1AnnouncementsItemWithAnnouncement404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1AnnouncementsItemWithAnnouncement500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1AnnouncementsItemWithAnnouncement503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1AnnouncementsItemWithAnnouncementGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1AnnouncementsItemWithAnnouncementGetResponseable), nil
}
// ToGetRequestInformation returns details of a specific announcement with the given ID.
// returns a *RequestInformation when successful
func (m *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1AnnouncementsWithAnnouncementItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder when successful
func (m *V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1AnnouncementsWithAnnouncementItemRequestBuilder) {
    return NewV1beta1AnnouncementsWithAnnouncementItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
