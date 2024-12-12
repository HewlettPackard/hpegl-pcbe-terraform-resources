package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1EmailSubscriptionsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\email-subscriptions
type V1beta1EmailSubscriptionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1EmailSubscriptionsRequestBuilderGetQueryParameters returns the email subscriptions currently defined in the application.  The "activated" field for eachsubscription instance will be set to 'true' if the user has activated the subscription.  The user is identified by the contents of the bearer token that is included in the request.
type V1beta1EmailSubscriptionsRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of email subscriptions returned in the response.The returned set of email subscriptions must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The comparisons supported are the following:  - “eq” : Is a property equal to value. Valid for number, boolean and string properties.  - "contains" : Does a property have the text specified as a value  Syntax:   - “eq” : filter=\<property\> eq \<value\>  - "contains" : filter=contains(name, 'Houston')User can include "and" to filter on multiple fields   "filter=\<property1\> eq \<value1\> and contains(\<property2\>,\<value2\>)"Filters are supported on the following attributes:  - name  - serviceName  - activated
    Filter *string `uriparametername:"filter"`
    // The number of results to return. Default limit is 12.
    Limit *int32 `uriparametername:"limit"`
    // The number of results to skip. This is used for paging results. Default offset is 0.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response. The service currently only supports specification of all fields.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a direction indicator ("asc" or "desc"). If no direction indicator is specified the default order is ascending. - e.g. "sort=name" is equivalent to "sort=name asc". Service currently only supports sorting by 1 property per request. Supported fields include:- name- serviceName- activatedEven though secondary sort is not supported, queries of the database will include a secondary sort by "name asc" to guarantee consistent paging behavior. Default sort is "name asc" if not specified by the user.
    Sort *string `uriparametername:"sort"`
}
// V1beta1EmailSubscriptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1EmailSubscriptionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1EmailSubscriptionsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.emailSubscriptions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder when successful
func (m *V1beta1EmailSubscriptionsRequestBuilder) ById(id string)(*V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.emailSubscriptions.item collection
// returns a *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder when successful
func (m *V1beta1EmailSubscriptionsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1EmailSubscriptionsRequestBuilderInternal instantiates a new V1beta1EmailSubscriptionsRequestBuilder and sets the default values.
func NewV1beta1EmailSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1EmailSubscriptionsRequestBuilder) {
    m := &V1beta1EmailSubscriptionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/email-subscriptions{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1EmailSubscriptionsRequestBuilder instantiates a new V1beta1EmailSubscriptionsRequestBuilder and sets the default values.
func NewV1beta1EmailSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1EmailSubscriptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1EmailSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the email subscriptions currently defined in the application.  The "activated" field for eachsubscription instance will be set to 'true' if the user has activated the subscription.  The user is identified by the contents of the bearer token that is included in the request.
// Deprecated: This method is obsolete. Use GetAsEmailSubscriptionsGetResponse instead.
// returns a V1beta1EmailSubscriptionsResponseable when successful
// returns a V1beta1EmailSubscriptions400Error error when the service returns a 400 status code
// returns a V1beta1EmailSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1EmailSubscriptions503Error error when the service returns a 503 status code
func (m *V1beta1EmailSubscriptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsRequestBuilderGetRequestConfiguration)(V1beta1EmailSubscriptionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1EmailSubscriptions400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1EmailSubscriptions401ErrorFromDiscriminatorValue,
        "503": CreateV1beta1EmailSubscriptions503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1EmailSubscriptionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1EmailSubscriptionsResponseable), nil
}
// GetAsEmailSubscriptionsGetResponse returns the email subscriptions currently defined in the application.  The "activated" field for eachsubscription instance will be set to 'true' if the user has activated the subscription.  The user is identified by the contents of the bearer token that is included in the request.
// returns a V1beta1EmailSubscriptionsGetResponseable when successful
// returns a V1beta1EmailSubscriptions400Error error when the service returns a 400 status code
// returns a V1beta1EmailSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1EmailSubscriptions503Error error when the service returns a 503 status code
func (m *V1beta1EmailSubscriptionsRequestBuilder) GetAsEmailSubscriptionsGetResponse(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsRequestBuilderGetRequestConfiguration)(V1beta1EmailSubscriptionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1EmailSubscriptions400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1EmailSubscriptions401ErrorFromDiscriminatorValue,
        "503": CreateV1beta1EmailSubscriptions503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1EmailSubscriptionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1EmailSubscriptionsGetResponseable), nil
}
// ToGetRequestInformation returns the email subscriptions currently defined in the application.  The "activated" field for eachsubscription instance will be set to 'true' if the user has activated the subscription.  The user is identified by the contents of the bearer token that is included in the request.
// returns a *RequestInformation when successful
func (m *V1beta1EmailSubscriptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1EmailSubscriptionsRequestBuilder when successful
func (m *V1beta1EmailSubscriptionsRequestBuilder) WithUrl(rawUrl string)(*V1beta1EmailSubscriptionsRequestBuilder) {
    return NewV1beta1EmailSubscriptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
