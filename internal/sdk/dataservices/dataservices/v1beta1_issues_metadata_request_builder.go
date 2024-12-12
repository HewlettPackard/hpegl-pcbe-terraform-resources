package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1IssuesMetadataRequestBuilder builds and executes requests for operations under \data-services\v1beta1\issues-metadata
type V1beta1IssuesMetadataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1IssuesMetadataRequestBuilderGetQueryParameters returns the list of values of category, services, and severity supported by Issues.Functionalities like query parameters, filtering, sorting, grouping, and paging are not supported.
type V1beta1IssuesMetadataRequestBuilderGetQueryParameters struct {
    // The numbers of items to return
    Limit *int32 `uriparametername:"limit"`
    // The number of items to skip before starting to collect the result set
    Offset *int32 `uriparametername:"offset"`
}
// V1beta1IssuesMetadataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1IssuesMetadataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1IssuesMetadataRequestBuilderGetQueryParameters
}
// NewV1beta1IssuesMetadataRequestBuilderInternal instantiates a new V1beta1IssuesMetadataRequestBuilder and sets the default values.
func NewV1beta1IssuesMetadataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1IssuesMetadataRequestBuilder) {
    m := &V1beta1IssuesMetadataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/issues-metadata{?limit*,offset*}", pathParameters),
    }
    return m
}
// NewV1beta1IssuesMetadataRequestBuilder instantiates a new V1beta1IssuesMetadataRequestBuilder and sets the default values.
func NewV1beta1IssuesMetadataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1IssuesMetadataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1IssuesMetadataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the list of values of category, services, and severity supported by Issues.Functionalities like query parameters, filtering, sorting, grouping, and paging are not supported.
// Deprecated: This method is obsolete. Use GetAsIssuesMetadataGetResponse instead.
// returns a V1beta1IssuesMetadataResponseable when successful
// returns a V1beta1IssuesMetadata400Error error when the service returns a 400 status code
// returns a V1beta1IssuesMetadata401Error error when the service returns a 401 status code
// returns a V1beta1IssuesMetadata403Error error when the service returns a 403 status code
// returns a V1beta1IssuesMetadata404Error error when the service returns a 404 status code
// returns a V1beta1IssuesMetadata405Error error when the service returns a 405 status code
// returns a V1beta1IssuesMetadata422Error error when the service returns a 422 status code
// returns a V1beta1IssuesMetadata500Error error when the service returns a 500 status code
// returns a V1beta1IssuesMetadata503Error error when the service returns a 503 status code
func (m *V1beta1IssuesMetadataRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1IssuesMetadataRequestBuilderGetRequestConfiguration)(V1beta1IssuesMetadataResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1IssuesMetadata400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1IssuesMetadata401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1IssuesMetadata403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1IssuesMetadata404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1IssuesMetadata405ErrorFromDiscriminatorValue,
        "422": CreateV1beta1IssuesMetadata422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1IssuesMetadata500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1IssuesMetadata503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1IssuesMetadataResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1IssuesMetadataResponseable), nil
}
// GetAsIssuesMetadataGetResponse returns the list of values of category, services, and severity supported by Issues.Functionalities like query parameters, filtering, sorting, grouping, and paging are not supported.
// returns a V1beta1IssuesMetadataGetResponseable when successful
// returns a V1beta1IssuesMetadata400Error error when the service returns a 400 status code
// returns a V1beta1IssuesMetadata401Error error when the service returns a 401 status code
// returns a V1beta1IssuesMetadata403Error error when the service returns a 403 status code
// returns a V1beta1IssuesMetadata404Error error when the service returns a 404 status code
// returns a V1beta1IssuesMetadata405Error error when the service returns a 405 status code
// returns a V1beta1IssuesMetadata422Error error when the service returns a 422 status code
// returns a V1beta1IssuesMetadata500Error error when the service returns a 500 status code
// returns a V1beta1IssuesMetadata503Error error when the service returns a 503 status code
func (m *V1beta1IssuesMetadataRequestBuilder) GetAsIssuesMetadataGetResponse(ctx context.Context, requestConfiguration *V1beta1IssuesMetadataRequestBuilderGetRequestConfiguration)(V1beta1IssuesMetadataGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1IssuesMetadata400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1IssuesMetadata401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1IssuesMetadata403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1IssuesMetadata404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1IssuesMetadata405ErrorFromDiscriminatorValue,
        "422": CreateV1beta1IssuesMetadata422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1IssuesMetadata500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1IssuesMetadata503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1IssuesMetadataGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1IssuesMetadataGetResponseable), nil
}
// ToGetRequestInformation returns the list of values of category, services, and severity supported by Issues.Functionalities like query parameters, filtering, sorting, grouping, and paging are not supported.
// returns a *RequestInformation when successful
func (m *V1beta1IssuesMetadataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1IssuesMetadataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1IssuesMetadataRequestBuilder when successful
func (m *V1beta1IssuesMetadataRequestBuilder) WithUrl(rawUrl string)(*V1beta1IssuesMetadataRequestBuilder) {
    return NewV1beta1IssuesMetadataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
