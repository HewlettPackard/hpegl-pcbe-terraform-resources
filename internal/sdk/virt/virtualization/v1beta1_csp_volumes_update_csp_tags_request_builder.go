package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspVolumesUpdateCspTagsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-volumes\update-csp-tags
type V1beta1CspVolumesUpdateCspTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspVolumesUpdateCspTagsRequestBuilderPostQueryParameters perform specified CSP tag updates to a set of volumes.  The volumes must not be in aDELETED state.  An error occurrence can result in partial completion.
type V1beta1CspVolumesUpdateCspTagsRequestBuilderPostQueryParameters struct {
    // Identifies a target volume by its ID. A separate instance of this parameteris used for each target volume. A maximum of 100 IDs is allowed.
    Id []string `uriparametername:"id"`
}
// V1beta1CspVolumesUpdateCspTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspVolumesUpdateCspTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspVolumesUpdateCspTagsRequestBuilderPostQueryParameters
}
// NewV1beta1CspVolumesUpdateCspTagsRequestBuilderInternal instantiates a new V1beta1CspVolumesUpdateCspTagsRequestBuilder and sets the default values.
func NewV1beta1CspVolumesUpdateCspTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesUpdateCspTagsRequestBuilder) {
    m := &V1beta1CspVolumesUpdateCspTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-volumes/update-csp-tags{?id*}", pathParameters),
    }
    return m
}
// NewV1beta1CspVolumesUpdateCspTagsRequestBuilder instantiates a new V1beta1CspVolumesUpdateCspTagsRequestBuilder and sets the default values.
func NewV1beta1CspVolumesUpdateCspTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspVolumesUpdateCspTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspVolumesUpdateCspTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform specified CSP tag updates to a set of volumes.  The volumes must not be in aDELETED state.  An error occurrence can result in partial completion.
// Deprecated: This method is obsolete. Use PostAsUpdateCspTagsPostResponse instead.
// returns a V1beta1CspVolumesUpdateCspTagsResponseable when successful
// returns a V1beta1CspVolumesUpdateCspTags400Error error when the service returns a 400 status code
// returns a V1beta1CspVolumesUpdateCspTags401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumesUpdateCspTags403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumesUpdateCspTags404Error error when the service returns a 404 status code
// returns a V1beta1CspVolumesUpdateCspTags409Error error when the service returns a 409 status code
// returns a V1beta1CspVolumesUpdateCspTags500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesUpdateCspTagsRequestBuilder) Post(ctx context.Context, body V1beta1CspVolumesUpdateCspTagsPostRequestBodyable, requestConfiguration *V1beta1CspVolumesUpdateCspTagsRequestBuilderPostRequestConfiguration)(V1beta1CspVolumesUpdateCspTagsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspVolumesUpdateCspTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspVolumesUpdateCspTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumesUpdateCspTags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspVolumesUpdateCspTags404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspVolumesUpdateCspTags409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumesUpdateCspTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesUpdateCspTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesUpdateCspTagsResponseable), nil
}
// PostAsUpdateCspTagsPostResponse perform specified CSP tag updates to a set of volumes.  The volumes must not be in aDELETED state.  An error occurrence can result in partial completion.
// returns a V1beta1CspVolumesUpdateCspTagsPostResponseable when successful
// returns a V1beta1CspVolumesUpdateCspTags400Error error when the service returns a 400 status code
// returns a V1beta1CspVolumesUpdateCspTags401Error error when the service returns a 401 status code
// returns a V1beta1CspVolumesUpdateCspTags403Error error when the service returns a 403 status code
// returns a V1beta1CspVolumesUpdateCspTags404Error error when the service returns a 404 status code
// returns a V1beta1CspVolumesUpdateCspTags409Error error when the service returns a 409 status code
// returns a V1beta1CspVolumesUpdateCspTags500Error error when the service returns a 500 status code
func (m *V1beta1CspVolumesUpdateCspTagsRequestBuilder) PostAsUpdateCspTagsPostResponse(ctx context.Context, body V1beta1CspVolumesUpdateCspTagsPostRequestBodyable, requestConfiguration *V1beta1CspVolumesUpdateCspTagsRequestBuilderPostRequestConfiguration)(V1beta1CspVolumesUpdateCspTagsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspVolumesUpdateCspTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspVolumesUpdateCspTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspVolumesUpdateCspTags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspVolumesUpdateCspTags404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspVolumesUpdateCspTags409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspVolumesUpdateCspTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspVolumesUpdateCspTagsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspVolumesUpdateCspTagsPostResponseable), nil
}
// ToPostRequestInformation perform specified CSP tag updates to a set of volumes.  The volumes must not be in aDELETED state.  An error occurrence can result in partial completion.
// returns a *RequestInformation when successful
func (m *V1beta1CspVolumesUpdateCspTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1CspVolumesUpdateCspTagsPostRequestBodyable, requestConfiguration *V1beta1CspVolumesUpdateCspTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
// returns a *V1beta1CspVolumesUpdateCspTagsRequestBuilder when successful
func (m *V1beta1CspVolumesUpdateCspTagsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspVolumesUpdateCspTagsRequestBuilder) {
    return NewV1beta1CspVolumesUpdateCspTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
