package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1DatastoresItemUnmountRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\datastores\{id}\unmount
type V1beta1DatastoresItemUnmountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DatastoresItemUnmountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresItemUnmountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DatastoresItemUnmountRequestBuilderInternal instantiates a new V1beta1DatastoresItemUnmountRequestBuilder and sets the default values.
func NewV1beta1DatastoresItemUnmountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresItemUnmountRequestBuilder) {
    m := &V1beta1DatastoresItemUnmountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/datastores/{id}/unmount", pathParameters),
    }
    return m
}
// NewV1beta1DatastoresItemUnmountRequestBuilder instantiates a new V1beta1DatastoresItemUnmountRequestBuilder and sets the default values.
func NewV1beta1DatastoresItemUnmountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresItemUnmountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DatastoresItemUnmountRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unmount the datastore
// Deprecated: This method is obsolete. Use PostAsUnmountPostResponse instead.
// returns a V1beta1DatastoresItemUnmountResponseable when successful
// returns a V1beta1DatastoresItemUnmount400Error error when the service returns a 400 status code
// returns a V1beta1DatastoresItemUnmount401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemUnmount403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemUnmount404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemUnmount500Error error when the service returns a 500 status code
// returns a V1beta1DatastoresItemUnmount503Error error when the service returns a 503 status code
func (m *V1beta1DatastoresItemUnmountRequestBuilder) Post(ctx context.Context, body []string, requestConfiguration *V1beta1DatastoresItemUnmountRequestBuilderPostRequestConfiguration)(V1beta1DatastoresItemUnmountResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DatastoresItemUnmount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DatastoresItemUnmount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemUnmount403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemUnmount404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemUnmount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DatastoresItemUnmount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresItemUnmountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresItemUnmountResponseable), nil
}
// PostAsUnmountPostResponse unmount the datastore
// returns a V1beta1DatastoresItemUnmountPostResponseable when successful
// returns a V1beta1DatastoresItemUnmount400Error error when the service returns a 400 status code
// returns a V1beta1DatastoresItemUnmount401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemUnmount403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemUnmount404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemUnmount500Error error when the service returns a 500 status code
// returns a V1beta1DatastoresItemUnmount503Error error when the service returns a 503 status code
func (m *V1beta1DatastoresItemUnmountRequestBuilder) PostAsUnmountPostResponse(ctx context.Context, body []string, requestConfiguration *V1beta1DatastoresItemUnmountRequestBuilderPostRequestConfiguration)(V1beta1DatastoresItemUnmountPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DatastoresItemUnmount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DatastoresItemUnmount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemUnmount403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemUnmount404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemUnmount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DatastoresItemUnmount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresItemUnmountPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresItemUnmountPostResponseable), nil
}
// ToPostRequestInformation unmount the datastore
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresItemUnmountRequestBuilder) ToPostRequestInformation(ctx context.Context, body []string, requestConfiguration *V1beta1DatastoresItemUnmountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    cast := make([]interface{}, len(body))
    for i, v := range body {
        cast[i] = v
    }
    requestInfo.SetContentFromScalarCollection(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", cast)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1DatastoresItemUnmountRequestBuilder when successful
func (m *V1beta1DatastoresItemUnmountRequestBuilder) WithUrl(rawUrl string)(*V1beta1DatastoresItemUnmountRequestBuilder) {
    return NewV1beta1DatastoresItemUnmountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
