package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstancesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instances
type V1beta1CspMachineInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstancesRequestBuilderGetQueryParameters returns a list of cloud service provider (CSP) machine instances according to the given queryparameters for paging, filtering, and sorting.
type V1beta1CspMachineInstancesRequestBuilderGetQueryParameters struct {
    // An expression by which to filter the results.These operators compare a property value to a literal value:* 'eq' - Equal.* 'ne' - Not equal.A 'contains' function filters by a substring match.  Combine it with 'tolower' for a case insensitivesubstring match.An 'in' operator matches a given value in a collection.  See the example for retrieving the resources ina protection group.Use 'and' and 'or' with parentheses to combine expressions.These fields are available for filtering:* accountInfo/id* cspId* cspInfo/availabilityZone* cspInfo/cspRegion* cspInfo/instanceType* cspInfo/networkInfo/subnetInfo/id* cspInfo/networkInfo/vpcInfo/id* cspInfo/networkInfo/publicIpAddress* cspInfo/networkInfo/privateIpAddress* cspInfo/state* cspType* name* protectionGroupInfo/id* protectionStatus* state* subscriptionInfo/id* resourceGroupInfo/id
    Filter *string `uriparametername:"filter"`
    // An expression containing CSP tag key/value pairs by which to filter the results, by matchingagainst the cspInfo/cspTags property.  An 'eq' operator connects a tag key and value.  Multiple'eq' conditions can be joined with 'or' operators.  An item is included if it has at least onetag key/value pair that matches a tag key/value pair in this expression.
    FilterCspTags *string `uriparametername:"filter%2Dcsp%2Dtags"`
    // The maximum number of items to include in the response.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to omit from the beginning of the result set.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Offset *int32 `uriparametername:"offset"`
    // A resource property by which to sort, followed by an optional directionindicator: "asc" (ascending, the default) or "desc" (descending).These fields are available for sorting:* accountInfo/id* cspId* cspInfo/availabilityZone* cspInfo/createdAt* cspInfo/cspRegion* cspInfo/instanceType* cspInfo/networkInfo/privateIpAddress* cspInfo/networkInfo/publicIpAddress* cspInfo/networkInfo/subnetInfo/id* cspInfo/networkInfo/vpcInfo/id* cspInfo/state* id* name* protectionStatus* state
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspMachineInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspMachineInstancesRequestBuilderGetQueryParameters
}
// V1beta1CspMachineInstancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspMachineInstances.item collection
// returns a *V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder when successful
func (m *V1beta1CspMachineInstancesRequestBuilder) ById(id string)(*V1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1CspMachineInstancesCspMachineInstancesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspMachineInstancesRequestBuilderInternal instantiates a new V1beta1CspMachineInstancesRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesRequestBuilder) {
    m := &V1beta1CspMachineInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instances{?filter*,filter%2Dcsp%2Dtags*,limit*,offset*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstancesRequestBuilder instantiates a new V1beta1CspMachineInstancesRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of cloud service provider (CSP) machine instances according to the given queryparameters for paging, filtering, and sorting.
// Deprecated: This method is obsolete. Use GetAsCspMachineInstancesGetResponse instead.
// returns a V1beta1CspMachineInstancesResponseable when successful
// returns a V1beta1CspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstancesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstances403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesResponseable), nil
}
// GetAsCspMachineInstancesGetResponse returns a list of cloud service provider (CSP) machine instances according to the given queryparameters for paging, filtering, and sorting.
// returns a V1beta1CspMachineInstancesGetResponseable when successful
// returns a V1beta1CspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesRequestBuilder) GetAsCspMachineInstancesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesRequestBuilderGetRequestConfiguration)(V1beta1CspMachineInstancesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstances403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesGetResponseable), nil
}
// Post create CSP Machine Instance
// Deprecated: This method is obsolete. Use PostAsCspMachineInstancesPostResponse instead.
// returns a V1beta1CspMachineInstancesResponseable when successful
// returns a V1beta1CspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstances409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesRequestBuilder) Post(ctx context.Context, body V1beta1CspMachineInstancesPostRequestBodyable, requestConfiguration *V1beta1CspMachineInstancesRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstances404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstances409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesResponseable), nil
}
// PostAsCspMachineInstancesPostResponse create CSP Machine Instance
// returns a V1beta1CspMachineInstancesPostResponseable when successful
// returns a V1beta1CspMachineInstances400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstances401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstances403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstances404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstances409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstances500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesRequestBuilder) PostAsCspMachineInstancesPostResponse(ctx context.Context, body V1beta1CspMachineInstancesPostRequestBodyable, requestConfiguration *V1beta1CspMachineInstancesRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstances400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstances401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstances403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstances404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstances409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstances500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesPostResponseable), nil
}
// ToGetRequestInformation returns a list of cloud service provider (CSP) machine instances according to the given queryparameters for paging, filtering, and sorting.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create CSP Machine Instance
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1CspMachineInstancesPostRequestBodyable, requestConfiguration *V1beta1CspMachineInstancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateCspTags the updateCspTags property
// returns a *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder when successful
func (m *V1beta1CspMachineInstancesRequestBuilder) UpdateCspTags()(*V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) {
    return NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstancesRequestBuilder when successful
func (m *V1beta1CspMachineInstancesRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstancesRequestBuilder) {
    return NewV1beta1CspMachineInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
