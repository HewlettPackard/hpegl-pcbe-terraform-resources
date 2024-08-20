package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspSubscriptionsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-subscriptions
type V1beta1CspAccountsItemCspSubscriptionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetQueryParameters returns a list of subscriptions in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.For a cloud service provider that doesn't support subscriptions, an empty list is returned.
type V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetQueryParameters struct {
    // A resource property by which to sort, followed by an optional directionindicator: "asc" (ascending, the default) or "desc" (descending).These fields are available for sorting:* cspId* name
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetQueryParameters
}
// BySubscriptionId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspAccounts.item.cspSubscriptions.item collection
// returns a *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) BySubscriptionId(subscriptionId string)(*V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if subscriptionId != "" {
        urlTplParams["subscriptionId"] = subscriptionId
    }
    return NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspSubscriptionsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) {
    m := &V1beta1CspAccountsItemCspSubscriptionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-subscriptions{?sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilder instantiates a new V1beta1CspAccountsItemCspSubscriptionsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of subscriptions in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.For a cloud service provider that doesn't support subscriptions, an empty list is returned.
// Deprecated: This method is obsolete. Use GetAsCspSubscriptionsGetResponse instead.
// returns a V1beta1CspAccountsItemCspSubscriptionsResponseable when successful
// returns a V1beta1CspAccountsItemCspSubscriptions400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubscriptions403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubscriptions404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubscriptions500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubscriptionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubscriptions400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubscriptions401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubscriptions403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubscriptions404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubscriptions500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubscriptionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubscriptionsResponseable), nil
}
// GetAsCspSubscriptionsGetResponse returns a list of subscriptions in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.For a cloud service provider that doesn't support subscriptions, an empty list is returned.
// returns a V1beta1CspAccountsItemCspSubscriptionsGetResponseable when successful
// returns a V1beta1CspAccountsItemCspSubscriptions400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubscriptions403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubscriptions404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubscriptions500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) GetAsCspSubscriptionsGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubscriptionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubscriptions400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubscriptions401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubscriptions403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubscriptions404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubscriptions500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubscriptionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubscriptionsGetResponseable), nil
}
// ToGetRequestInformation returns a list of subscriptions in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.For a cloud service provider that doesn't support subscriptions, an empty list is returned.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubscriptionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
