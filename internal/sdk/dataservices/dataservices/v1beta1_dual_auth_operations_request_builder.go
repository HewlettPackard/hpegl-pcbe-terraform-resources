package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1DualAuthOperationsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\dual-auth-operations
type V1beta1DualAuthOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DualAuthOperationsRequestBuilderGetQueryParameters returns the list of Dual Authorization operations for the current account. The list will include only the resource types (Application Resource) the user has read permission for. The user must have permission to read pending operations.
type V1beta1DualAuthOperationsRequestBuilderGetQueryParameters struct {
    // The expression to use for filtering responses. The following comparisons are supported:“eq” : Valid for number, boolean and string properties.“gt” :  Valid for number or string timestamp properties.“lt” :  Valid for number or string timestamp properties“in” : Valid for an array of stringsSyntax:“eq” : filter=<property> eq <value> {host:port}/data-services/v1beta1/dual-auth-operations?filter=<property> eq <value>“gt” : filter=<property> gt <value> {host:port}/data-services/v1beta1/dual-auth-operations?filter=<property> gt <value>“lt” : filter=<property> lt <value> {host:port}/data-services/v1beta1/dual-auth-operations?filter=<property> lt <value>“in” : filter=<property> in <value> {host:port}/data-services/v1beta1/dual-auth-operations?filter=<property> in <value>* Use "and" to combine filter inputs {host:port}/data-services/v1beta1/dual-auth-operations?filter=<property1> eq <value1> and <property2> lt <value2>* To filter multiple values on one property e.g. filter=resourceType in ('foo','bar') {host:port}/data-services/v1beta1/dual-auth-operations?filter=foo%bar%20in%20resourceTypeExamples:GET /data-services/v1beta1/dual-auth-operations?filter=resourceType eq 'ISSUE'GET /data-services/v1beta1/dual-auth-operations?filter=resourceType eq 'ISSUE' and state eq 'CREATED'GET /data-services/v1beta1/dual-auth-operations?filter=relatedObjectType in ('NIMBLE-VOLUME')Filters are supported on following attributes:resourceUri,resourceName,resourceType,requestedOperation,operationDescription,requestedByUri,requestedByEmail,requestedAt,customerId,checkedByUri,checkedByEmail,checkedAt,sourceServiceExternalName,state
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging. The limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // Use offset in conjunction with limit for paging. The offset is the number of items from the beginning of the result set to the first item included in the response.
    Offset *int32 `uriparametername:"offset"`
    // Limits the properties returned with a resource or collection-level GET. Specify a comma-separated list of properties. If this is omitted, all properties are returned.
    Select *string `uriparametername:"select"`
    // The property to sort by followed by a direction indicator ("asc" or "desc"). If no direction indicator is specified the default order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1DualAuthOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DualAuthOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1DualAuthOperationsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.dualAuthOperations.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder when successful
func (m *V1beta1DualAuthOperationsRequestBuilder) ById(id string)(*V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.dualAuthOperations.item collection
// returns a *V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder when successful
func (m *V1beta1DualAuthOperationsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1DualAuthOperationsDualAuthOperationsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1DualAuthOperationsRequestBuilderInternal instantiates a new V1beta1DualAuthOperationsRequestBuilder and sets the default values.
func NewV1beta1DualAuthOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DualAuthOperationsRequestBuilder) {
    m := &V1beta1DualAuthOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/dual-auth-operations{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1DualAuthOperationsRequestBuilder instantiates a new V1beta1DualAuthOperationsRequestBuilder and sets the default values.
func NewV1beta1DualAuthOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DualAuthOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DualAuthOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the list of Dual Authorization operations for the current account. The list will include only the resource types (Application Resource) the user has read permission for. The user must have permission to read pending operations.
// Deprecated: This method is obsolete. Use GetAsDualAuthOperationsGetResponse instead.
// returns a V1beta1DualAuthOperationsResponseable when successful
// returns a V1beta1DualAuthOperations400Error error when the service returns a 400 status code
// returns a V1beta1DualAuthOperations401Error error when the service returns a 401 status code
// returns a V1beta1DualAuthOperations403Error error when the service returns a 403 status code
// returns a V1beta1DualAuthOperations404Error error when the service returns a 404 status code
// returns a V1beta1DualAuthOperations422Error error when the service returns a 422 status code
// returns a V1beta1DualAuthOperations500Error error when the service returns a 500 status code
// returns a V1beta1DualAuthOperations503Error error when the service returns a 503 status code
func (m *V1beta1DualAuthOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1DualAuthOperationsRequestBuilderGetRequestConfiguration)(V1beta1DualAuthOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DualAuthOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DualAuthOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DualAuthOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DualAuthOperations404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1DualAuthOperations422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DualAuthOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DualAuthOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DualAuthOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DualAuthOperationsResponseable), nil
}
// GetAsDualAuthOperationsGetResponse returns the list of Dual Authorization operations for the current account. The list will include only the resource types (Application Resource) the user has read permission for. The user must have permission to read pending operations.
// returns a V1beta1DualAuthOperationsGetResponseable when successful
// returns a V1beta1DualAuthOperations400Error error when the service returns a 400 status code
// returns a V1beta1DualAuthOperations401Error error when the service returns a 401 status code
// returns a V1beta1DualAuthOperations403Error error when the service returns a 403 status code
// returns a V1beta1DualAuthOperations404Error error when the service returns a 404 status code
// returns a V1beta1DualAuthOperations422Error error when the service returns a 422 status code
// returns a V1beta1DualAuthOperations500Error error when the service returns a 500 status code
// returns a V1beta1DualAuthOperations503Error error when the service returns a 503 status code
func (m *V1beta1DualAuthOperationsRequestBuilder) GetAsDualAuthOperationsGetResponse(ctx context.Context, requestConfiguration *V1beta1DualAuthOperationsRequestBuilderGetRequestConfiguration)(V1beta1DualAuthOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DualAuthOperations400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DualAuthOperations401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DualAuthOperations403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DualAuthOperations404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1DualAuthOperations422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DualAuthOperations500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DualAuthOperations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DualAuthOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DualAuthOperationsGetResponseable), nil
}
// ToGetRequestInformation returns the list of Dual Authorization operations for the current account. The list will include only the resource types (Application Resource) the user has read permission for. The user must have permission to read pending operations.
// returns a *RequestInformation when successful
func (m *V1beta1DualAuthOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1DualAuthOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1DualAuthOperationsRequestBuilder when successful
func (m *V1beta1DualAuthOperationsRequestBuilder) WithUrl(rawUrl string)(*V1beta1DualAuthOperationsRequestBuilder) {
    return NewV1beta1DualAuthOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
