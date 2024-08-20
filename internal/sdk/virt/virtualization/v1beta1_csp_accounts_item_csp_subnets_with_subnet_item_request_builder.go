package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-subnets\{subnetId}
type V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) {
    m := &V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-subnets/{subnetId}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder instantiates a new V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get details of a subnet in the specified cloud account.
// Deprecated: This method is obsolete. Use GetAsWithSubnetGetResponse instead.
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseable when successful
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubnetsItemWithSubnetResponseable), nil
}
// GetAsWithSubnetGetResponse get details of a subnet in the specified cloud account.
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable when successful
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubnetsItemWithSubnet500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) GetAsWithSubnetGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnet500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubnetsItemWithSubnetGetResponseable), nil
}
// ToGetRequestInformation get details of a subnet in the specified cloud account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder) {
    return NewV1beta1CspAccountsItemCspSubnetsWithSubnetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
