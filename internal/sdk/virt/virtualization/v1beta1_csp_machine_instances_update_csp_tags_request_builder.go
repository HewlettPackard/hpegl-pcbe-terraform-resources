package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-machine-instances\update-csp-tags
type V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostQueryParameters perform specified CSP tag updates to a set of machine instances.  The machine instancesmust not be in a DELETED state.  An error occurrence can result in partial completion.
type V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostQueryParameters struct {
    // Identifies a target machine instance by its ID. A separate instance of this parameteris used for each target machine instance. A maximum of 100 IDs is allowed.
    Id []string `uriparametername:"id"`
}
// V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostQueryParameters
}
// NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilderInternal instantiates a new V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) {
    m := &V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-machine-instances/update-csp-tags{?id*}", pathParameters),
    }
    return m
}
// NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilder instantiates a new V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder and sets the default values.
func NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform specified CSP tag updates to a set of machine instances.  The machine instancesmust not be in a DELETED state.  An error occurrence can result in partial completion.
// Deprecated: This method is obsolete. Use PostAsUpdateCspTagsPostResponse instead.
// returns a V1beta1CspMachineInstancesUpdateCspTagsResponseable when successful
// returns a V1beta1CspMachineInstancesUpdateCspTags400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) Post(ctx context.Context, body V1beta1CspMachineInstancesUpdateCspTagsPostRequestBodyable, requestConfiguration *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesUpdateCspTagsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesUpdateCspTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesUpdateCspTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesUpdateCspTags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesUpdateCspTags404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesUpdateCspTags409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesUpdateCspTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesUpdateCspTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesUpdateCspTagsResponseable), nil
}
// PostAsUpdateCspTagsPostResponse perform specified CSP tag updates to a set of machine instances.  The machine instancesmust not be in a DELETED state.  An error occurrence can result in partial completion.
// returns a V1beta1CspMachineInstancesUpdateCspTagsPostResponseable when successful
// returns a V1beta1CspMachineInstancesUpdateCspTags400Error error when the service returns a 400 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags401Error error when the service returns a 401 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags403Error error when the service returns a 403 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags404Error error when the service returns a 404 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags409Error error when the service returns a 409 status code
// returns a V1beta1CspMachineInstancesUpdateCspTags500Error error when the service returns a 500 status code
func (m *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) PostAsUpdateCspTagsPostResponse(ctx context.Context, body V1beta1CspMachineInstancesUpdateCspTagsPostRequestBodyable, requestConfiguration *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostRequestConfiguration)(V1beta1CspMachineInstancesUpdateCspTagsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspMachineInstancesUpdateCspTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspMachineInstancesUpdateCspTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspMachineInstancesUpdateCspTags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspMachineInstancesUpdateCspTags404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspMachineInstancesUpdateCspTags409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspMachineInstancesUpdateCspTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspMachineInstancesUpdateCspTagsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspMachineInstancesUpdateCspTagsPostResponseable), nil
}
// ToPostRequestInformation perform specified CSP tag updates to a set of machine instances.  The machine instancesmust not be in a DELETED state.  An error occurrence can result in partial completion.
// returns a *RequestInformation when successful
func (m *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1CspMachineInstancesUpdateCspTagsPostRequestBodyable, requestConfiguration *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder when successful
func (m *V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspMachineInstancesUpdateCspTagsRequestBuilder) {
    return NewV1beta1CspMachineInstancesUpdateCspTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
