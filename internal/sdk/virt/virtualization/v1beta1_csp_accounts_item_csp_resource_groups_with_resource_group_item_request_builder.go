package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-resource-groups\{resourceGroupId}
type V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) {
    m := &V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-resource-groups/{resourceGroupId}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder instantiates a new V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get details of a resource group in the specified cloud account, where resource group refers toa collection of related assets.
// Deprecated: This method is obsolete. Use GetAsWithResourceGroupGetResponse instead.
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupResponseable when successful
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupResponseable), nil
}
// GetAsWithResourceGroupGetResponse get details of a resource group in the specified cloud account, where resource group refers toa collection of related assets.
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponseable when successful
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) GetAsWithResourceGroupGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroup500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspResourceGroupsItemWithResourceGroupGetResponseable), nil
}
// ToGetRequestInformation get details of a resource group in the specified cloud account, where resource group refers toa collection of related assets.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder) {
    return NewV1beta1CspAccountsItemCspResourceGroupsWithResourceGroupItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
