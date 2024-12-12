package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\software-releases\{id}
type V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetQueryParameters get a single Software Release by its ID.
type V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetQueryParameters struct {
    // Comma separated properties to return in the result. If omitted, all properties will be returned. Selecting nested properties of an object is not supported.
    Select []string `uriparametername:"select"`
}
// V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetQueryParameters
}
// NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderInternal instantiates a new V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) {
    m := &V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/software-releases/{id}{?select}", pathParameters),
    }
    return m
}
// NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder instantiates a new V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder and sets the default values.
func NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Download the download property
// returns a *V1beta1SoftwareReleasesItemDownloadRequestBuilder when successful
func (m *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) Download()(*V1beta1SoftwareReleasesItemDownloadRequestBuilder) {
    return NewV1beta1SoftwareReleasesItemDownloadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Downloads the downloads property
// returns a *V1beta1SoftwareReleasesItemDownloadsRequestBuilder when successful
func (m *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) Downloads()(*V1beta1SoftwareReleasesItemDownloadsRequestBuilder) {
    return NewV1beta1SoftwareReleasesItemDownloadsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a single Software Release by its ID.
// Deprecated: This method is obsolete. Use GetAsSoftwareReleasesGetResponse instead.
// returns a V1beta1SoftwareReleasesItemSoftwareReleasesResponseable when successful
// returns a V1beta1SoftwareReleasesItemSoftwareReleases400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetRequestConfiguration)(V1beta1SoftwareReleasesItemSoftwareReleasesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleasesItemSoftwareReleases400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleasesItemSoftwareReleases401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareReleasesItemSoftwareReleases404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleasesItemSoftwareReleases500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleasesItemSoftwareReleases503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareReleasesItemSoftwareReleasesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareReleasesItemSoftwareReleasesResponseable), nil
}
// GetAsSoftwareReleasesGetResponse get a single Software Release by its ID.
// returns a V1beta1SoftwareReleasesItemSoftwareReleasesGetResponseable when successful
// returns a V1beta1SoftwareReleasesItemSoftwareReleases400Error error when the service returns a 400 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases401Error error when the service returns a 401 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases404Error error when the service returns a 404 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases500Error error when the service returns a 500 status code
// returns a V1beta1SoftwareReleasesItemSoftwareReleases503Error error when the service returns a 503 status code
func (m *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) GetAsSoftwareReleasesGetResponse(ctx context.Context, requestConfiguration *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetRequestConfiguration)(V1beta1SoftwareReleasesItemSoftwareReleasesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SoftwareReleasesItemSoftwareReleases400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SoftwareReleasesItemSoftwareReleases401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SoftwareReleasesItemSoftwareReleases404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SoftwareReleasesItemSoftwareReleases500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SoftwareReleasesItemSoftwareReleases503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SoftwareReleasesItemSoftwareReleasesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SoftwareReleasesItemSoftwareReleasesGetResponseable), nil
}
// ToGetRequestInformation get a single Software Release by its ID.
// returns a *RequestInformation when successful
func (m *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder when successful
func (m *V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder) {
    return NewV1beta1SoftwareReleasesSoftwareReleasesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
