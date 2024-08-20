package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemAddHypervisorClusterRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\add-hypervisor-cluster
type V1beta1SystemsItemAddHypervisorClusterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemAddHypervisorClusterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemAddHypervisorClusterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemAddHypervisorClusterRequestBuilderInternal instantiates a new V1beta1SystemsItemAddHypervisorClusterRequestBuilder and sets the default values.
func NewV1beta1SystemsItemAddHypervisorClusterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemAddHypervisorClusterRequestBuilder) {
    m := &V1beta1SystemsItemAddHypervisorClusterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/add-hypervisor-cluster", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemAddHypervisorClusterRequestBuilder instantiates a new V1beta1SystemsItemAddHypervisorClusterRequestBuilder and sets the default values.
func NewV1beta1SystemsItemAddHypervisorClusterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemAddHypervisorClusterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemAddHypervisorClusterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post adds the specified hypervisor cluster to the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// Deprecated: This method is obsolete. Use PostAsAddHypervisorClusterPostResponse instead.
// returns a V1beta1SystemsItemAddHypervisorClusterResponseable when successful
// returns a V1beta1SystemsItemAddHypervisorCluster400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemAddHypervisorCluster401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemAddHypervisorCluster403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemAddHypervisorCluster404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemAddHypervisorCluster409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemAddHypervisorCluster500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemAddHypervisorClusterRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemAddHypervisorClusterPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAddHypervisorClusterRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemAddHypervisorClusterResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemAddHypervisorCluster400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemAddHypervisorCluster401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemAddHypervisorCluster403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemAddHypervisorCluster404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemAddHypervisorCluster409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemAddHypervisorCluster500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemAddHypervisorClusterResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemAddHypervisorClusterResponseable), nil
}
// PostAsAddHypervisorClusterPostResponse adds the specified hypervisor cluster to the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a V1beta1SystemsItemAddHypervisorClusterPostResponseable when successful
// returns a V1beta1SystemsItemAddHypervisorCluster400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemAddHypervisorCluster401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemAddHypervisorCluster403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemAddHypervisorCluster404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemAddHypervisorCluster409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemAddHypervisorCluster500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemAddHypervisorClusterRequestBuilder) PostAsAddHypervisorClusterPostResponse(ctx context.Context, body V1beta1SystemsItemAddHypervisorClusterPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAddHypervisorClusterRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemAddHypervisorClusterPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemAddHypervisorCluster400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemAddHypervisorCluster401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemAddHypervisorCluster403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemAddHypervisorCluster404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemAddHypervisorCluster409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemAddHypervisorCluster500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemAddHypervisorClusterPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemAddHypervisorClusterPostResponseable), nil
}
// ToPostRequestInformation adds the specified hypervisor cluster to the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemAddHypervisorClusterRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemAddHypervisorClusterPostRequestBodyable, requestConfiguration *V1beta1SystemsItemAddHypervisorClusterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemAddHypervisorClusterRequestBuilder when successful
func (m *V1beta1SystemsItemAddHypervisorClusterRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemAddHypervisorClusterRequestBuilder) {
    return NewV1beta1SystemsItemAddHypervisorClusterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
