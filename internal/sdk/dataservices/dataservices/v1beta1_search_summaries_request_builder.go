package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SearchSummariesRequestBuilder builds and executes requests for operations under \data-services\v1beta1\search-summaries
type V1beta1SearchSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SearchSummariesRequestBuilderGetQueryParameters returns a list of services that have search hits and the count of hits in each service.The response can be filtered using the `filter` query parameter.
type V1beta1SearchSummariesRequestBuilderGetQueryParameters struct {
    // The filter parameter contains an OData filter expression to furtherfilter the resources returned in addition to the query search string.Filters are supported only on following attributes:- nameOnly the following OData filter operators are supported:- in  : The attribute value is one of the values in an array of stringsThe filter parameter should be url encoded.
    Filter *string `uriparametername:"filter"`
    // Query contains the text to search in DSCC resources.
    Query *string `uriparametername:"query"`
}
// V1beta1SearchSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SearchSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SearchSummariesRequestBuilderGetQueryParameters
}
// NewV1beta1SearchSummariesRequestBuilderInternal instantiates a new V1beta1SearchSummariesRequestBuilder and sets the default values.
func NewV1beta1SearchSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SearchSummariesRequestBuilder) {
    m := &V1beta1SearchSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/search-summaries?query={query}{&filter*}", pathParameters),
    }
    return m
}
// NewV1beta1SearchSummariesRequestBuilder instantiates a new V1beta1SearchSummariesRequestBuilder and sets the default values.
func NewV1beta1SearchSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SearchSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SearchSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of services that have search hits and the count of hits in each service.The response can be filtered using the `filter` query parameter.
// Deprecated: This method is obsolete. Use GetAsSearchSummariesGetResponse instead.
// returns a V1beta1SearchSummariesResponseable when successful
// returns a V1beta1SearchSummaries400Error error when the service returns a 400 status code
// returns a V1beta1SearchSummaries401Error error when the service returns a 401 status code
// returns a V1beta1SearchSummaries500Error error when the service returns a 500 status code
func (m *V1beta1SearchSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SearchSummariesRequestBuilderGetRequestConfiguration)(V1beta1SearchSummariesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SearchSummaries400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SearchSummaries401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SearchSummaries500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SearchSummariesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SearchSummariesResponseable), nil
}
// GetAsSearchSummariesGetResponse returns a list of services that have search hits and the count of hits in each service.The response can be filtered using the `filter` query parameter.
// returns a V1beta1SearchSummariesGetResponseable when successful
// returns a V1beta1SearchSummaries400Error error when the service returns a 400 status code
// returns a V1beta1SearchSummaries401Error error when the service returns a 401 status code
// returns a V1beta1SearchSummaries500Error error when the service returns a 500 status code
func (m *V1beta1SearchSummariesRequestBuilder) GetAsSearchSummariesGetResponse(ctx context.Context, requestConfiguration *V1beta1SearchSummariesRequestBuilderGetRequestConfiguration)(V1beta1SearchSummariesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SearchSummaries400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SearchSummaries401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SearchSummaries500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SearchSummariesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SearchSummariesGetResponseable), nil
}
// ToGetRequestInformation returns a list of services that have search hits and the count of hits in each service.The response can be filtered using the `filter` query parameter.
// returns a *RequestInformation when successful
func (m *V1beta1SearchSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SearchSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SearchSummariesRequestBuilder when successful
func (m *V1beta1SearchSummariesRequestBuilder) WithUrl(rawUrl string)(*V1beta1SearchSummariesRequestBuilder) {
    return NewV1beta1SearchSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
