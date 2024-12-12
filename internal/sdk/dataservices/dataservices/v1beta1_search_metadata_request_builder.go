package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SearchMetadataRequestBuilder builds and executes requests for operations under \data-services\v1beta1\search-metadata
type V1beta1SearchMetadataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SearchMetadataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SearchMetadataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SearchMetadataRequestBuilderInternal instantiates a new V1beta1SearchMetadataRequestBuilder and sets the default values.
func NewV1beta1SearchMetadataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SearchMetadataRequestBuilder) {
    m := &V1beta1SearchMetadataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/search-metadata", pathParameters),
    }
    return m
}
// NewV1beta1SearchMetadataRequestBuilder instantiates a new V1beta1SearchMetadataRequestBuilder and sets the default values.
func NewV1beta1SearchMetadataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SearchMetadataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SearchMetadataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get metadata provides descriptive and human readable attributes for machinereadable names and identifiers that appear in other Global Search APIrequests and responses.Returns lists of values for:- types: resource types that Global Search is capable of searching- services: DSCC services that Global Search is capable of searching
// Deprecated: This method is obsolete. Use GetAsSearchMetadataGetResponse instead.
// returns a V1beta1SearchMetadataResponseable when successful
// returns a V1beta1SearchMetadata400Error error when the service returns a 400 status code
// returns a V1beta1SearchMetadata401Error error when the service returns a 401 status code
// returns a V1beta1SearchMetadata500Error error when the service returns a 500 status code
func (m *V1beta1SearchMetadataRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SearchMetadataRequestBuilderGetRequestConfiguration)(V1beta1SearchMetadataResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SearchMetadata400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SearchMetadata401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SearchMetadata500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SearchMetadataResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SearchMetadataResponseable), nil
}
// GetAsSearchMetadataGetResponse metadata provides descriptive and human readable attributes for machinereadable names and identifiers that appear in other Global Search APIrequests and responses.Returns lists of values for:- types: resource types that Global Search is capable of searching- services: DSCC services that Global Search is capable of searching
// returns a V1beta1SearchMetadataGetResponseable when successful
// returns a V1beta1SearchMetadata400Error error when the service returns a 400 status code
// returns a V1beta1SearchMetadata401Error error when the service returns a 401 status code
// returns a V1beta1SearchMetadata500Error error when the service returns a 500 status code
func (m *V1beta1SearchMetadataRequestBuilder) GetAsSearchMetadataGetResponse(ctx context.Context, requestConfiguration *V1beta1SearchMetadataRequestBuilderGetRequestConfiguration)(V1beta1SearchMetadataGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SearchMetadata400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SearchMetadata401ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SearchMetadata500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SearchMetadataGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SearchMetadataGetResponseable), nil
}
// ToGetRequestInformation metadata provides descriptive and human readable attributes for machinereadable names and identifiers that appear in other Global Search APIrequests and responses.Returns lists of values for:- types: resource types that Global Search is capable of searching- services: DSCC services that Global Search is capable of searching
// returns a *RequestInformation when successful
func (m *V1beta1SearchMetadataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SearchMetadataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SearchMetadataRequestBuilder when successful
func (m *V1beta1SearchMetadataRequestBuilder) WithUrl(rawUrl string)(*V1beta1SearchMetadataRequestBuilder) {
    return NewV1beta1SearchMetadataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
