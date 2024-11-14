package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSoftwarePrechecksRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\software-prechecks
type V1beta1SystemsItemSoftwarePrechecksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSoftwarePrechecksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSoftwarePrechecksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemSoftwarePrechecksRequestBuilderInternal instantiates a new V1beta1SystemsItemSoftwarePrechecksRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwarePrechecksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwarePrechecksRequestBuilder) {
    m := &V1beta1SystemsItemSoftwarePrechecksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/software-prechecks", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSoftwarePrechecksRequestBuilder instantiates a new V1beta1SystemsItemSoftwarePrechecksRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwarePrechecksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwarePrechecksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSoftwarePrechecksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate a set of software update prechecks on the specified system and hypervisor clusters in it using the specified software catalog version.This operation validates whether the software on system and its constituent clusters are ready to be updated to the target software catalog version.
// Deprecated: This method is obsolete. Use PostAsSoftwarePrechecksPostResponse instead.
// returns a V1beta1SystemsItemSoftwarePrechecksResponseable when successful
// returns a V1beta1SystemsItemSoftwarePrechecks400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwarePrechecks401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwarePrechecks403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwarePrechecks404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwarePrechecks409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwarePrechecks500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwarePrechecksRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemSoftwarePrechecksPostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwarePrechecksRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwarePrechecksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwarePrechecks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwarePrechecks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwarePrechecks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwarePrechecks404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwarePrechecks409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwarePrechecks500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwarePrechecksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwarePrechecksResponseable), nil
}
// PostAsSoftwarePrechecksPostResponse initiate a set of software update prechecks on the specified system and hypervisor clusters in it using the specified software catalog version.This operation validates whether the software on system and its constituent clusters are ready to be updated to the target software catalog version.
// returns a V1beta1SystemsItemSoftwarePrechecksPostResponseable when successful
// returns a V1beta1SystemsItemSoftwarePrechecks400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwarePrechecks401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwarePrechecks403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwarePrechecks404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwarePrechecks409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwarePrechecks500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwarePrechecksRequestBuilder) PostAsSoftwarePrechecksPostResponse(ctx context.Context, body V1beta1SystemsItemSoftwarePrechecksPostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwarePrechecksRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwarePrechecksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwarePrechecks400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwarePrechecks401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwarePrechecks403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwarePrechecks404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwarePrechecks409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwarePrechecks500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwarePrechecksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwarePrechecksPostResponseable), nil
}
// ToPostRequestInformation initiate a set of software update prechecks on the specified system and hypervisor clusters in it using the specified software catalog version.This operation validates whether the software on system and its constituent clusters are ready to be updated to the target software catalog version.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSoftwarePrechecksRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemSoftwarePrechecksPostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwarePrechecksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemSoftwarePrechecksRequestBuilder when successful
func (m *V1beta1SystemsItemSoftwarePrechecksRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSoftwarePrechecksRequestBuilder) {
    return NewV1beta1SystemsItemSoftwarePrechecksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
