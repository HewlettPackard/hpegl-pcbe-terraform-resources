package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-subscriptions\{subscriptionId}
type V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) {
    m := &V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-subscriptions/{subscriptionId}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder instantiates a new V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get details of a subscription in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.
// Deprecated: This method is obsolete. Use GetAsWithSubscriptionGetResponse instead.
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionResponseable when successful
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionResponseable), nil
}
// GetAsWithSubscriptionGetResponse get details of a subscription in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseable when successful
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspSubscriptionsItemWithSubscription500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) GetAsWithSubscriptionGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscription500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspSubscriptionsItemWithSubscriptionGetResponseable), nil
}
// ToGetRequestInformation get details of a subscription in the specified cloud account, where subscription refers toa logical container that holds a collection of connected business or technical resources.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder) {
    return NewV1beta1CspAccountsItemCspSubscriptionsWithSubscriptionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
