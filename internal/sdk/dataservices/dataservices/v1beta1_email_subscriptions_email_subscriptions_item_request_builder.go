package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\email-subscriptions\{id}
type V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderInternal instantiates a new V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder and sets the default values.
func NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) {
    m := &V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/email-subscriptions/{id}", pathParameters),
    }
    return m
}
// NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder instantiates a new V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder and sets the default values.
func NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deactivate the email subscription for the user.  The user is identified by the contents of the bearer token that is included in the request.
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions404Error error when the service returns a 404 status code
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions503Error error when the service returns a 503 status code
func (m *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions503ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Post activate the email subscription for the user.  The user is identified by the contents of the bearer token that is included in the request.
// Deprecated: This method is obsolete. Use PostAsEmailSubscriptionsPostResponse instead.
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptionsResponseable when successful
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions404Error error when the service returns a 404 status code
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions503Error error when the service returns a 503 status code
func (m *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderPostRequestConfiguration)(V1beta1EmailSubscriptionsItemEmailSubscriptionsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1EmailSubscriptionsItemEmailSubscriptionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1EmailSubscriptionsItemEmailSubscriptionsResponseable), nil
}
// PostAsEmailSubscriptionsPostResponse activate the email subscription for the user.  The user is identified by the contents of the bearer token that is included in the request.
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseable when successful
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions401Error error when the service returns a 401 status code
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions404Error error when the service returns a 404 status code
// returns a V1beta1EmailSubscriptionsItemEmailSubscriptions503Error error when the service returns a 503 status code
func (m *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) PostAsEmailSubscriptionsPostResponse(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderPostRequestConfiguration)(V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions404ErrorFromDiscriminatorValue,
        "503": CreateV1beta1EmailSubscriptionsItemEmailSubscriptions503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1EmailSubscriptionsItemEmailSubscriptionsPostResponseable), nil
}
// ToDeleteRequestInformation deactivate the email subscription for the user.  The user is identified by the contents of the bearer token that is included in the request.
// returns a *RequestInformation when successful
func (m *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation activate the email subscription for the user.  The user is identified by the contents of the bearer token that is included in the request.
// returns a *RequestInformation when successful
func (m *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder when successful
func (m *V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder) {
    return NewV1beta1EmailSubscriptionsEmailSubscriptionsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
