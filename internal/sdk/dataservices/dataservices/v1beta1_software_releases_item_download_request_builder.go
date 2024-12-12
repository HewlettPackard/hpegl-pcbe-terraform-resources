package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SoftwareReleasesItemDownloadRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-releases\{id}\download
type V1beta1SoftwareReleasesItemDownloadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SoftwareReleasesItemDownloadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SoftwareReleasesItemDownloadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SoftwareReleasesItemDownloadRequestBuilderInternal instantiates a new V1beta1SoftwareReleasesItemDownloadRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesItemDownloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesItemDownloadRequestBuilder) {
    m := &V1beta1SoftwareReleasesItemDownloadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-releases/{id}/download", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareReleasesItemDownloadRequestBuilder instantiates a new V1beta1SoftwareReleasesItemDownloadRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesItemDownloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesItemDownloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareReleasesItemDownloadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post download a file within a Software Release.
// returns a V1beta1SoftwareReleasesItemDownload400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleasesItemDownload401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleasesItemDownload403Error error when the service returns a 403 status code
// returns a V1beta1SoftwareReleasesItemDownload404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareReleasesItemDownload500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleasesItemDownload503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesItemDownloadRequestBuilder) Post(ctx context.Context, body V1beta1SoftwareReleasesItemDownloadPostRequestBodyable, requestConfiguration *V1beta1SoftwareReleasesItemDownloadRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleasesItemDownload400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleasesItemDownload401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SoftwareReleasesItemDownload403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareReleasesItemDownload404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleasesItemDownload500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleasesItemDownload503ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation download a file within a Software Release.
// returns a *RequestInformation when successful
func (m *V1beta1SoftwareReleasesItemDownloadRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SoftwareReleasesItemDownloadPostRequestBodyable, requestConfiguration *V1beta1SoftwareReleasesItemDownloadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SoftwareReleasesItemDownloadRequestBuilder when successful
func (m *V1beta1SoftwareReleasesItemDownloadRequestBuilder) WithUrl(rawUrl string)(*V1beta1SoftwareReleasesItemDownloadRequestBuilder) {
    return NewV1beta1SoftwareReleasesItemDownloadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
