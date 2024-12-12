package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SettingsSettingsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\settings\{id}
type V1beta1SettingsSettingsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SettingsSettingsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SettingsSettingsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1SettingsSettingsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SettingsSettingsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1SettingsSettingsItemRequestBuilderInternal instantiates a new V1beta1SettingsSettingsItemRequestBuilder and sets the default values.
func NewV1beta1SettingsSettingsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SettingsSettingsItemRequestBuilder) {
    m := &V1beta1SettingsSettingsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/settings/{id}", pathParameters),
    }
    return m
}
// NewV1beta1SettingsSettingsItemRequestBuilder instantiates a new V1beta1SettingsSettingsItemRequestBuilder and sets the default values.
func NewV1beta1SettingsSettingsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SettingsSettingsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SettingsSettingsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the details of the setting with the given ID. If a value does not exist for the setting, the API returns the default value.
// Deprecated: This method is obsolete. Use GetAsSettingsGetResponse instead.
// returns a V1beta1SettingsItemSettingsResponseable when successful
// returns a V1beta1SettingsItemSettings400Error error when the service returns a 400 status code
// returns a V1beta1SettingsItemSettings401Error error when the service returns a 401 status code
// returns a V1beta1SettingsItemSettings403Error error when the service returns a 403 status code
// returns a V1beta1SettingsItemSettings404Error error when the service returns a 404 status code
// returns a V1beta1SettingsItemSettings422Error error when the service returns a 422 status code
// returns a V1beta1SettingsItemSettings500Error error when the service returns a 500 status code
// returns a V1beta1SettingsItemSettings503Error error when the service returns a 503 status code
func (m *V1beta1SettingsSettingsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SettingsSettingsItemRequestBuilderGetRequestConfiguration)(V1beta1SettingsItemSettingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SettingsItemSettings400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SettingsItemSettings401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SettingsItemSettings403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SettingsItemSettings404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SettingsItemSettings422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SettingsItemSettings500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SettingsItemSettings503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SettingsItemSettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SettingsItemSettingsResponseable), nil
}
// GetAsSettingsGetResponse returns the details of the setting with the given ID. If a value does not exist for the setting, the API returns the default value.
// returns a V1beta1SettingsItemSettingsGetResponseable when successful
// returns a V1beta1SettingsItemSettings400Error error when the service returns a 400 status code
// returns a V1beta1SettingsItemSettings401Error error when the service returns a 401 status code
// returns a V1beta1SettingsItemSettings403Error error when the service returns a 403 status code
// returns a V1beta1SettingsItemSettings404Error error when the service returns a 404 status code
// returns a V1beta1SettingsItemSettings422Error error when the service returns a 422 status code
// returns a V1beta1SettingsItemSettings500Error error when the service returns a 500 status code
// returns a V1beta1SettingsItemSettings503Error error when the service returns a 503 status code
func (m *V1beta1SettingsSettingsItemRequestBuilder) GetAsSettingsGetResponse(ctx context.Context, requestConfiguration *V1beta1SettingsSettingsItemRequestBuilderGetRequestConfiguration)(V1beta1SettingsItemSettingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SettingsItemSettings400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SettingsItemSettings401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SettingsItemSettings403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SettingsItemSettings404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SettingsItemSettings422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SettingsItemSettings500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SettingsItemSettings503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SettingsItemSettingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SettingsItemSettingsGetResponseable), nil
}
// Patch changes the value of the given setting.
// Deprecated: This method is obsolete. Use PatchAsSettingsPatchResponse instead.
// returns a V1beta1SettingsItemSettingsResponseable when successful
// returns a V1beta1SettingsItemSettings400Error error when the service returns a 400 status code
// returns a V1beta1SettingsItemSettings401Error error when the service returns a 401 status code
// returns a V1beta1SettingsItemSettings403Error error when the service returns a 403 status code
// returns a V1beta1SettingsItemSettings404Error error when the service returns a 404 status code
// returns a V1beta1SettingsItemSettings412Error error when the service returns a 412 status code
// returns a V1beta1SettingsItemSettings422Error error when the service returns a 422 status code
// returns a V1beta1SettingsItemSettings500Error error when the service returns a 500 status code
// returns a V1beta1SettingsItemSettings503Error error when the service returns a 503 status code
func (m *V1beta1SettingsSettingsItemRequestBuilder) Patch(ctx context.Context, body V1beta1SettingsItemSettingsPatchRequestBodyable, requestConfiguration *V1beta1SettingsSettingsItemRequestBuilderPatchRequestConfiguration)(V1beta1SettingsItemSettingsResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SettingsItemSettings400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SettingsItemSettings401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SettingsItemSettings403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SettingsItemSettings404ErrorFromDiscriminatorValue,
        "412": CreateV1beta1SettingsItemSettings412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SettingsItemSettings422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SettingsItemSettings500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SettingsItemSettings503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SettingsItemSettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SettingsItemSettingsResponseable), nil
}
// PatchAsSettingsPatchResponse changes the value of the given setting.
// returns a V1beta1SettingsItemSettingsPatchResponseable when successful
// returns a V1beta1SettingsItemSettings400Error error when the service returns a 400 status code
// returns a V1beta1SettingsItemSettings401Error error when the service returns a 401 status code
// returns a V1beta1SettingsItemSettings403Error error when the service returns a 403 status code
// returns a V1beta1SettingsItemSettings404Error error when the service returns a 404 status code
// returns a V1beta1SettingsItemSettings412Error error when the service returns a 412 status code
// returns a V1beta1SettingsItemSettings422Error error when the service returns a 422 status code
// returns a V1beta1SettingsItemSettings500Error error when the service returns a 500 status code
// returns a V1beta1SettingsItemSettings503Error error when the service returns a 503 status code
func (m *V1beta1SettingsSettingsItemRequestBuilder) PatchAsSettingsPatchResponse(ctx context.Context, body V1beta1SettingsItemSettingsPatchRequestBodyable, requestConfiguration *V1beta1SettingsSettingsItemRequestBuilderPatchRequestConfiguration)(V1beta1SettingsItemSettingsPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SettingsItemSettings400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SettingsItemSettings401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SettingsItemSettings403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SettingsItemSettings404ErrorFromDiscriminatorValue,
        "412": CreateV1beta1SettingsItemSettings412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SettingsItemSettings422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SettingsItemSettings500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SettingsItemSettings503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SettingsItemSettingsPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SettingsItemSettingsPatchResponseable), nil
}
// ToGetRequestInformation returns the details of the setting with the given ID. If a value does not exist for the setting, the API returns the default value.
// returns a *RequestInformation when successful
func (m *V1beta1SettingsSettingsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SettingsSettingsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation changes the value of the given setting.
// returns a *RequestInformation when successful
func (m *V1beta1SettingsSettingsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1SettingsItemSettingsPatchRequestBodyable, requestConfiguration *V1beta1SettingsSettingsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/merge-patch+json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1SettingsSettingsItemRequestBuilder when successful
func (m *V1beta1SettingsSettingsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SettingsSettingsItemRequestBuilder) {
    return NewV1beta1SettingsSettingsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
