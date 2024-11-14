package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\remove-hypervisor-clusters
type V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemRemoveHypervisorClustersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemRemoveHypervisorClustersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemRemoveHypervisorClustersRequestBuilderInternal instantiates a new V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorClustersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) {
    m := &V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/remove-hypervisor-clusters", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemRemoveHypervisorClustersRequestBuilder instantiates a new V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder and sets the default values.
func NewV1beta1SystemsItemRemoveHypervisorClustersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemRemoveHypervisorClustersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post removes the specified hypervisor clusters from the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// Deprecated: This method is obsolete. Use PostAsRemoveHypervisorClustersPostResponse instead.
// returns a V1beta1SystemsItemRemoveHypervisorClustersResponseable when successful
// returns a V1beta1SystemsItemRemoveHypervisorClusters400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemRemoveHypervisorClustersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemRemoveHypervisorClustersResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemRemoveHypervisorClusters400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemRemoveHypervisorClusters401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemRemoveHypervisorClusters403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemRemoveHypervisorClusters404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemRemoveHypervisorClusters409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemRemoveHypervisorClusters500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemRemoveHypervisorClustersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemRemoveHypervisorClustersResponseable), nil
}
// PostAsRemoveHypervisorClustersPostResponse removes the specified hypervisor clusters from the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a V1beta1SystemsItemRemoveHypervisorClustersPostResponseable when successful
// returns a V1beta1SystemsItemRemoveHypervisorClusters400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemRemoveHypervisorClusters500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) PostAsRemoveHypervisorClustersPostResponse(ctx context.Context, body V1beta1SystemsItemRemoveHypervisorClustersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemRemoveHypervisorClustersPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemRemoveHypervisorClusters400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemRemoveHypervisorClusters401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemRemoveHypervisorClusters403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemRemoveHypervisorClusters404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemRemoveHypervisorClusters409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemRemoveHypervisorClusters500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemRemoveHypervisorClustersPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemRemoveHypervisorClustersPostResponseable), nil
}
// ToPostRequestInformation removes the specified hypervisor clusters from the specified system. The user must have permissions to update this system. Returns a task identifier to be used by clients to track the progress of the operation.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemRemoveHypervisorClustersPostRequestBodyable, requestConfiguration *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder when successful
func (m *V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemRemoveHypervisorClustersRequestBuilder) {
    return NewV1beta1SystemsItemRemoveHypervisorClustersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
