package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstancesItemRefreshRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instances\{id}\refresh
type V1beta1CspMachineInstancesItemRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstancesItemRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesItemRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspMachineInstancesItemRefreshRequestBuilderInternal instantiates a new V1beta1CspMachineInstancesItemRefreshRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesItemRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesItemRefreshRequestBuilder) {
    m := &V1beta1CspMachineInstancesItemRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instances/{id}/refresh", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstancesItemRefreshRequestBuilder instantiates a new V1beta1CspMachineInstancesItemRefreshRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesItemRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesItemRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstancesItemRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates the properties of a specified machine instance to match the settings in the cloud account. Also updates the CSP properties of any CSP volumes to which it is attached. Protection groups with dynamic membership are updated to include or exclude the machine instance according to its updated settings.
// Deprecated: This method is obsolete. Use PostAsRefreshPostResponse instead.
// returns a V1beta1CspMachineInstancesItemRefreshResponseable when successful
// returns a V1beta1CspMachineInstancesItemRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesItemRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemRefreshRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesItemRefreshResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspMachineInstancesItemRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemRefreshResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemRefreshResponseable), nil
}
// PostAsRefreshPostResponse updates the properties of a specified machine instance to match the settings in the cloud account. Also updates the CSP properties of any CSP volumes to which it is attached. Protection groups with dynamic membership are updated to include or exclude the machine instance according to its updated settings.
// returns a V1beta1CspMachineInstancesItemRefreshPostResponseable when successful
// returns a V1beta1CspMachineInstancesItemRefresh401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesItemRefresh403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesItemRefresh404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesItemRefresh409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesItemRefresh500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesItemRefreshRequestBuilder) PostAsRefreshPostResponse(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemRefreshRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesItemRefreshPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1CspMachineInstancesItemRefresh401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesItemRefresh403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesItemRefresh404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesItemRefresh409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesItemRefresh500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesItemRefreshPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesItemRefreshPostResponseable), nil
}
// ToPostRequestInformation updates the properties of a specified machine instance to match the settings in the cloud account. Also updates the CSP properties of any CSP volumes to which it is attached. Protection groups with dynamic membership are updated to include or exclude the machine instance according to its updated settings.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesItemRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspMachineInstancesItemRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspMachineInstancesItemRefreshRequestBuilder when successful
func (m *V1beta1CspMachineInstancesItemRefreshRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstancesItemRefreshRequestBuilder) {
    return NewV1beta1CspMachineInstancesItemRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
