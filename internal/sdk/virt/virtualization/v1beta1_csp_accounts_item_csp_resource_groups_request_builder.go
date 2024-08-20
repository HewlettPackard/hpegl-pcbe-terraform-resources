package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspResourceGroupsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-resource-groups
type V1beta1CspAccountsItemCspResourceGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetQueryParameters returns a list of resource groups in the specified cloud account, where resource group refers toa collection of related assets.For a cloud service provider that doesn't support resource groups, an empty list is returned.
type V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetQueryParameters struct {
    // An expression by which to filter the results.These operators compare a property value to a literal value:* 'eq' - EqualThese fields are available for filtering:* subscriptionInfo/id
    Filter *string `uriparametername:"filter"`
    // A resource property by which to sort, followed by an optional directionindicator: "asc" (ascending, the default) or "desc" (descending).These fields are available for sorting:* cspId* name
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetQueryParameters
}
// ByResourceGroupId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspAccounts.item.cspResourceGroups.item collection
// returns a *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) ByResourceGroupId(resourceGroupId string)(*V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if resourceGroupId != "" {
        urlTplParams["resourceGroupId"] = resourceGroupId
    }
    return NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspResourceGroupsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) {
    m := &V1beta1CspAccountsItemCspResourceGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-resource-groups{?filter*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilder instantiates a new V1beta1CspAccountsItemCspResourceGroupsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of resource groups in the specified cloud account, where resource group refers toa collection of related assets.For a cloud service provider that doesn't support resource groups, an empty list is returned.
// Deprecated: This method is obsolete. Use GetAsCspResourceGroupsGetResponse instead.
// returns a V1beta1CspAccountsItemCspResourceGroupsResponseable when successful
// returns a V1beta1CspAccountsItemCspResourceGroups400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspResourceGroups401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspResourceGroups403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspResourceGroups404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspResourceGroups500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspResourceGroupsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspResourceGroups400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspResourceGroups401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspResourceGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspResourceGroups404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspResourceGroups500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspResourceGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspResourceGroupsResponseable), nil
}
// GetAsCspResourceGroupsGetResponse returns a list of resource groups in the specified cloud account, where resource group refers toa collection of related assets.For a cloud service provider that doesn't support resource groups, an empty list is returned.
// returns a V1beta1CspAccountsItemCspResourceGroupsGetResponseable when successful
// returns a V1beta1CspAccountsItemCspResourceGroups400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspResourceGroups401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspResourceGroups403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspResourceGroups404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspResourceGroups500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) GetAsCspResourceGroupsGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspResourceGroupsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspResourceGroups400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspResourceGroups401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspResourceGroups403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspResourceGroups404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspResourceGroups500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspResourceGroupsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspResourceGroupsGetResponseable), nil
}
// ToGetRequestInformation returns a list of resource groups in the specified cloud account, where resource group refers toa collection of related assets.For a cloud service provider that doesn't support resource groups, an empty list is returned.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspResourceGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
