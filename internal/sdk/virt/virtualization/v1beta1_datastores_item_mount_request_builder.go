package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1DatastoresItemMountRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\datastores\{id}\mount
type V1beta1DatastoresItemMountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DatastoresItemMountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DatastoresItemMountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1DatastoresItemMountRequestBuilderInternal instantiates a new V1beta1DatastoresItemMountRequestBuilder and sets the default values.
func NewV1beta1DatastoresItemMountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresItemMountRequestBuilder) {
    m := &V1beta1DatastoresItemMountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/datastores/{id}/mount", pathParameters),
    }
    return m
}
// NewV1beta1DatastoresItemMountRequestBuilder instantiates a new V1beta1DatastoresItemMountRequestBuilder and sets the default values.
func NewV1beta1DatastoresItemMountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DatastoresItemMountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DatastoresItemMountRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mount the datastore
// Deprecated: This method is obsolete. Use PostAsMountPostResponse instead.
// returns a V1beta1DatastoresItemMountResponseable when successful
// returns a V1beta1DatastoresItemMount400Error error when the service returns a 400 status code
// returns a V1beta1DatastoresItemMount401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemMount403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemMount404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemMount500Error error when the service returns a 500 status code
// returns a V1beta1DatastoresItemMount503Error error when the service returns a 503 status code
func (m *V1beta1DatastoresItemMountRequestBuilder) Post(ctx context.Context, body []string, requestConfiguration *V1beta1DatastoresItemMountRequestBuilderPostRequestConfiguration)(V1beta1DatastoresItemMountResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DatastoresItemMount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DatastoresItemMount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemMount403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemMount404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemMount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DatastoresItemMount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresItemMountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresItemMountResponseable), nil
}
// PostAsMountPostResponse mount the datastore
// returns a V1beta1DatastoresItemMountPostResponseable when successful
// returns a V1beta1DatastoresItemMount400Error error when the service returns a 400 status code
// returns a V1beta1DatastoresItemMount401Error error when the service returns a 401 status code
// returns a V1beta1DatastoresItemMount403Error error when the service returns a 403 status code
// returns a V1beta1DatastoresItemMount404Error error when the service returns a 404 status code
// returns a V1beta1DatastoresItemMount500Error error when the service returns a 500 status code
// returns a V1beta1DatastoresItemMount503Error error when the service returns a 503 status code
func (m *V1beta1DatastoresItemMountRequestBuilder) PostAsMountPostResponse(ctx context.Context, body []string, requestConfiguration *V1beta1DatastoresItemMountRequestBuilderPostRequestConfiguration)(V1beta1DatastoresItemMountPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DatastoresItemMount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DatastoresItemMount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DatastoresItemMount403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1DatastoresItemMount404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DatastoresItemMount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DatastoresItemMount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DatastoresItemMountPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DatastoresItemMountPostResponseable), nil
}
// ToPostRequestInformation mount the datastore
// returns a *RequestInformation when successful
func (m *V1beta1DatastoresItemMountRequestBuilder) ToPostRequestInformation(ctx context.Context, body []string, requestConfiguration *V1beta1DatastoresItemMountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1DatastoresItemMountRequestBuilder when successful
func (m *V1beta1DatastoresItemMountRequestBuilder) WithUrl(rawUrl string)(*V1beta1DatastoresItemMountRequestBuilder) {
    return NewV1beta1DatastoresItemMountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
