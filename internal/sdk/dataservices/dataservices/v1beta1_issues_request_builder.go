package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1IssuesRequestBuilder builds and executes requests for operations under \data-services\v1beta1\issues
type V1beta1IssuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1IssuesRequestBuilderGetQueryParameters returns the list of active issues for the account and user obtained from the request header. The list includesissues only for the resource-types that the user has permissions to view. Active issues are issues in the CREATED state. Clients using this API must process the returned issues for any desired groupings.
type V1beta1IssuesRequestBuilderGetQueryParameters struct {
    // The expression used to filter responses. You can filter on the following properties: `issueType`, `severity`, `category`, `state`, `createdAt`, `services`, `sourceResourceId`, `sourceResourceType`. You can combine multiple comparison operators using “and”.The returned set of resources must match the criteria in the filter query parameterA comparison compares a property name to a literal. The comparisons supported are the following:“eq” : Is a property equal to value. Valid for number, boolean and string properties.“gt” : Is a property greater than a value. Valid for number or string timestamp properties.“lt” : Is a property less than a value. Valid for number or string timestamp properties“in” : Is a value in a property. The property is an array of number, boolean or string properties."contains": Is a substring value that is equal to a portion of the property value. Valid for strings.Syntax: “eq” : filter=<property> eq <value> {host:port}/data-services/v1beta1/issues?filter=<property> eq <value>“gt” : filter=<property> gt <value> {host:port}/data-services/v1beta1/issues?filter=<property> gt <value>“lt” : filter=<property> lt <value> {host:port}/data-services/v1beta1/issues?filter=<property> lt <value>“in” : filter=<property> in <value> {host:port}/data-services/v1beta1/issues?filter=<property> in <value>“contains” : filter=contains(property,value) {host:port}/data-services/v1beta1/issues?filter=contains(property,value)* Can use and to add more filter inputs {host:port}/data-services/v1beta1/issues?filter=<property1> eq <value1> and <property2> lt <value2> * To filter multiple values on one property e.g. filter=severity in ('CRITICAL','WARNING') {host:port}/data-services/v1beta1/issues?filter=severity%20in%20CRITICAL%2CWARNINGExamples:GET /data-services/v1beta1/issues?filter=issueType eq 'ISSUE'GET /data-services/v1beta1/issues?filter=issueType eq 'ISSUE' & state eq 'CREATED'GET /data-services/v1beta1/issues?filter=contains(sourceResourceType,'orchestrator')GET /data-services/v1beta1/issues?filter='data-ops-manager' in servicesFilters are supported on following attributes:issueType,severity,category,state,createdAt,services,sourceResourceId,sourceResourceType
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging. The limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // Use offset in conjunction with limit for paging. The offset is the number of items from the beginningof the result set to the first item included in the response.
    Offset *int32 `uriparametername:"offset"`
    // The select query parameter is used to limit the properties returned with a resource or collection-level GET.Multiple properties can be listed to be returned. The server must only return the set of properties requested by the client.The property “select” is the name of the select query parameter; its value is the list of properties to return separated by commas.
    Select *string `uriparametername:"select"`
    // resource property to sort, with an order appendedOrder may only be either “asc” (ascending) or “desc” (descending)
    Sort *string `uriparametername:"sort"`
}
// V1beta1IssuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1IssuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1IssuesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.issues.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1IssuesIssuesItemRequestBuilder when successful
func (m *V1beta1IssuesRequestBuilder) ById(id string)(*V1beta1IssuesIssuesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1IssuesIssuesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.issues.item collection
// returns a *V1beta1IssuesIssuesItemRequestBuilder when successful
func (m *V1beta1IssuesRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1IssuesIssuesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1IssuesIssuesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1IssuesRequestBuilderInternal instantiates a new V1beta1IssuesRequestBuilder and sets the default values.
func NewV1beta1IssuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1IssuesRequestBuilder) {
    m := &V1beta1IssuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/issues{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1IssuesRequestBuilder instantiates a new V1beta1IssuesRequestBuilder and sets the default values.
func NewV1beta1IssuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1IssuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1IssuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the list of active issues for the account and user obtained from the request header. The list includesissues only for the resource-types that the user has permissions to view. Active issues are issues in the CREATED state. Clients using this API must process the returned issues for any desired groupings.
// Deprecated: This method is obsolete. Use GetAsIssuesGetResponse instead.
// returns a V1beta1IssuesResponseable when successful
// returns a V1beta1Issues400Error error when the service returns a 400 status code
// returns a V1beta1Issues401Error error when the service returns a 401 status code
// returns a V1beta1Issues403Error error when the service returns a 403 status code
// returns a V1beta1Issues404Error error when the service returns a 404 status code
// returns a V1beta1Issues405Error error when the service returns a 405 status code
// returns a V1beta1Issues422Error error when the service returns a 422 status code
// returns a V1beta1Issues500Error error when the service returns a 500 status code
// returns a V1beta1Issues503Error error when the service returns a 503 status code
func (m *V1beta1IssuesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1IssuesRequestBuilderGetRequestConfiguration)(V1beta1IssuesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Issues400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Issues401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Issues403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Issues404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1Issues405ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Issues422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Issues500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Issues503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1IssuesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1IssuesResponseable), nil
}
// GetAsIssuesGetResponse returns the list of active issues for the account and user obtained from the request header. The list includesissues only for the resource-types that the user has permissions to view. Active issues are issues in the CREATED state. Clients using this API must process the returned issues for any desired groupings.
// returns a V1beta1IssuesGetResponseable when successful
// returns a V1beta1Issues400Error error when the service returns a 400 status code
// returns a V1beta1Issues401Error error when the service returns a 401 status code
// returns a V1beta1Issues403Error error when the service returns a 403 status code
// returns a V1beta1Issues404Error error when the service returns a 404 status code
// returns a V1beta1Issues405Error error when the service returns a 405 status code
// returns a V1beta1Issues422Error error when the service returns a 422 status code
// returns a V1beta1Issues500Error error when the service returns a 500 status code
// returns a V1beta1Issues503Error error when the service returns a 503 status code
func (m *V1beta1IssuesRequestBuilder) GetAsIssuesGetResponse(ctx context.Context, requestConfiguration *V1beta1IssuesRequestBuilderGetRequestConfiguration)(V1beta1IssuesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Issues400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Issues401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Issues403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Issues404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1Issues405ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Issues422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Issues500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Issues503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1IssuesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1IssuesGetResponseable), nil
}
// ToGetRequestInformation returns the list of active issues for the account and user obtained from the request header. The list includesissues only for the resource-types that the user has permissions to view. Active issues are issues in the CREATED state. Clients using this API must process the returned issues for any desired groupings.
// returns a *RequestInformation when successful
func (m *V1beta1IssuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1IssuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1IssuesRequestBuilder when successful
func (m *V1beta1IssuesRequestBuilder) WithUrl(rawUrl string)(*V1beta1IssuesRequestBuilder) {
    return NewV1beta1IssuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
