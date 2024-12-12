package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1IssuesIssuesItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\issues\{id}
type V1beta1IssuesIssuesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1IssuesIssuesItemRequestBuilderGetQueryParameters returns an active issue with the specified id for the account obtained from the request header. The issue must be in the CREATED state
type V1beta1IssuesIssuesItemRequestBuilderGetQueryParameters struct {
    // Limits the properties returned with a resource or collection-level GET. Specify a comma-separated list of properties.(e.g.: "?select=id,type,customerId,services,createdAt,lastOccurredAt,generation,resourceUri")
    Select *string `uriparametername:"select"`
}
// V1beta1IssuesIssuesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1IssuesIssuesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1IssuesIssuesItemRequestBuilderGetQueryParameters
}
// NewV1beta1IssuesIssuesItemRequestBuilderInternal instantiates a new V1beta1IssuesIssuesItemRequestBuilder and sets the default values.
func NewV1beta1IssuesIssuesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1IssuesIssuesItemRequestBuilder) {
    m := &V1beta1IssuesIssuesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/issues/{id}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1IssuesIssuesItemRequestBuilder instantiates a new V1beta1IssuesIssuesItemRequestBuilder and sets the default values.
func NewV1beta1IssuesIssuesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1IssuesIssuesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1IssuesIssuesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns an active issue with the specified id for the account obtained from the request header. The issue must be in the CREATED state
// Deprecated: This method is obsolete. Use GetAsIssuesGetResponse instead.
// returns a V1beta1IssuesItemIssuesResponseable when successful
// returns a V1beta1IssuesItemIssues400Error error when the service returns a 400 status code
// returns a V1beta1IssuesItemIssues401Error error when the service returns a 401 status code
// returns a V1beta1IssuesItemIssues403Error error when the service returns a 403 status code
// returns a V1beta1IssuesItemIssues404Error error when the service returns a 404 status code
// returns a V1beta1IssuesItemIssues405Error error when the service returns a 405 status code
// returns a V1beta1IssuesItemIssues409Error error when the service returns a 409 status code
// returns a V1beta1IssuesItemIssues422Error error when the service returns a 422 status code
// returns a V1beta1IssuesItemIssues500Error error when the service returns a 500 status code
// returns a V1beta1IssuesItemIssues503Error error when the service returns a 503 status code
func (m *V1beta1IssuesIssuesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1IssuesIssuesItemRequestBuilderGetRequestConfiguration)(V1beta1IssuesItemIssuesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1IssuesItemIssues400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1IssuesItemIssues401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1IssuesItemIssues403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1IssuesItemIssues404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1IssuesItemIssues405ErrorFromDiscriminatorValue,
        "409": CreateV1beta1IssuesItemIssues409ErrorFromDiscriminatorValue,
        "422": CreateV1beta1IssuesItemIssues422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1IssuesItemIssues500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1IssuesItemIssues503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1IssuesItemIssuesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1IssuesItemIssuesResponseable), nil
}
// GetAsIssuesGetResponse returns an active issue with the specified id for the account obtained from the request header. The issue must be in the CREATED state
// returns a V1beta1IssuesItemIssuesGetResponseable when successful
// returns a V1beta1IssuesItemIssues400Error error when the service returns a 400 status code
// returns a V1beta1IssuesItemIssues401Error error when the service returns a 401 status code
// returns a V1beta1IssuesItemIssues403Error error when the service returns a 403 status code
// returns a V1beta1IssuesItemIssues404Error error when the service returns a 404 status code
// returns a V1beta1IssuesItemIssues405Error error when the service returns a 405 status code
// returns a V1beta1IssuesItemIssues409Error error when the service returns a 409 status code
// returns a V1beta1IssuesItemIssues422Error error when the service returns a 422 status code
// returns a V1beta1IssuesItemIssues500Error error when the service returns a 500 status code
// returns a V1beta1IssuesItemIssues503Error error when the service returns a 503 status code
func (m *V1beta1IssuesIssuesItemRequestBuilder) GetAsIssuesGetResponse(ctx context.Context, requestConfiguration *V1beta1IssuesIssuesItemRequestBuilderGetRequestConfiguration)(V1beta1IssuesItemIssuesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1IssuesItemIssues400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1IssuesItemIssues401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1IssuesItemIssues403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1IssuesItemIssues404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1IssuesItemIssues405ErrorFromDiscriminatorValue,
        "409": CreateV1beta1IssuesItemIssues409ErrorFromDiscriminatorValue,
        "422": CreateV1beta1IssuesItemIssues422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1IssuesItemIssues500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1IssuesItemIssues503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1IssuesItemIssuesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1IssuesItemIssuesGetResponseable), nil
}
// ToGetRequestInformation returns an active issue with the specified id for the account obtained from the request header. The issue must be in the CREATED state
// returns a *RequestInformation when successful
func (m *V1beta1IssuesIssuesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1IssuesIssuesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1IssuesIssuesItemRequestBuilder when successful
func (m *V1beta1IssuesIssuesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1IssuesIssuesItemRequestBuilder) {
    return NewV1beta1IssuesIssuesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
