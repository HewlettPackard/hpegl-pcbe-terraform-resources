package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspTagKeysRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-tag-keys
type V1beta1CspAccountsItemCspTagKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspTagKeysRequestBuilderGetQueryParameters returns the cloud service provider (CSP) tag keys attached to csp-machine-instance and csp-volumeresources in the specified cloud account and region.
type V1beta1CspAccountsItemCspTagKeysRequestBuilderGetQueryParameters struct {
    // A cloud provider region by which to filter the results.
    CspRegion *string `uriparametername:"csp%2Dregion"`
}
// V1beta1CspAccountsItemCspTagKeysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspTagKeysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsItemCspTagKeysRequestBuilderGetQueryParameters
}
// NewV1beta1CspAccountsItemCspTagKeysRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspTagKeysRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspTagKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspTagKeysRequestBuilder) {
    m := &V1beta1CspAccountsItemCspTagKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-tag-keys{?csp%2Dregion*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspTagKeysRequestBuilder instantiates a new V1beta1CspAccountsItemCspTagKeysRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspTagKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspTagKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspTagKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the cloud service provider (CSP) tag keys attached to csp-machine-instance and csp-volumeresources in the specified cloud account and region.
// Deprecated: This method is obsolete. Use GetAsCspTagKeysGetResponse instead.
// returns a V1beta1CspAccountsItemCspTagKeysResponseable when successful
// returns a V1beta1CspAccountsItemCspTagKeys400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspTagKeys401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspTagKeys403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspTagKeys404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspTagKeys500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspTagKeysRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspTagKeysRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspTagKeysResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspTagKeys400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspTagKeys401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspTagKeys403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspTagKeys404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspTagKeys500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspTagKeysResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspTagKeysResponseable), nil
}
// GetAsCspTagKeysGetResponse returns the cloud service provider (CSP) tag keys attached to csp-machine-instance and csp-volumeresources in the specified cloud account and region.
// returns a V1beta1CspAccountsItemCspTagKeysGetResponseable when successful
// returns a V1beta1CspAccountsItemCspTagKeys400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspTagKeys401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspTagKeys403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspTagKeys404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspTagKeys500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspTagKeysRequestBuilder) GetAsCspTagKeysGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspTagKeysRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspTagKeysGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspTagKeys400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspTagKeys401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspTagKeys403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspTagKeys404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspTagKeys500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspTagKeysGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspTagKeysGetResponseable), nil
}
// ToGetRequestInformation returns the cloud service provider (CSP) tag keys attached to csp-machine-instance and csp-volumeresources in the specified cloud account and region.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspTagKeysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspTagKeysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsItemCspTagKeysRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspTagKeysRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspTagKeysRequestBuilder) {
    return NewV1beta1CspAccountsItemCspTagKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
