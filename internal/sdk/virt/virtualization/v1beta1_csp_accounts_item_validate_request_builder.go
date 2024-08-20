package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsItemValidateRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts\{account-id}\validate
type V1beta1CspAccountsItemValidateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsItemValidateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsItemValidateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1beta1CspAccountsItemValidateRequestBuilderInternal instantiates a new V1beta1CspAccountsItemValidateRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemValidateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemValidateRequestBuilder) {
    m := &V1beta1CspAccountsItemValidateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts/{account%2Did}/validate", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsItemValidateRequestBuilder instantiates a new V1beta1CspAccountsItemValidateRequestBuilder and sets the default values.
func NewV1beta1CspAccountsItemValidateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsItemValidateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsItemValidateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this API is used to validate that DSCC has access to a cloud service provider(CSP) account and that all the required permissions are in place to allow DSCCto manage data protection for the account.
// Deprecated: This method is obsolete. Use PostAsValidatePostResponse instead.
// returns a V1beta1CspAccountsItemValidateResponseable when successful
// returns a V1beta1CspAccountsItemValidate400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemValidate401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemValidate403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemValidate404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemValidate409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemValidate412Error error when the service returns a 412 status code
// returns a V1beta1CspAccountsItemValidate500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemValidate503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsItemValidateRequestBuilder) Post(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemValidateRequestBuilderPostRequestConfiguration)(V1beta1CspAccountsItemValidateResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemValidate400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemValidate401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemValidate403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemValidate404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemValidate409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1CspAccountsItemValidate412ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemValidate500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemValidate503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemValidateResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemValidateResponseable), nil
}
// PostAsValidatePostResponse this API is used to validate that DSCC has access to a cloud service provider(CSP) account and that all the required permissions are in place to allow DSCCto manage data protection for the account.
// returns a V1beta1CspAccountsItemValidatePostResponseable when successful
// returns a V1beta1CspAccountsItemValidate400Error error when the service returns a 400 status code
// returns a V1beta1CspAccountsItemValidate401Error error when the service returns a 401 status code
// returns a V1beta1CspAccountsItemValidate403Error error when the service returns a 403 status code
// returns a V1beta1CspAccountsItemValidate404Error error when the service returns a 404 status code
// returns a V1beta1CspAccountsItemValidate409Error error when the service returns a 409 status code
// returns a V1beta1CspAccountsItemValidate412Error error when the service returns a 412 status code
// returns a V1beta1CspAccountsItemValidate500Error error when the service returns a 500 status code
// returns a V1beta1CspAccountsItemValidate503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsItemValidateRequestBuilder) PostAsValidatePostResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemValidateRequestBuilderPostRequestConfiguration)(V1beta1CspAccountsItemValidatePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccountsItemValidate400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccountsItemValidate401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccountsItemValidate403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1CspAccountsItemValidate404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1CspAccountsItemValidate409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1CspAccountsItemValidate412ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccountsItemValidate500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccountsItemValidate503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsItemValidatePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsItemValidatePostResponseable), nil
}
// ToPostRequestInformation this API is used to validate that DSCC has access to a cloud service provider(CSP) account and that all the required permissions are in place to allow DSCCto manage data protection for the account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsItemValidateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsItemValidateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1CspAccountsItemValidateRequestBuilder when successful
func (m *V1beta1CspAccountsItemValidateRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsItemValidateRequestBuilder) {
    return NewV1beta1CspAccountsItemValidateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
