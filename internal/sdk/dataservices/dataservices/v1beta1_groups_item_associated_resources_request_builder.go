package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1GroupsItemAssociatedResourcesRequestBuilder builds and executes requests for operations under \data-services\v1beta1\groups\{id}\associated-resources
type V1beta1GroupsItemAssociatedResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1GroupsItemAssociatedResourcesRequestBuilderGetQueryParameters returns list of resources associated with a specific group.
type V1beta1GroupsItemAssociatedResourcesRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameterA comparison compares a property name to a literal. The comparisons supported are the following:  - “eq” : Is a property equal to value. Valid for number, boolean and string properties.  - "contains" : Does a property have the text specified as a valueSyntax:   - “eq” : filter=\<property\> eq \<value\>   - "contains" : filter=contains(type, 'vol') User can include 'and' to filter on multiple fields  "filter=type eq volume and contains(name, 'MyVolume')Filters are supported on following attributes:  - type  - name
    Filter *string `uriparametername:"filter"`
    // The number of results to return. Default limit is 12, if not specified by the user.
    Limit *int32 `uriparametername:"limit"`
    // The number of results to skip. This is used for paging results. Default offset  is 0, if notspecified by the user.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response. Currently only support specification of all fields:
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc"). If no direction indicator is specified the default order is ascending. - e.g. "sort=type desc". Currently only supportsorting by 1 property per request. Supported fields include:- type- nameEven though secondary sort is not supported, queries of the database will include a secondary sort by "name asc" to guarantee consistent paging behavior. Of course, there will no secondary sort if user specified "name" as the sorting field.Default sort is "name asc" if not specified by the user.
    Sort *string `uriparametername:"sort"`
}
// V1beta1GroupsItemAssociatedResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1GroupsItemAssociatedResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1GroupsItemAssociatedResourcesRequestBuilderGetQueryParameters
}
// NewV1beta1GroupsItemAssociatedResourcesRequestBuilderInternal instantiates a new V1beta1GroupsItemAssociatedResourcesRequestBuilder and sets the default values.
func NewV1beta1GroupsItemAssociatedResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1GroupsItemAssociatedResourcesRequestBuilder) {
    m := &V1beta1GroupsItemAssociatedResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/groups/{id}/associated-resources{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1GroupsItemAssociatedResourcesRequestBuilder instantiates a new V1beta1GroupsItemAssociatedResourcesRequestBuilder and sets the default values.
func NewV1beta1GroupsItemAssociatedResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1GroupsItemAssociatedResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1GroupsItemAssociatedResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns list of resources associated with a specific group.
// Deprecated: This method is obsolete. Use GetAsAssociatedResourcesGetResponse instead.
// returns a V1beta1GroupsItemAssociatedResourcesResponseable when successful
// returns a V1beta1GroupsItemAssociatedResources400Error error when the service returns a 400 status code
// returns a V1beta1GroupsItemAssociatedResources403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemAssociatedResources404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemAssociatedResources503Error error when the service returns a 503 status code
func (m *V1beta1GroupsItemAssociatedResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1GroupsItemAssociatedResourcesRequestBuilderGetRequestConfiguration)(V1beta1GroupsItemAssociatedResourcesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1GroupsItemAssociatedResources400ErrorFromDiscriminatorValue,
        "403": CreateV1beta1GroupsItemAssociatedResources403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemAssociatedResources404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemAssociatedResources503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsItemAssociatedResourcesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsItemAssociatedResourcesResponseable), nil
}
// GetAsAssociatedResourcesGetResponse returns list of resources associated with a specific group.
// returns a V1beta1GroupsItemAssociatedResourcesGetResponseable when successful
// returns a V1beta1GroupsItemAssociatedResources400Error error when the service returns a 400 status code
// returns a V1beta1GroupsItemAssociatedResources403Error error when the service returns a 403 status code
// returns a V1beta1GroupsItemAssociatedResources404Error error when the service returns a 404 status code
// returns a V1beta1GroupsItemAssociatedResources503Error error when the service returns a 503 status code
func (m *V1beta1GroupsItemAssociatedResourcesRequestBuilder) GetAsAssociatedResourcesGetResponse(ctx context.Context, requestConfiguration *V1beta1GroupsItemAssociatedResourcesRequestBuilderGetRequestConfiguration)(V1beta1GroupsItemAssociatedResourcesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1GroupsItemAssociatedResources400ErrorFromDiscriminatorValue,
        "403": CreateV1beta1GroupsItemAssociatedResources403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1GroupsItemAssociatedResources404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1GroupsItemAssociatedResources503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsItemAssociatedResourcesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsItemAssociatedResourcesGetResponseable), nil
}
// ToGetRequestInformation returns list of resources associated with a specific group.
// returns a *RequestInformation when successful
func (m *V1beta1GroupsItemAssociatedResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1GroupsItemAssociatedResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1GroupsItemAssociatedResourcesRequestBuilder when successful
func (m *V1beta1GroupsItemAssociatedResourcesRequestBuilder) WithUrl(rawUrl string)(*V1beta1GroupsItemAssociatedResourcesRequestBuilder) {
    return NewV1beta1GroupsItemAssociatedResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
