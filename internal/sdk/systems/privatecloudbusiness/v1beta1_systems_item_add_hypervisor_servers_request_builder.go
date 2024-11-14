package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemAddHypervisorServersRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\add-hypervisor-servers
type V1beta1SystemsItemAddHypervisorServersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemAddHypervisorServersRequestBuilderInternal instantiates a new V1beta1SystemsItemAddHypervisorServersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemAddHypervisorServersRequestBuilder) {
    m := &V1beta1SystemsItemAddHypervisorServersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/add-hypervisor-servers", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemAddHypervisorServersRequestBuilder instantiates a new V1beta1SystemsItemAddHypervisorServersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemAddHypervisorServersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemAddHypervisorServersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post adds the specified hypervisor servers to the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// Deprecated: This method is obsolete. Use PostAsAddHypervisorServersPostResponse instead.
// returns a V1beta1SystemsItemAddHypervisorServersResponseable when successful
// returns a V1beta1SystemsItemAddHypervisorServers400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemAddHypervisorServers401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemAddHypervisorServers403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemAddHypervisorServers404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemAddHypervisorServers409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemAddHypervisorServers500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemAddHypervisorServersRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemAddHypervisorServersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemAddHypervisorServersResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemAddHypervisorServers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemAddHypervisorServers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemAddHypervisorServers403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemAddHypervisorServers404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemAddHypervisorServers409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemAddHypervisorServers500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemAddHypervisorServersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemAddHypervisorServersResponseable), nil
}
// PostAsAddHypervisorServersPostResponse adds the specified hypervisor servers to the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a V1beta1SystemsItemAddHypervisorServersPostResponseable when successful
// returns a V1beta1SystemsItemAddHypervisorServers400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemAddHypervisorServers401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemAddHypervisorServers403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemAddHypervisorServers404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemAddHypervisorServers409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemAddHypervisorServers500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemAddHypervisorServersRequestBuilder) PostAsAddHypervisorServersPostResponse(ctx context.Context, body V1beta1SystemsItemAddHypervisorServersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemAddHypervisorServersPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemAddHypervisorServers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemAddHypervisorServers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemAddHypervisorServers403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemAddHypervisorServers404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemAddHypervisorServers409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemAddHypervisorServers500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemAddHypervisorServersPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemAddHypervisorServersPostResponseable), nil
}
// ToPostRequestInformation adds the specified hypervisor servers to the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemAddHypervisorServersRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemAddHypervisorServersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAddHypervisorServersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemAddHypervisorServersRequestBuilder when successful
func (m *V1beta1SystemsItemAddHypervisorServersRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemAddHypervisorServersRequestBuilder) {
    return NewV1beta1SystemsItemAddHypervisorServersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
