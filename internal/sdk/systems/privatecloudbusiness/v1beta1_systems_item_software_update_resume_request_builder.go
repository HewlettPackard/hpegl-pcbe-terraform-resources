package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\software-update-resume
type V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSoftwareUpdateResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSoftwareUpdateResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilderInternal instantiates a new V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) {
    m := &V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/software-update-resume", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilder instantiates a new V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resume the failed and paused software update on the specified system.Invoke this API only after fixing the problems reported in software update.
// Deprecated: This method is obsolete. Use PostAsSoftwareUpdateResumePostResponse instead.
// returns a V1beta1SystemsItemSoftwareUpdateResumeResponseable when successful
// returns a V1beta1SystemsItemSoftwareUpdateResume400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareUpdateResumeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareUpdateResume400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareUpdateResume401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareUpdateResume403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareUpdateResume404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareUpdateResume409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareUpdateResume500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareUpdateResumeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareUpdateResumeResponseable), nil
}
// PostAsSoftwareUpdateResumePostResponse resume the failed and paused software update on the specified system.Invoke this API only after fixing the problems reported in software update.
// returns a V1beta1SystemsItemSoftwareUpdateResumePostResponseable when successful
// returns a V1beta1SystemsItemSoftwareUpdateResume400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume403Error error when the service returns a 403 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume409Error error when the service returns a 409 status code
// returns a V1beta1SystemsItemSoftwareUpdateResume500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) PostAsSoftwareUpdateResumePostResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilderPostRequestConfiguration)(V1beta1SystemsItemSoftwareUpdateResumePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSoftwareUpdateResume400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSoftwareUpdateResume401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SystemsItemSoftwareUpdateResume403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSoftwareUpdateResume404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SystemsItemSoftwareUpdateResume409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSoftwareUpdateResume500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSoftwareUpdateResumePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSoftwareUpdateResumePostResponseable), nil
}
// ToPostRequestInformation resume the failed and paused software update on the specified system.Invoke this API only after fixing the problems reported in software update.
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder when successful
func (m *V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSoftwareUpdateResumeRequestBuilder) {
    return NewV1beta1SystemsItemSoftwareUpdateResumeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
