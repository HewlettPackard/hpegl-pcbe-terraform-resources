package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemRemoveHypervisorServersRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\remove-hypervisor-servers
type V1beta1SystemsItemRemoveHypervisorServersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemRemoveHypervisorServersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemRemoveHypervisorServersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemRemoveHypervisorServersRequestBuilderInternal instantiates a new V1beta1SystemsItemRemoveHypervisorServersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorServersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) {
    m := &V1beta1SystemsItemRemoveHypervisorServersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/remove-hypervisor-servers", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemRemoveHypervisorServersRequestBuilder instantiates a new V1beta1SystemsItemRemoveHypervisorServersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorServersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemRemoveHypervisorServersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post removes the specified hypervisor servers from the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// Deprecated: This method is obsolete. Use PostAsRemoveHypervisorServersPostResponse instead.
// returns a V1beta1SystemsItemRemoveHypervisorServersResponseable when successful
// returns a V1beta1SystemsItemRemoveHypervisorServers400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemRemoveHypervisorServersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemRemoveHypervisorServersRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemRemoveHypervisorServersResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemRemoveHypervisorServers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemRemoveHypervisorServers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemRemoveHypervisorServers403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemRemoveHypervisorServers404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemRemoveHypervisorServers409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemRemoveHypervisorServers500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemRemoveHypervisorServersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemRemoveHypervisorServersResponseable), nil
}
// PostAsRemoveHypervisorServersPostResponse removes the specified hypervisor servers from the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a V1beta1SystemsItemRemoveHypervisorServersPostResponseable when successful
// returns a V1beta1SystemsItemRemoveHypervisorServers400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemRemoveHypervisorServers500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) PostAsRemoveHypervisorServersPostResponse(ctx context.Context, body V1beta1SystemsItemRemoveHypervisorServersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemRemoveHypervisorServersRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemRemoveHypervisorServersPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemRemoveHypervisorServers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemRemoveHypervisorServers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemRemoveHypervisorServers403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemRemoveHypervisorServers404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemRemoveHypervisorServers409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemRemoveHypervisorServers500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemRemoveHypervisorServersPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemRemoveHypervisorServersPostResponseable), nil
}
// ToPostRequestInformation removes the specified hypervisor servers from the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemRemoveHypervisorServersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemRemoveHypervisorServersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemRemoveHypervisorServersRequestBuilder when successful
func (m *V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemRemoveHypervisorServersRequestBuilder) {
    return NewV1beta1SystemsItemRemoveHypervisorServersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
