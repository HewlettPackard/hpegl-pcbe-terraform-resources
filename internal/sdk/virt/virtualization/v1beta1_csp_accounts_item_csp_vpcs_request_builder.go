package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspVpcsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-vpcs
type V1beta1CspAccountsItemCspVpcsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspVpcsRequestBuilderGetQueryParameters returns a list of virtual private clouds (VPCs) in the specified cloud account.
type V1beta1CspAccountsItemCspVpcsRequestBuilderGetQueryParameters struct {
    // An expression by which to filter the results.These operators compare a property value to a literal value:* 'eq' - Equal* 'ne' - Not equalA 'contains' function filters by a substring match.  Combine it with 'tolower' for a case insensitivesubstring match.Use 'and' and 'or' with parentheses to combine expressions.These fields are available for filtering:* cspId* cspInfo/cspRegion* name* subscriptionInfo/id
    Filter *string `uriparametername:"filter"`
    // A resource property by which to sort, followed by an optional directionindicator: "asc" (ascending, the default) or "desc" (descending).These fields are available for sorting:* cspInfo/cspRegion* id* name
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspAccountsItemCspVpcsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspVpcsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsItemCspVpcsRequestBuilderGetQueryParameters
}
// ByVpcId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspAccounts.item.cspVpcs.item collection
// returns a *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspVpcsRequestBuilder) ByVpcId(vpcId string)(*V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if vpcId != "" {
        urlTplParams["vpcId"] = vpcId
    }
    return NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspAccountsItemCspVpcsRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspVpcsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspVpcsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspVpcsRequestBuilder) {
    m := &V1beta1CspAccountsItemCspVpcsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-vpcs{?filter*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspVpcsRequestBuilder instantiates a new V1beta1CspAccountsItemCspVpcsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspVpcsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspVpcsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspVpcsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of virtual private clouds (VPCs) in the specified cloud account.
// Deprecated: This method is obsolete. Use GetAsCspVpcsGetResponse instead.
// returns a V1beta1CspAccountsItemCspVpcsResponseable when successful
// returns a V1beta1CspAccountsItemCspVpcs400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspVpcs401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspVpcs403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspVpcs404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspVpcs500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspVpcsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspVpcsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspVpcsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspVpcs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspVpcs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspVpcs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspVpcs404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspVpcs500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspVpcsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspVpcsResponseable), nil
}
// GetAsCspVpcsGetResponse returns a list of virtual private clouds (VPCs) in the specified cloud account.
// returns a V1beta1CspAccountsItemCspVpcsGetResponseable when successful
// returns a V1beta1CspAccountsItemCspVpcs400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspVpcs401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspVpcs403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspVpcs404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspVpcs500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspVpcsRequestBuilder) GetAsCspVpcsGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspVpcsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspVpcsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspVpcs400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspVpcs401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspVpcs403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspVpcs404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspVpcs500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspVpcsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspVpcsGetResponseable), nil
}
// ToGetRequestInformation returns a list of virtual private clouds (VPCs) in the specified cloud account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspVpcsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspVpcsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsItemCspVpcsRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspVpcsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspVpcsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspVpcsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
