package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-vpcs\{vpcId}
type V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) {
    m := &V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-vpcs/{vpcId}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder instantiates a new V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get details of a virtual private cloud (VPC) in the specified cloud account.
// Deprecated: This method is obsolete. Use GetAsWithVpcGetResponse instead.
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpcResponseable when successful
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspVpcsItemWithVpcResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspVpcsItemWithVpcResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspVpcsItemWithVpcResponseable), nil
}
// GetAsWithVpcGetResponse get details of a virtual private cloud (VPC) in the specified cloud account.
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseable when successful
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspVpcsItemWithVpc500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) GetAsWithVpcGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspVpcsItemWithVpc500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspVpcsItemWithVpcGetResponseable), nil
}
// ToGetRequestInformation get details of a virtual private cloud (VPC) in the specified cloud account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder) {
    return NewV1beta1CspAccountsItemCspVpcsWithVpcItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
