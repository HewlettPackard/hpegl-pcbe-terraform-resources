package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSoftwareUpdateRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\software-update
type V1beta1SystemsItemSoftwareUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSoftwareUpdateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSoftwareUpdateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemSoftwareUpdateRequestBuilderInternal instantiates a new V1beta1SystemsItemSoftwareUpdateRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareUpdateRequestBuilder) {
    m := &V1beta1SystemsItemSoftwareUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/software-update", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSoftwareUpdateRequestBuilder instantiates a new V1beta1SystemsItemSoftwareUpdateRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSoftwareUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate a software update on the specified system and the hypervisor clusters in it using the specified software catalog version.Before invoking this API, ensure that software update prechecks on the same parameters successfully completed and is valid.
// Deprecated: This method is obsolete. Use PostAsSoftwareUpdatePostResponse instead.
// returns a V1beta1SystemsItemSoftwareUpdateResponseable when successful
// returns a V1beta1SystemsItemSoftwareUpdate400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareUpdate401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareUpdate403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareUpdate404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareUpdate409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareUpdate500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareUpdateRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemSoftwareUpdatePostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwareUpdateRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareUpdateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareUpdate400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareUpdate401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareUpdate403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareUpdate404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareUpdate409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareUpdate500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareUpdateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareUpdateResponseable), nil
}
// PostAsSoftwareUpdatePostResponse initiate a software update on the specified system and the hypervisor clusters in it using the specified software catalog version.Before invoking this API, ensure that software update prechecks on the same parameters successfully completed and is valid.
// returns a V1beta1SystemsItemSoftwareUpdatePostResponseable when successful
// returns a V1beta1SystemsItemSoftwareUpdate400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareUpdate401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareUpdate403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareUpdate404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareUpdate409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareUpdate500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareUpdateRequestBuilder) PostAsSoftwareUpdatePostResponse(ctx context.Context, body V1beta1SystemsItemSoftwareUpdatePostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwareUpdateRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareUpdatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareUpdate400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareUpdate401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareUpdate403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareUpdate404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareUpdate409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareUpdate500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareUpdatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareUpdatePostResponseable), nil
}
// ToPostRequestInformation initiate a software update on the specified system and the hypervisor clusters in it using the specified software catalog version.Before invoking this API, ensure that software update prechecks on the same parameters successfully completed and is valid.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSoftwareUpdateRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemSoftwareUpdatePostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwareUpdateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemSoftwareUpdateRequestBuilder when successful
func (m *V1beta1SystemsItemSoftwareUpdateRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSoftwareUpdateRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareUpdateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
