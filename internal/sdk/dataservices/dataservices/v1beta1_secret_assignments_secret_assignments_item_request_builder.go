package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\secret-assignments\{id}
type V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderInternal instantiates a new V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder and sets the default values.
func NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) {
    m := &V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/secret-assignments/{id}", pathParameters),
    }
    return m
}
// NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder instantiates a new V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder and sets the default values.
func NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get reports the attributes of the specified assignment.
// Deprecated: This method is obsolete. Use GetAsSecretAssignmentsGetResponse instead.
// returns a V1beta1SecretAssignmentsItemSecretAssignmentsResponseable when successful
// returns a V1beta1SecretAssignmentsItemSecretAssignments400Error error when the service returns a 400 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments401Error error when the service returns a 401 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments403Error error when the service returns a 403 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments404Error error when the service returns a 404 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments422Error error when the service returns a 422 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments500Error error when the service returns a 500 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments503Error error when the service returns a 503 status code
func (m *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderGetRequestConfiguration)(V1beta1SecretAssignmentsItemSecretAssignmentsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretAssignmentsItemSecretAssignments400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretAssignmentsItemSecretAssignments401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretAssignmentsItemSecretAssignments403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretAssignmentsItemSecretAssignments404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretAssignmentsItemSecretAssignments422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretAssignmentsItemSecretAssignments500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretAssignmentsItemSecretAssignments503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretAssignmentsItemSecretAssignmentsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretAssignmentsItemSecretAssignmentsResponseable), nil
}
// GetAsSecretAssignmentsGetResponse reports the attributes of the specified assignment.
// returns a V1beta1SecretAssignmentsItemSecretAssignmentsGetResponseable when successful
// returns a V1beta1SecretAssignmentsItemSecretAssignments400Error error when the service returns a 400 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments401Error error when the service returns a 401 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments403Error error when the service returns a 403 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments404Error error when the service returns a 404 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments422Error error when the service returns a 422 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments500Error error when the service returns a 500 status code
// returns a V1beta1SecretAssignmentsItemSecretAssignments503Error error when the service returns a 503 status code
func (m *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) GetAsSecretAssignmentsGetResponse(ctx context.Context, requestConfiguration *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderGetRequestConfiguration)(V1beta1SecretAssignmentsItemSecretAssignmentsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretAssignmentsItemSecretAssignments400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretAssignmentsItemSecretAssignments401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretAssignmentsItemSecretAssignments403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretAssignmentsItemSecretAssignments404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretAssignmentsItemSecretAssignments422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretAssignmentsItemSecretAssignments500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretAssignmentsItemSecretAssignments503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretAssignmentsItemSecretAssignmentsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretAssignmentsItemSecretAssignmentsGetResponseable), nil
}
// ToGetRequestInformation reports the attributes of the specified assignment.
// returns a *RequestInformation when successful
func (m *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder when successful
func (m *V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder) {
    return NewV1beta1SecretAssignmentsSecretAssignmentsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
