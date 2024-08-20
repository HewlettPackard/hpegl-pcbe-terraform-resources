package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspVolumesRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-volumes
type V1beta1CspVolumesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspVolumesRequestBuilderGetQueryParameters returns a list of cloud service provider (CSP) volumes according to the given queryparameters for paging, filtering, and sorting.
type V1beta1CspVolumesRequestBuilderGetQueryParameters struct {
    // An expression by which to filter the results.These operators compare a property value to a literal value:* 'eq' - Equal* 'ne' - Not equalA 'contains' function filters by a substring match.  Combine it with 'tolower' for a case insensitivesubstring match.An 'in' operator matches a given value in a collection.  See the example for retrieving the resources ina protection group.Use 'and' and 'or' with parentheses to combine expressions.These fields are available for filtering:* accountInfo/id* cspId* cspInfo/availabilityZone* cspInfo/cspRegion* cspInfo/volumeType* cspType* name* protectionGroupInfo/id* protectionStatus* state* subscriptionInfo/id* resourceGroupInfo/id
    Filter *string `uriparametername:"filter"`
    // An expression containing CSP tag key/value pairs by which to filter the results, by matchingagainst the cspInfo/cspTags property.  An 'eq' operator connects a tag key and value.  Multiple'eq' conditions can be joined with 'or' operators.  An item is included if it has at least onetag key/value pair that matches a tag key/value pair in this expression.
    FilterCspTags *string `uriparametername:"filter%2Dcsp%2Dtags"`
    // The maximum number of items to include in the response.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to omit from the beginning of the result set.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Offset *int32 `uriparametername:"offset"`
    // A resource property by which to sort, followed by an optional directionindicator: "asc" (ascending, the default) or "desc" (descending).These fields are available for sorting:* accountInfo/id* cspId* cspInfo/availabilityZone* cspInfo/createdAt* cspInfo/cspRegion* cspInfo/sizeInGiB* cspInfo/volumeType* id* name* protectionStatus* state
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspVolumesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspVolumesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspVolumesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspVolumes.item collection
// returns a *V1beta1CspVolumesCspVolumesItemRequestBuilder when successful
func (m *V1beta1CspVolumesRequestBuilder) ById(id string)(*V1beta1CspVolumesCspVolumesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1CspVolumesCspVolumesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspVolumesRequestBuilderInternal instantiates a new V1beta1CspVolumesRequestBuilder and sets the default values.
func NewV1beta1CspVolumesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesRequestBuilder) {
    m := &V1beta1CspVolumesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-volumes{?filter*,filter%2Dcsp%2Dtags*,limit*,offset*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspVolumesRequestBuilder instantiates a new V1beta1CspVolumesRequestBuilder and sets the default values.
func NewV1beta1CspVolumesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspVolumesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of cloud service provider (CSP) volumes according to the given queryparameters for paging, filtering, and sorting.
// Deprecated: This method is obsolete. Use GetAsCspVolumesGetResponse instead.
// returns a V1beta1CspVolumesResponseable when successful
// returns a V1beta1CspVolumes400Error error when the service returns a 400 status code
// returns a V1beta1CspVolumes401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumes403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumes500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspVolumesRequestBuilderGetRequestConfiguration)(V1beta1CspVolumesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspVolumes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspVolumes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumes403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesResponseable), nil
}
// GetAsCspVolumesGetResponse returns a list of cloud service provider (CSP) volumes according to the given queryparameters for paging, filtering, and sorting.
// returns a V1beta1CspVolumesGetResponseable when successful
// returns a V1beta1CspVolumes400Error error when the service returns a 400 status code
// returns a V1beta1CspVolumes401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumes403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumes500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesRequestBuilder) GetAsCspVolumesGetResponse(ctx context.Context, requestConfiguration *V1beta1CspVolumesRequestBuilderGetRequestConfiguration)(V1beta1CspVolumesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspVolumes400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspVolumes401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumes403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumes500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesGetResponseable), nil
}
// ToGetRequestInformation returns a list of cloud service provider (CSP) volumes according to the given queryparameters for paging, filtering, and sorting.
// returns a *RequestInformation when successful
func (m *V1beta1CspVolumesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspVolumesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateCspTags the updateCspTags property
// returns a *V1beta1CspVolumesUpdateCspTagsRequestBuilder when successful
func (m *V1beta1CspVolumesRequestBuilder) UpdateCspTags()(*V1beta1CspVolumesUpdateCspTagsRequestBuilder) {
    return NewV1beta1CspVolumesUpdateCspTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspVolumesRequestBuilder when successful
func (m *V1beta1CspVolumesRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspVolumesRequestBuilder) {
    return NewV1beta1CspVolumesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
