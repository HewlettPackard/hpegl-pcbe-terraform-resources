package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemCspTagsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\csp-tags
type V1beta1CspAccountsItemCspTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemCspTagsRequestBuilderGetQueryParameters returns the values associated with a given cloud service provider (CSP) tag key attached tocsp-machine-instance and csp-volume resources in the specified cloud account and region.
type V1beta1CspAccountsItemCspTagsRequestBuilderGetQueryParameters struct {
    // A cloud provider region by which to filter the results.
    CspRegion *string `uriparametername:"csp%2Dregion"`
    // An expression by which to filter the results.  A filter for a specified "key"is *required*, and it is the only property available for filtering.
    Filter *string `uriparametername:"filter"`
}
// V1beta1CspAccountsItemCspTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemCspTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsItemCspTagsRequestBuilderGetQueryParameters
}
// NewV1beta1CspAccountsItemCspTagsRequestBuilderInternal instantiates a new V1beta1CspAccountsItemCspTagsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspTagsRequestBuilder) {
    m := &V1beta1CspAccountsItemCspTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/csp-tags?filter={filter}{&csp%2Dregion*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemCspTagsRequestBuilder instantiates a new V1beta1CspAccountsItemCspTagsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemCspTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemCspTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemCspTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the values associated with a given cloud service provider (CSP) tag key attached tocsp-machine-instance and csp-volume resources in the specified cloud account and region.
// Deprecated: This method is obsolete. Use GetAsCspTagsGetResponse instead.
// returns a V1beta1CspAccountsItemCspTagsResponseable when successful
// returns a V1beta1CspAccountsItemCspTags400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspTags401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspTags403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspTags404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspTags500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspTagsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspTagsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspTags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspTags404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspTagsResponseable), nil
}
// GetAsCspTagsGetResponse returns the values associated with a given cloud service provider (CSP) tag key attached tocsp-machine-instance and csp-volume resources in the specified cloud account and region.
// returns a V1beta1CspAccountsItemCspTagsGetResponseable when successful
// returns a V1beta1CspAccountsItemCspTags400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemCspTags401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemCspTags403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemCspTags404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemCspTags500Error error when the service returns a 500 status code
func (m *V1beta1CspAccountsItemCspTagsRequestBuilder) GetAsCspTagsGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspTagsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemCspTagsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemCspTags400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemCspTags401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemCspTags403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemCspTags404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemCspTags500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemCspTagsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemCspTagsGetResponseable), nil
}
// ToGetRequestInformation returns the values associated with a given cloud service provider (CSP) tag key attached tocsp-machine-instance and csp-volume resources in the specified cloud account and region.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemCspTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemCspTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsItemCspTagsRequestBuilder when successful
func (m *V1beta1CspAccountsItemCspTagsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemCspTagsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
