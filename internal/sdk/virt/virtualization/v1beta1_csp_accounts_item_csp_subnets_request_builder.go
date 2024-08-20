package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspSubnetsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-subnets
type V1beta1CspAccountsItemCspSubnetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspSubnetsRequestBuilderGetQueryParameters returns a list of subnets in the specified cloud account.
type V1beta1CspAccountsItemCspSubnetsRequestBuilderGetQueryParameters struct {
    // An expression by which to filter the results.These operators compare a property value to a literal value:  * 'eq' - Equal  * 'ne' - Not equalA 'contains' function filters by a substring match.  Combine it with 'tolower' for a case insensitivesubstring match.Use 'and' and 'or' with parentheses to combine expressions.These fields are available for filtering:* cspId* cspInfo/availabilityZone* cspInfo/cspRegion* name* vpcInfo/id
    Filter *string `uriparametername:"filter"`
    // A resource property by which to sort, followed by an optional directionindicator: "asc" (ascending, the default) or "desc" (descending).These fields are available for sorting:* cspInfo/availabilityZone* cspInfo/cspRegion* id* name
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspAccountsItemCspSubnetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspSubnetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsItemCspSubnetsRequestBuilderGetQueryParameters
}
// BySubnetId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspAccounts.item.cspSubnets.item collection
// returns a *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspSubnetsRequestBuilder) BySubnetId(subnetId string)(*V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if subnetId != "" {
        urlTplParams["subnetId"] = subnetId
    }
    return NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspAccountsItemCspSubnetsRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspSubnetsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubnetsRequestBuilder) {
    m := &V1beta1CspAccountsItemCspSubnetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-subnets{?filter*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspSubnetsRequestBuilder instantiates a new V1beta1CspAccountsItemCspSubnetsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubnetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspSubnetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of subnets in the specified cloud account.
// Deprecated: This method is obsolete. Use GetAsCspSubnetsGetResponse instead.
// returns a V1beta1CspAccountsItemCspSubnetsResponseable when successful
// returns a V1beta1CspAccountsItemCspSubnets400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubnets401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubnets403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubnets404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubnets500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubnetsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubnetsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubnetsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubnets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubnets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubnets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubnets404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubnets500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubnetsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubnetsResponseable), nil
}
// GetAsCspSubnetsGetResponse returns a list of subnets in the specified cloud account.
// returns a V1beta1CspAccountsItemCspSubnetsGetResponseable when successful
// returns a V1beta1CspAccountsItemCspSubnets400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubnets401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubnets403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubnets404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubnets500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubnetsRequestBuilder) GetAsCspSubnetsGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubnetsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubnetsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubnets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubnets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubnets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubnets404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubnets500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubnetsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubnetsGetResponseable), nil
}
// ToGetRequestInformation returns a list of subnets in the specified cloud account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspSubnetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubnetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsItemCspSubnetsRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspSubnetsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspSubnetsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspSubnetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
