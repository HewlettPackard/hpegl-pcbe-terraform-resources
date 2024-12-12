package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1RecentSearchesRequestBuilder builds and executes requests for operations under \data-services\v1beta1\recent-searches
type V1beta1RecentSearchesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1RecentSearchesRequestBuilderGetQueryParameters returns ths DSCC resources that a user selected or viewed through Global Search.Results are sorted with most recent first.
type V1beta1RecentSearchesRequestBuilderGetQueryParameters struct {
    // Limit is the maximum number of items to include in the response.If not specified, The default limit is 5.Limit should be used in conjunction with offset.
    Limit *int32 `uriparametername:"limit"`
}
// V1beta1RecentSearchesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1RecentSearchesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1RecentSearchesRequestBuilderGetQueryParameters
}
// V1beta1RecentSearchesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1RecentSearchesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1RecentSearchesRequestBuilderInternal instantiates a new V1beta1RecentSearchesRequestBuilder and sets the default values.
func NewV1beta1RecentSearchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RecentSearchesRequestBuilder) {
    m := &V1beta1RecentSearchesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/recent-searches{?limit*}", pathParameters),
    }
    return m
}
// NewV1beta1RecentSearchesRequestBuilder instantiates a new V1beta1RecentSearchesRequestBuilder and sets the default values.
func NewV1beta1RecentSearchesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1RecentSearchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1RecentSearchesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns ths DSCC resources that a user selected or viewed through Global Search.Results are sorted with most recent first.
// Deprecated: This method is obsolete. Use GetAsRecentSearchesGetResponse instead.
// returns a V1beta1RecentSearchesResponseable when successful
// returns a V1beta1RecentSearches400Error error when the service returns a 400 status code
// returns a V1beta1RecentSearches401Error error when the service returns a 401 status code
// returns a V1beta1RecentSearches500Error error when the service returns a 500 status code
func (m *V1beta1RecentSearchesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1RecentSearchesRequestBuilderGetRequestConfiguration)(V1beta1RecentSearchesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1RecentSearches400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1RecentSearches401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1RecentSearches500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1RecentSearchesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1RecentSearchesResponseable), nil
}
// GetAsRecentSearchesGetResponse returns ths DSCC resources that a user selected or viewed through Global Search.Results are sorted with most recent first.
// returns a V1beta1RecentSearchesGetResponseable when successful
// returns a V1beta1RecentSearches400Error error when the service returns a 400 status code
// returns a V1beta1RecentSearches401Error error when the service returns a 401 status code
// returns a V1beta1RecentSearches500Error error when the service returns a 500 status code
func (m *V1beta1RecentSearchesRequestBuilder) GetAsRecentSearchesGetResponse(ctx context.Context, requestConfiguration *V1beta1RecentSearchesRequestBuilderGetRequestConfiguration)(V1beta1RecentSearchesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1RecentSearches400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1RecentSearches401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1RecentSearches500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1RecentSearchesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1RecentSearchesGetResponseable), nil
}
// Post saves the details of a resource that a user opened through Search. After performinga search, a user may select one of the search results. This API saves that selectedresource.
// returns a UntypedNodeable when successful
// returns a V1beta1RecentSearches400Error error when the service returns a 400 status code
// returns a V1beta1RecentSearches401Error error when the service returns a 401 status code
// returns a V1beta1RecentSearches500Error error when the service returns a 500 status code
func (m *V1beta1RecentSearchesRequestBuilder) Post(ctx context.Context, body V1beta1RecentSearchesPostRequestBodyable, requestConfiguration *V1beta1RecentSearchesRequestBuilderPostRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1RecentSearches400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1RecentSearches401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1RecentSearches500ErrorFromDiscriminatorValue,
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
// ToGetRequestInformation returns ths DSCC resources that a user selected or viewed through Global Search.Results are sorted with most recent first.
// returns a *RequestInformation when successful
func (m *V1beta1RecentSearchesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1RecentSearchesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation saves the details of a resource that a user opened through Search. After performinga search, a user may select one of the search results. This API saves that selectedresource.
// returns a *RequestInformation when successful
func (m *V1beta1RecentSearchesRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1RecentSearchesPostRequestBodyable, requestConfiguration *V1beta1RecentSearchesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1RecentSearchesRequestBuilder when successful
func (m *V1beta1RecentSearchesRequestBuilder) WithUrl(rawUrl string)(*V1beta1RecentSearchesRequestBuilder) {
    return NewV1beta1RecentSearchesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
