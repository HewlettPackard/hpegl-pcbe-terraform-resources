package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemAssignSecretRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\assign-secret
type V1beta1SystemsItemAssignSecretRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemAssignSecretRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemAssignSecretRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemAssignSecretRequestBuilderInternal instantiates a new V1beta1SystemsItemAssignSecretRequestBuilder and sets the default values.
func NewV1beta1SystemsItemAssignSecretRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemAssignSecretRequestBuilder) {
    m := &V1beta1SystemsItemAssignSecretRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/assign-secret", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemAssignSecretRequestBuilder instantiates a new V1beta1SystemsItemAssignSecretRequestBuilder and sets the default values.
func NewV1beta1SystemsItemAssignSecretRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemAssignSecretRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemAssignSecretRequestBuilderInternal(urlParams, requestAdapter)
}
// Post changes which secret is used to access the specified servers in a particular way.System update permission is required. A returned task identifier can be used to track the progress of the operation.
// Deprecated: This method is obsolete. Use PostAsAssignSecretPostResponse instead.
// returns a V1beta1SystemsItemAssignSecretResponseable when successful
// returns a V1beta1SystemsItemAssignSecret400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemAssignSecret401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemAssignSecret403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemAssignSecret404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemAssignSecret409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemAssignSecret500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemAssignSecretRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemAssignSecretPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAssignSecretRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemAssignSecretResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemAssignSecret400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemAssignSecret401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemAssignSecret403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemAssignSecret404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemAssignSecret409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemAssignSecret500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemAssignSecretResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemAssignSecretResponseable), nil
}
// PostAsAssignSecretPostResponse changes which secret is used to access the specified servers in a particular way.System update permission is required. A returned task identifier can be used to track the progress of the operation.
// returns a V1beta1SystemsItemAssignSecretPostResponseable when successful
// returns a V1beta1SystemsItemAssignSecret400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemAssignSecret401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemAssignSecret403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemAssignSecret404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemAssignSecret409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemAssignSecret500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemAssignSecretRequestBuilder) PostAsAssignSecretPostResponse(ctx context.Context, body V1beta1SystemsItemAssignSecretPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAssignSecretRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemAssignSecretPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemAssignSecret400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemAssignSecret401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemAssignSecret403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemAssignSecret404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemAssignSecret409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemAssignSecret500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemAssignSecretPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemAssignSecretPostResponseable), nil
}
// ToPostRequestInformation changes which secret is used to access the specified servers in a particular way.System update permission is required. A returned task identifier can be used to track the progress of the operation.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemAssignSecretRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemAssignSecretPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAssignSecretRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemsItemAssignSecretRequestBuilder when successful
func (m *V1beta1SystemsItemAssignSecretRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemAssignSecretRequestBuilder) {
    return NewV1beta1SystemsItemAssignSecretRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
