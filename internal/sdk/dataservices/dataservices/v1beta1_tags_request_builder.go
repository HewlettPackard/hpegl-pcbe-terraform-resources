package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0b277d2fa239a4fb0c3da10b998ec83f8fc587245ff4515a79a866345755eb44 "github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices/dataservices/v1beta1/tags"
)

// V1beta1TagsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\tags
type V1beta1TagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1TagsRequestBuilderGetQueryParameters returns a list of all tags when no `select` query parameter is provided. When a `select` query parameter is provided, then key or value of the tags are returned.
type V1beta1TagsRequestBuilderGetQueryParameters struct {
    // The set of tags returned in the response.The supported comparisons are:  - “eq” : Valid for number, boolean and string properties.  - “ne” : Valid for number, boolean and string properties.  - "contains"Syntax:  - “eq” : filter=\<property\> eq \<value\>  - “ne” : filter=\<property\> ne \<value\>  - "startswith" : filter=startswith(key, 'Houston')  - "startswith" : filter=startswith(value, 'Houston')You can use "and" to filter on multiple fields   "filter=\<property1\> eq \<value1\> and \<property2\> ne \<value2\>"Examples:  GET /data-services/v1beta1/tags?filter=key eq Houston  GET /data-services/v1beta1/tags?filter=startswith(key, Houston) and value eq VolumeFilters are supported on following attributes:  - key  - value
    Filter *string `uriparametername:"filter"`
    // The number of results to return.
    Limit *int32 `uriparametername:"limit"`
    // The number of results to skip. This is used for paging results.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response. Service currentlyonly supports specification of all fields.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc"). Default order is ascending. Service only supports sorting by 1 property per request. Supported fields include:- key- value - If specified, a secondary sort by "key asc" is included to guarantee consistent paging behavior.
    // Deprecated: This property is deprecated, use SortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc"). Default order is ascending. Service only supports sorting by 1 property per request. Supported fields include:- key- value - If specified, a secondary sort by "key asc" is included to guarantee consistent paging behavior.
    SortAsGetSortQueryParameterType *i0b277d2fa239a4fb0c3da10b998ec83f8fc587245ff4515a79a866345755eb44.GetSortQueryParameterType `uriparametername:"sort"`
}
// V1beta1TagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1TagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1TagsRequestBuilderGetQueryParameters
}
// NewV1beta1TagsRequestBuilderInternal instantiates a new V1beta1TagsRequestBuilder and sets the default values.
func NewV1beta1TagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1TagsRequestBuilder) {
    m := &V1beta1TagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/tags{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1TagsRequestBuilder instantiates a new V1beta1TagsRequestBuilder and sets the default values.
func NewV1beta1TagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1TagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1TagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of all tags when no `select` query parameter is provided. When a `select` query parameter is provided, then key or value of the tags are returned.
// Deprecated: This method is obsolete. Use GetAsTagsGetResponse instead.
// returns a V1beta1TagsResponseable when successful
// returns a V1beta1Tags400Error error when the service returns a 400 status code
// returns a V1beta1Tags401Error error when the service returns a 401 status code
// returns a V1beta1Tags403Error error when the service returns a 403 status code
// returns a V1beta1Tags404Error error when the service returns a 404 status code
// returns a V1beta1Tags500Error error when the service returns a 500 status code
// returns a V1beta1Tags503Error error when the service returns a 503 status code
func (m *V1beta1TagsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1TagsRequestBuilderGetRequestConfiguration)(V1beta1TagsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Tags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Tags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Tags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Tags404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Tags500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Tags503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1TagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1TagsResponseable), nil
}
// GetAsTagsGetResponse returns a list of all tags when no `select` query parameter is provided. When a `select` query parameter is provided, then key or value of the tags are returned.
// returns a V1beta1TagsGetResponseable when successful
// returns a V1beta1Tags400Error error when the service returns a 400 status code
// returns a V1beta1Tags401Error error when the service returns a 401 status code
// returns a V1beta1Tags403Error error when the service returns a 403 status code
// returns a V1beta1Tags404Error error when the service returns a 404 status code
// returns a V1beta1Tags500Error error when the service returns a 500 status code
// returns a V1beta1Tags503Error error when the service returns a 503 status code
func (m *V1beta1TagsRequestBuilder) GetAsTagsGetResponse(ctx context.Context, requestConfiguration *V1beta1TagsRequestBuilderGetRequestConfiguration)(V1beta1TagsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Tags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Tags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Tags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Tags404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Tags500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Tags503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1TagsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1TagsGetResponseable), nil
}
// ToGetRequestInformation returns a list of all tags when no `select` query parameter is provided. When a `select` query parameter is provided, then key or value of the tags are returned.
// returns a *RequestInformation when successful
func (m *V1beta1TagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1TagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1TagsRequestBuilder when successful
func (m *V1beta1TagsRequestBuilder) WithUrl(rawUrl string)(*V1beta1TagsRequestBuilder) {
    return NewV1beta1TagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
