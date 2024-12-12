package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SearchRequestBuilder builds and executes requests for operations under \data-services\v1beta1\search
type V1beta1SearchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SearchRequestBuilderGetQueryParameters returns a list of search results that are visible to the user.The response can be paged by using the `limit` and `offset` query parametersand filtered using the `filter` query parameter.
type V1beta1SearchRequestBuilderGetQueryParameters struct {
    // Indicates whether or not the response will contain aggregated search results.When true:- This is intended to present a summarized list of search results.- The query parameters offset and limit are ignored. Do not set them.- The response contains the top 5 search hits from each of the resource types that the requesting user has permission to view.When false:- This is intended to present detailed, paginated and filtered search results.- The query parameters filter, offset and limit are applied.- The response contains a page of search hits.
    Aggregated *bool `uriparametername:"aggregated"`
    // The filter parameter contains an OData filter expression to furtherfilter the resources returned in addition to the query search string.Filters are supported only on following attributes:- resource/service- resource/typeOnly the following OData filter operators are supported:- in  : The attribute value is one of the values in an array of strings- and : Logical and between multiple filter predicatesThe filter parameter should be url encoded.
    Filter *string `uriparametername:"filter"`
    // Limit is the maximum number of items to include in the response.If not specified, The default limit is 10.Limit should be used in conjunction with `offset`.```offset=30&limit=10```
    Limit *int32 `uriparametername:"limit"`
    // Offset is the number of items from the beginning of the result set tothe first item included in the response.If not specified, the default offset is 0.Offset should be used in conjunction with `limit`.```offset=30&limit=10```
    Offset *int32 `uriparametername:"offset"`
    // Query contains the text to search in DSCC resources.
    Query *string `uriparametername:"query"`
}
// V1beta1SearchRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SearchRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SearchRequestBuilderGetQueryParameters
}
// NewV1beta1SearchRequestBuilderInternal instantiates a new V1beta1SearchRequestBuilder and sets the default values.
func NewV1beta1SearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SearchRequestBuilder) {
    m := &V1beta1SearchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/search?query={query}{&aggregated*,filter*,limit*,offset*}", pathParameters),
    }
    return m
}
// NewV1beta1SearchRequestBuilder instantiates a new V1beta1SearchRequestBuilder and sets the default values.
func NewV1beta1SearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SearchRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of search results that are visible to the user.The response can be paged by using the `limit` and `offset` query parametersand filtered using the `filter` query parameter.
// Deprecated: This method is obsolete. Use GetAsSearchGetResponse instead.
// returns a V1beta1SearchResponseable when successful
// returns a V1beta1Search400Error error when the service returns a 400 status code
// returns a V1beta1Search401Error error when the service returns a 401 status code
// returns a V1beta1Search500Error error when the service returns a 500 status code
func (m *V1beta1SearchRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SearchRequestBuilderGetRequestConfiguration)(V1beta1SearchResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Search400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Search401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Search500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SearchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SearchResponseable), nil
}
// GetAsSearchGetResponse returns a list of search results that are visible to the user.The response can be paged by using the `limit` and `offset` query parametersand filtered using the `filter` query parameter.
// returns a V1beta1SearchGetResponseable when successful
// returns a V1beta1Search400Error error when the service returns a 400 status code
// returns a V1beta1Search401Error error when the service returns a 401 status code
// returns a V1beta1Search500Error error when the service returns a 500 status code
func (m *V1beta1SearchRequestBuilder) GetAsSearchGetResponse(ctx context.Context, requestConfiguration *V1beta1SearchRequestBuilderGetRequestConfiguration)(V1beta1SearchGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Search400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Search401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Search500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SearchGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SearchGetResponseable), nil
}
// ToGetRequestInformation returns a list of search results that are visible to the user.The response can be paged by using the `limit` and `offset` query parametersand filtered using the `filter` query parameter.
// returns a *RequestInformation when successful
func (m *V1beta1SearchRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SearchRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SearchRequestBuilder when successful
func (m *V1beta1SearchRequestBuilder) WithUrl(rawUrl string)(*V1beta1SearchRequestBuilder) {
    return NewV1beta1SearchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
