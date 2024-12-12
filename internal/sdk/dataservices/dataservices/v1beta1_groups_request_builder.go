package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1GroupsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\groups
type V1beta1GroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1GroupsRequestBuilderGetQueryParameters returns all groups for a customer
type V1beta1GroupsRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of groups returned in the response.The returned set of groups must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The comparisons supported are the following:  - “eq” : Is a property equal to value. Valid for number, boolean and string properties.  - “gt” : Is a property greater than a value. Valid for number or string timestamp properties.  - “lt” : Is a property less than a value. Valid for number or string timestamp properties  - "contains" : Does a property have the text specified as a valueSyntax:   - “eq” : filter=\<property\> eq \<value\>  - “gt” : filter=\<property\> gt \<value\>  - “lt” : filter=\<property\> lt \<value\>  - "contains" : filter=contains(name, 'Houston')  - "contains" : filter=contains(description, 'Houston')User can include "and" to filter on multiple fields   "filter=\<property1\> eq \<value1\> and \<property2\> lt \<value2\>"Filters are supported on following attributes:  - name  - updatedBy  - updatedAt
    Filter *string `uriparametername:"filter"`
    // The number of results to return. Default limit is 12, if not specified by the user.
    Limit *int32 `uriparametername:"limit"`
    // The number of results to skip. This is used for paging results. Default offset  is 0, if notspecified by the user.
    Offset *int32 `uriparametername:"offset"`
    // List of permissions to check for group access, each of which has the form "resource-type.permission" If omitted, default permissions group.create and group.read will be used.
    Permissions []string `uriparametername:"permissions"`
    // A list of properties to include in the response. Service currently only supports specification of all fields.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc"). If no direction indicator is specified the default order is ascending. - e.g. "sort=name" is equivalent to "sort=name asc". Service currently only supports sorting by 1 property per request. Supported fields include:- name- updatedBy- updatedAt- description- associationCountEven though secondary sort is not supported, queries of the database will include a secondary sort by "name asc" to guarantee consistent paging behavior. Of course, there will no secondary sort if user specified "name" as the sorting field.Default sort is "name asc" if not specified by the user.
    Sort *string `uriparametername:"sort"`
}
// V1beta1GroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1GroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1GroupsRequestBuilderGetQueryParameters
}
// V1beta1GroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1GroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.groups.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1GroupsGroupsItemRequestBuilder when successful
func (m *V1beta1GroupsRequestBuilder) ById(id string)(*V1beta1GroupsGroupsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1GroupsGroupsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.groups.item collection
// returns a *V1beta1GroupsGroupsItemRequestBuilder when successful
func (m *V1beta1GroupsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1GroupsGroupsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1GroupsGroupsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1GroupsRequestBuilderInternal instantiates a new V1beta1GroupsRequestBuilder and sets the default values.
func NewV1beta1GroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1GroupsRequestBuilder) {
    m := &V1beta1GroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/groups{?filter*,limit*,offset*,permissions*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1GroupsRequestBuilder instantiates a new V1beta1GroupsRequestBuilder and sets the default values.
func NewV1beta1GroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1GroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1GroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all groups for a customer
// Deprecated: This method is obsolete. Use GetAsGroupsGetResponse instead.
// returns a V1beta1GroupsResponseable when successful
// returns a V1beta1Groups400Error error when the service returns a 400 status code
// returns a V1beta1Groups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1GroupsRequestBuilderGetRequestConfiguration)(V1beta1GroupsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Groups400ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Groups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsResponseable), nil
}
// GetAsGroupsGetResponse returns all groups for a customer
// returns a V1beta1GroupsGetResponseable when successful
// returns a V1beta1Groups400Error error when the service returns a 400 status code
// returns a V1beta1Groups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsRequestBuilder) GetAsGroupsGetResponse(ctx context.Context, requestConfiguration *V1beta1GroupsRequestBuilderGetRequestConfiguration)(V1beta1GroupsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Groups400ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Groups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsGetResponseable), nil
}
// Post create a new group
// Deprecated: This method is obsolete. Use PostAsGroupsPostResponse instead.
// returns a V1beta1GroupsResponseable when successful
// returns a V1beta1Groups400Error error when the service returns a 400 status code
// returns a V1beta1Groups403Error error when the service returns a 403 status code
// returns a V1beta1Groups409Error error when the service returns a 409 status code
// returns a V1beta1Groups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsRequestBuilder) Post(ctx context.Context, body V1beta1GroupsPostRequestBodyable, requestConfiguration *V1beta1GroupsRequestBuilderPostRequestConfiguration)(V1beta1GroupsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Groups400ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Groups403ErrorFromDiscriminatorValue,
        "409": CreateV1beta1Groups409ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Groups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsResponseable), nil
}
// PostAsGroupsPostResponse create a new group
// returns a V1beta1GroupsPostResponseable when successful
// returns a V1beta1Groups400Error error when the service returns a 400 status code
// returns a V1beta1Groups403Error error when the service returns a 403 status code
// returns a V1beta1Groups409Error error when the service returns a 409 status code
// returns a V1beta1Groups503Error error when the service returns a 503 status code
func (m *V1beta1GroupsRequestBuilder) PostAsGroupsPostResponse(ctx context.Context, body V1beta1GroupsPostRequestBodyable, requestConfiguration *V1beta1GroupsRequestBuilderPostRequestConfiguration)(V1beta1GroupsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Groups400ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Groups403ErrorFromDiscriminatorValue,
        "409": CreateV1beta1Groups409ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Groups503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1GroupsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1GroupsPostResponseable), nil
}
// ToGetRequestInformation returns all groups for a customer
// returns a *RequestInformation when successful
func (m *V1beta1GroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1GroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new group
// returns a *RequestInformation when successful
func (m *V1beta1GroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1GroupsPostRequestBodyable, requestConfiguration *V1beta1GroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1GroupsRequestBuilder when successful
func (m *V1beta1GroupsRequestBuilder) WithUrl(rawUrl string)(*V1beta1GroupsRequestBuilder) {
    return NewV1beta1GroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
