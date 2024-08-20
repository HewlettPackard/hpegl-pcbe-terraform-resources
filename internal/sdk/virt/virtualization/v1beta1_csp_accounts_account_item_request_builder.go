package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsAccountItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}
type V1beta1CspAccountsAccountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsAccountItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsAccountItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1CspAccountsAccountItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsAccountItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1CspAccountsAccountItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsAccountItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsAccountItemRequestBuilderInternal instantiates a new V1beta1CspAccountsAccountItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsAccountItemRequestBuilder) {
    m := &V1beta1CspAccountsAccountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsAccountItemRequestBuilder instantiates a new V1beta1CspAccountsAccountItemRequestBuilder and sets the default values.
func NewV1beta1CspAccountsAccountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsAccountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CspResourceGroups the cspResourceGroups property
// returns a *V1beta1CspAccountsItemCspResourceGroupsRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) CspResourceGroups()(*V1beta1CspAccountsItemCspResourceGroupsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspResourceGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspSubnets the cspSubnets property
// returns a *V1beta1CspAccountsItemCspSubnetsRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) CspSubnets()(*V1beta1CspAccountsItemCspSubnetsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspSubnetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspSubscriptions the cspSubscriptions property
// returns a *V1beta1CspAccountsItemCspSubscriptionsRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) CspSubscriptions()(*V1beta1CspAccountsItemCspSubscriptionsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspTagKeys the cspTagKeys property
// returns a *V1beta1CspAccountsItemCspTagKeysRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) CspTagKeys()(*V1beta1CspAccountsItemCspTagKeysRequestBuilder) {
    return NewV1beta1CspAccountsItemCspTagKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspTags the cspTags property
// returns a *V1beta1CspAccountsItemCspTagsRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) CspTags()(*V1beta1CspAccountsItemCspTagsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CspVpcs the cspVpcs property
// returns a *V1beta1CspAccountsItemCspVpcsRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) CspVpcs()(*V1beta1CspAccountsItemCspVpcsRequestBuilder) {
    return NewV1beta1CspAccountsItemCspVpcsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete this API will unregister the cloud service provider (CSP) account from theDSCC customer account.The unregistered CSP account will disable any scheduled jobs orchestrated bythe DSCC.
// Deprecated: This method is obsolete. Use DeleteAsAccountDeleteResponse instead.
// returns a V1beta1CspAccountsItemAccountResponseable when successful
// returns a V1beta1CspAccountsItemAccount400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemAccount401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemAccount403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemAccount404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemAccount409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemAccount500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemAccount503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsAccountItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderDeleteRequestConfiguration)(V1beta1CspAccountsItemAccountResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemAccount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemAccount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemAccount403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemAccount404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemAccount409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemAccount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemAccount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemAccountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemAccountResponseable), nil
}
// DeleteAsAccountDeleteResponse this API will unregister the cloud service provider (CSP) account from theDSCC customer account.The unregistered CSP account will disable any scheduled jobs orchestrated bythe DSCC.
// returns a V1beta1CspAccountsItemAccountDeleteResponseable when successful
// returns a V1beta1CspAccountsItemAccount400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemAccount401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemAccount403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemAccount404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemAccount409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemAccount500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemAccount503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsAccountItemRequestBuilder) DeleteAsAccountDeleteResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderDeleteRequestConfiguration)(V1beta1CspAccountsItemAccountDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemAccount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemAccount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemAccount403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemAccount404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemAccount409ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemAccount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemAccount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemAccountDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemAccountDeleteResponseable), nil
}
// Get returns details of a specified cloud service provider (CSP) account.
// Deprecated: This method is obsolete. Use GetAsAccountGetResponse instead.
// returns a V1beta1CspAccountsItemAccountResponseable when successful
// returns a V1beta1CspAccountsItemAccount400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemAccount401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemAccount403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemAccount500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemAccount503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsAccountItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemAccountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemAccount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemAccount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemAccount403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemAccount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemAccount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemAccountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemAccountResponseable), nil
}
// GetAsAccountGetResponse returns details of a specified cloud service provider (CSP) account.
// returns a V1beta1CspAccountsItemAccountGetResponseable when successful
// returns a V1beta1CspAccountsItemAccount400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemAccount401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemAccount403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemAccount500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemAccount503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsAccountItemRequestBuilder) GetAsAccountGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsItemAccountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemAccount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemAccount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemAccount403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemAccount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemAccount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemAccountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemAccountGetResponseable), nil
}
// Patch this API allows the configuration of the cloud service provider (CSP) accountto be modified.The configuration options include:* name* suspended
// Deprecated: This method is obsolete. Use PatchAsAccountPatchResponse instead.
// returns a V1beta1CspAccountsItemAccountResponseable when successful
// returns a V1beta1CspAccountsItemAccount400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemAccount401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemAccount403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemAccount412Error error when the service returns a 412 status code
// returns a V1beta1CspAccountsItemAccount500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemAccount503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsAccountItemRequestBuilder) Patch(ctx context.Context, body V1beta1CspAccountsItemAccountPatchRequestBodyable, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderPatchRequestConfiguration)(V1beta1CspAccountsItemAccountResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemAccount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemAccount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemAccount403ErrorFromDiscriminatorValue,
        "412": CreateV1beta1CspAccountsItemAccount412ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemAccount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemAccount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemAccountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemAccountResponseable), nil
}
// PatchAsAccountPatchResponse this API allows the configuration of the cloud service provider (CSP) accountto be modified.The configuration options include:* name* suspended
// returns a V1beta1CspAccountsItemAccountPatchResponseable when successful
// returns a V1beta1CspAccountsItemAccount400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemAccount401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemAccount403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemAccount412Error error when the service returns a 412 status code
// returns a V1beta1CspAccountsItemAccount500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemAccount503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsAccountItemRequestBuilder) PatchAsAccountPatchResponse(ctx context.Context, body V1beta1CspAccountsItemAccountPatchRequestBodyable, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderPatchRequestConfiguration)(V1beta1CspAccountsItemAccountPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemAccount400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemAccount401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemAccount403ErrorFromDiscriminatorValue,
        "412": CreateV1beta1CspAccountsItemAccount412ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemAccount500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemAccount503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemAccountPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemAccountPatchResponseable), nil
}
// RdsRefresh the rdsRefresh property
// returns a *V1beta1CspAccountsItemRdsRefreshRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) RdsRefresh()(*V1beta1CspAccountsItemRdsRefreshRequestBuilder) {
    return NewV1beta1CspAccountsItemRdsRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Refresh the refresh property
// returns a *V1beta1CspAccountsItemRefreshRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) Refresh()(*V1beta1CspAccountsItemRefreshRequestBuilder) {
    return NewV1beta1CspAccountsItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation this API will unregister the cloud service provider (CSP) account from theDSCC customer account.The unregistered CSP account will disable any scheduled jobs orchestrated bythe DSCC.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns details of a specified cloud service provider (CSP) account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation this API allows the configuration of the cloud service provider (CSP) accountto be modified.The configuration options include:* name* suspended
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body V1beta1CspAccountsItemAccountPatchRequestBodyable, requestConfiguration *V1beta1CspAccountsAccountItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Validate the validate property
// returns a *V1beta1CspAccountsItemValidateRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) Validate()(*V1beta1CspAccountsItemValidateRequestBuilder) {
    return NewV1beta1CspAccountsItemValidateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsAccountItemRequestBuilder when successful
func (m *V1beta1CspAccountsAccountItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsAccountItemRequestBuilder) {
    return NewV1beta1CspAccountsAccountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
