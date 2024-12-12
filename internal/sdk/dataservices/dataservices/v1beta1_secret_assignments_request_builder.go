package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1SecretAssignmentsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\secret-assignments
type V1beta1SecretAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SecretAssignmentsRequestBuilderGetQueryParameters reports the attributes of the assignments owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
type V1beta1SecretAssignmentsRequestBuilderGetQueryParameters struct {
    // An OData expression to filter responses by attribute. The OData logical operator "eq" is case-sensitive and supported for attributes "applianceId", "secretId", "goal", "service", or "status". The OData function "contains()" is not case-sensitive and supported for attribute "applianceId", "secretId", and "service". The OData logical operator "and" is supported for all attributes.
    Filter *string `uriparametername:"filter"`
    // The limit query parameter should be used in conjunction with offset for paging within a batched result set. The limit is the maximum number of items to include in the response. Example: offset=30&limit=10
    Limit *int32 `uriparametername:"limit"`
    // The offset query parameter should be used in conjunction with limit for paging within a batched result set. The offset is the number of items from the beginning of the batched result set to the first item included in the response. Example: offset=30&limit=10
    Offset *int32 `uriparametername:"offset"`
    // A response attribute to sort by, followed by a direction indicator ("asc" or "desc"). The attribute may be one of "applianceId", "secretId", "createdAt", "goal", "id", "label", "service", "status", or "updatedAt". Default: ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SecretAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SecretAssignmentsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.secretAssignments.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder when successful
func (m *V1beta1SecretAssignmentsRequestBuilder) ById(id string)(*V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.secretAssignments.item collection
// returns a *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder when successful
func (m *V1beta1SecretAssignmentsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SecretAssignmentsRequestBuilderInternal instantiates a new V1beta1SecretAssignmentsRequestBuilder and sets the default values.
func NewV1beta1SecretAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretAssignmentsRequestBuilder) {
    m := &V1beta1SecretAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/secret-assignments{?filter*,limit*,offset*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SecretAssignmentsRequestBuilder instantiates a new V1beta1SecretAssignmentsRequestBuilder and sets the default values.
func NewV1beta1SecretAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SecretAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get reports the attributes of the assignments owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
// Deprecated: This method is obsolete. Use GetAsSecretAssignmentsGetResponse instead.
// returns a V1beta1SecretAssignmentsResponseable when successful
// returns a V1beta1SecretAssignments400Error error when the service returns a 400 status code
// returns a V1beta1SecretAssignments401Error error when the service returns a 401 status code
// returns a V1beta1SecretAssignments403Error error when the service returns a 403 status code
// returns a V1beta1SecretAssignments404Error error when the service returns a 404 status code
// returns a V1beta1SecretAssignments422Error error when the service returns a 422 status code
// returns a V1beta1SecretAssignments500Error error when the service returns a 500 status code
// returns a V1beta1SecretAssignments503Error error when the service returns a 503 status code
func (m *V1beta1SecretAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SecretAssignmentsRequestBuilderGetRequestConfiguration)(V1beta1SecretAssignmentsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretAssignments400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretAssignments401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretAssignments403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretAssignments404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretAssignments422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretAssignments500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretAssignments503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretAssignmentsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretAssignmentsResponseable), nil
}
// GetAsSecretAssignmentsGetResponse reports the attributes of the assignments owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
// returns a V1beta1SecretAssignmentsGetResponseable when successful
// returns a V1beta1SecretAssignments400Error error when the service returns a 400 status code
// returns a V1beta1SecretAssignments401Error error when the service returns a 401 status code
// returns a V1beta1SecretAssignments403Error error when the service returns a 403 status code
// returns a V1beta1SecretAssignments404Error error when the service returns a 404 status code
// returns a V1beta1SecretAssignments422Error error when the service returns a 422 status code
// returns a V1beta1SecretAssignments500Error error when the service returns a 500 status code
// returns a V1beta1SecretAssignments503Error error when the service returns a 503 status code
func (m *V1beta1SecretAssignmentsRequestBuilder) GetAsSecretAssignmentsGetResponse(ctx context.Context, requestConfiguration *V1beta1SecretAssignmentsRequestBuilderGetRequestConfiguration)(V1beta1SecretAssignmentsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretAssignments400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretAssignments401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretAssignments403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretAssignments404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretAssignments422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretAssignments500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretAssignments503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretAssignmentsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretAssignmentsGetResponseable), nil
}
// ToGetRequestInformation reports the attributes of the assignments owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
// returns a *RequestInformation when successful
func (m *V1beta1SecretAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SecretAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SecretAssignmentsRequestBuilder when successful
func (m *V1beta1SecretAssignmentsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SecretAssignmentsRequestBuilder) {
    return NewV1beta1SecretAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
