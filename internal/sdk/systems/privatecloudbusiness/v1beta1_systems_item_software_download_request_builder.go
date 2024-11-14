package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSoftwareDownloadRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\software-download
type V1beta1SystemsItemSoftwareDownloadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSoftwareDownloadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSoftwareDownloadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemSoftwareDownloadRequestBuilderInternal instantiates a new V1beta1SystemsItemSoftwareDownloadRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareDownloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareDownloadRequestBuilder) {
    m := &V1beta1SystemsItemSoftwareDownloadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/software-download", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSoftwareDownloadRequestBuilder instantiates a new V1beta1SystemsItemSoftwareDownloadRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareDownloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareDownloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSoftwareDownloadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate a software download operation for the specified system and target software catalog version.This operation ensures the required software packages are available on the system for a software update to the target software catalog version.
// Deprecated: This method is obsolete. Use PostAsSoftwareDownloadPostResponse instead.
// returns a V1beta1SystemsItemSoftwareDownloadResponseable when successful
// returns a V1beta1SystemsItemSoftwareDownload400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareDownload401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareDownload403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareDownload404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareDownload405Error error when the service returns a 405 status code
// returns a V1beta1SystemsItemSoftwareDownload409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareDownload500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareDownloadRequestBuilder) Post(ctx context.Context, body V1beta1SystemsItemSoftwareDownloadPostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwareDownloadRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareDownloadResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareDownload400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareDownload401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareDownload403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareDownload404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1SystemsItemSoftwareDownload405ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareDownload409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareDownload500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareDownloadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareDownloadResponseable), nil
}
// PostAsSoftwareDownloadPostResponse initiate a software download operation for the specified system and target software catalog version.This operation ensures the required software packages are available on the system for a software update to the target software catalog version.
// returns a V1beta1SystemsItemSoftwareDownloadPostResponseable when successful
// returns a V1beta1SystemsItemSoftwareDownload400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareDownload401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareDownload403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareDownload404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareDownload405Error error when the service returns a 405 status code
// returns a V1beta1SystemsItemSoftwareDownload409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareDownload500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareDownloadRequestBuilder) PostAsSoftwareDownloadPostResponse(ctx context.Context, body V1beta1SystemsItemSoftwareDownloadPostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwareDownloadRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareDownloadPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareDownload400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareDownload401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareDownload403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareDownload404ErrorFromDiscriminatorValue,
        "405": CreateV1beta1SystemsItemSoftwareDownload405ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareDownload409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareDownload500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareDownloadPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareDownloadPostResponseable), nil
}
// ToPostRequestInformation initiate a software download operation for the specified system and target software catalog version.This operation ensures the required software packages are available on the system for a software update to the target software catalog version.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSoftwareDownloadRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SystemsItemSoftwareDownloadPostRequestBodyable, requestConfiguration *V1beta1SystemsItemSoftwareDownloadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemSoftwareDownloadRequestBuilder when successful
func (m *V1beta1SystemsItemSoftwareDownloadRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSoftwareDownloadRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareDownloadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
