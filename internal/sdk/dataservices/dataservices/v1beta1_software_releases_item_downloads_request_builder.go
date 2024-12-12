package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SoftwareReleasesItemDownloadsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-releases\{id}\downloads
type V1beta1SoftwareReleasesItemDownloadsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SoftwareReleasesItemDownloadsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SoftwareReleasesItemDownloadsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SoftwareReleasesItemDownloadsRequestBuilderInternal instantiates a new V1beta1SoftwareReleasesItemDownloadsRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesItemDownloadsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesItemDownloadsRequestBuilder) {
    m := &V1beta1SoftwareReleasesItemDownloadsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-releases/{id}/downloads", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareReleasesItemDownloadsRequestBuilder instantiates a new V1beta1SoftwareReleasesItemDownloadsRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesItemDownloadsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesItemDownloadsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareReleasesItemDownloadsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the URL to download a file within a Software Release.
// Deprecated: This method is obsolete. Use PostAsDownloadsPostResponse instead.
// returns a V1beta1SoftwareReleasesItemDownloadsResponseable when successful
// returns a V1beta1SoftwareReleasesItemDownloads400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleasesItemDownloads401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleasesItemDownloads403Error error when the service returns a 403 status code
// returns a V1beta1SoftwareReleasesItemDownloads404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareReleasesItemDownloads500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleasesItemDownloads503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesItemDownloadsRequestBuilder) Post(ctx context.Context, body V1beta1SoftwareReleasesItemDownloadsPostRequestBodyable, requestConfiguration *V1beta1SoftwareReleasesItemDownloadsRequestBuilderPostRequestConfiguration)(V1beta1SoftwareReleasesItemDownloadsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleasesItemDownloads400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleasesItemDownloads401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SoftwareReleasesItemDownloads403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareReleasesItemDownloads404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleasesItemDownloads500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleasesItemDownloads503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareReleasesItemDownloadsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareReleasesItemDownloadsResponseable), nil
}
// PostAsDownloadsPostResponse get the URL to download a file within a Software Release.
// returns a V1beta1SoftwareReleasesItemDownloadsPostResponseable when successful
// returns a V1beta1SoftwareReleasesItemDownloads400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleasesItemDownloads401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleasesItemDownloads403Error error when the service returns a 403 status code
// returns a V1beta1SoftwareReleasesItemDownloads404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareReleasesItemDownloads500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleasesItemDownloads503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesItemDownloadsRequestBuilder) PostAsDownloadsPostResponse(ctx context.Context, body V1beta1SoftwareReleasesItemDownloadsPostRequestBodyable, requestConfiguration *V1beta1SoftwareReleasesItemDownloadsRequestBuilderPostRequestConfiguration)(V1beta1SoftwareReleasesItemDownloadsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleasesItemDownloads400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleasesItemDownloads401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SoftwareReleasesItemDownloads403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareReleasesItemDownloads404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleasesItemDownloads500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleasesItemDownloads503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareReleasesItemDownloadsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareReleasesItemDownloadsPostResponseable), nil
}
// ToPostRequestInformation get the URL to download a file within a Software Release.
// returns a *RequestInformation when successful
func (m *V1beta1SoftwareReleasesItemDownloadsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1SoftwareReleasesItemDownloadsPostRequestBodyable, requestConfiguration *V1beta1SoftwareReleasesItemDownloadsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SoftwareReleasesItemDownloadsRequestBuilder when successful
func (m *V1beta1SoftwareReleasesItemDownloadsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SoftwareReleasesItemDownloadsRequestBuilder) {
    return NewV1beta1SoftwareReleasesItemDownloadsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
