package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1CspAccountsRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\csp-accounts
type V1beta1CspAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1CspAccountsRequestBuilderGetQueryParameters returns a list of cloud service provider (CSP) accounts according to the given queryparameters for paging, filtering, and sorting.
type V1beta1CspAccountsRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the list of CSP accounts returned in the response.The returned set of resources will match the criteria in the filter query parameter.A 'contains' expression can be used to filter the results based on case insensitive substring match.E.g. filter=contains(name, 'test-') will return all volumes with names containing the string 'test-' or 'Test-'.A comparison compares a property name to a literal. The comparisons supported are the following:* 'eq' : A property is equal to value. Valid for number, boolean and string properties.* 'ne' : A property is not equal to value. Valid for number, boolean and string properties.* 'gt' : A property is greater than a value. Valid for number and timestamp properties.* 'lt' : A property is less than a value. Valid for number and timestamp propertiesFiltering is supported with following attributes:* name* suspended* validationStatus* cspId* cspType* validatedAt* id
    Filter *string `uriparametername:"filter"`
    // The limit query parameter should be used in conjunction with offsetfor paging, e.g.: offset=30&limit=10. The limit is the maximumnumber of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // The offset query parameter should be used in conjunction with limitfor paging, e.g. : offset=30&limit=10. The offset is the number ofitems from the beginning of the result set to the first itemincluded in the response.
    Offset *int32 `uriparametername:"offset"`
    // A resource property by which to sort, followed by an optional directionindicator ("asc" or "desc"). If no direction indicator is specified thedefault order is ascending.Sorting is supported on the following properties:* name* suspended* validatationStatus* cspId* cspType* validatedAt* id
    Sort *string `uriparametername:"sort"`
}
// V1beta1CspAccountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1CspAccountsRequestBuilderGetQueryParameters
}
// V1beta1CspAccountsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1CspAccountsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccountId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.cspAccounts.item collection
// returns a *V1beta1CspAccountsAccountItemRequestBuilder when successful
func (m *V1beta1CspAccountsRequestBuilder) ByAccountId(accountId string)(*V1beta1CspAccountsAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accountId != "" {
        urlTplParams["account%2Did"] = accountId
    }
    return NewV1beta1CspAccountsAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1CspAccountsRequestBuilderInternal instantiates a new V1beta1CspAccountsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsRequestBuilder) {
    m := &V1beta1CspAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/csp-accounts{?filter*,limit*,offset*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1CspAccountsRequestBuilder instantiates a new V1beta1CspAccountsRequestBuilder and sets the default values.
func NewV1beta1CspAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1CspAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1CspAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of cloud service provider (CSP) accounts according to the given queryparameters for paging, filtering, and sorting.
// Deprecated: This method is obsolete. Use GetAsCspAccountsGetResponse instead.
// returns a V1beta1CspAccountsResponseable when successful
// returns a V1beta1CspAccounts400Error error when the service returns a 400 status code
// returns a V1beta1CspAccounts401Error error when the service returns a 401 status code
// returns a V1beta1CspAccounts403Error error when the service returns a 403 status code
// returns a V1beta1CspAccounts500Error error when the service returns a 500 status code
// returns a V1beta1CspAccounts503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1CspAccountsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccounts400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccounts401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccounts403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccounts500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccounts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsResponseable), nil
}
// GetAsCspAccountsGetResponse returns a list of cloud service provider (CSP) accounts according to the given queryparameters for paging, filtering, and sorting.
// returns a V1beta1CspAccountsGetResponseable when successful
// returns a V1beta1CspAccounts400Error error when the service returns a 400 status code
// returns a V1beta1CspAccounts401Error error when the service returns a 401 status code
// returns a V1beta1CspAccounts403Error error when the service returns a 403 status code
// returns a V1beta1CspAccounts500Error error when the service returns a 500 status code
// returns a V1beta1CspAccounts503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsRequestBuilder) GetAsCspAccountsGetResponse(ctx context.Context, requestConfiguration *V1beta1CspAccountsRequestBuilderGetRequestConfiguration)(V1beta1CspAccountsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccounts400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccounts401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccounts403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccounts500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccounts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsGetResponseable), nil
}
// Post this API is used to register a new cloud service provider (CSP) account withthe DSCC customer account.
// Deprecated: This method is obsolete. Use PostAsCspAccountsPostResponse instead.
// returns a V1beta1CspAccountsResponseable when successful
// returns a V1beta1CspAccounts400Error error when the service returns a 400 status code
// returns a V1beta1CspAccounts401Error error when the service returns a 401 status code
// returns a V1beta1CspAccounts403Error error when the service returns a 403 status code
// returns a V1beta1CspAccounts412Error error when the service returns a 412 status code
// returns a V1beta1CspAccounts500Error error when the service returns a 500 status code
// returns a V1beta1CspAccounts503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsRequestBuilder) Post(ctx context.Context, body V1beta1CspAccountsPostRequestBodyable, requestConfiguration *V1beta1CspAccountsRequestBuilderPostRequestConfiguration)(V1beta1CspAccountsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccounts400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccounts401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccounts403ErrorFromDiscriminatorValue,
        "412": CreateV1beta1CspAccounts412ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccounts500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccounts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsResponseable), nil
}
// PostAsCspAccountsPostResponse this API is used to register a new cloud service provider (CSP) account withthe DSCC customer account.
// returns a V1beta1CspAccountsPostResponseable when successful
// returns a V1beta1CspAccounts400Error error when the service returns a 400 status code
// returns a V1beta1CspAccounts401Error error when the service returns a 401 status code
// returns a V1beta1CspAccounts403Error error when the service returns a 403 status code
// returns a V1beta1CspAccounts412Error error when the service returns a 412 status code
// returns a V1beta1CspAccounts500Error error when the service returns a 500 status code
// returns a V1beta1CspAccounts503Error error when the service returns a 503 status code
func (m *V1beta1CspAccountsRequestBuilder) PostAsCspAccountsPostResponse(ctx context.Context, body V1beta1CspAccountsPostRequestBodyable, requestConfiguration *V1beta1CspAccountsRequestBuilderPostRequestConfiguration)(V1beta1CspAccountsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1CspAccounts400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1CspAccounts401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1CspAccounts403ErrorFromDiscriminatorValue,
        "412": CreateV1beta1CspAccounts412ErrorFromDiscriminatorValue,
        "500": CreateV1beta1CspAccounts500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1CspAccounts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1CspAccountsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1CspAccountsPostResponseable), nil
}
// ToGetRequestInformation returns a list of cloud service provider (CSP) accounts according to the given queryparameters for paging, filtering, and sorting.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1CspAccountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation this API is used to register a new cloud service provider (CSP) account withthe DSCC customer account.
// returns a *RequestInformation when successful
func (m *V1beta1CspAccountsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1CspAccountsPostRequestBodyable, requestConfiguration *V1beta1CspAccountsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1CspAccountsRequestBuilder when successful
func (m *V1beta1CspAccountsRequestBuilder) WithUrl(rawUrl string)(*V1beta1CspAccountsRequestBuilder) {
    return NewV1beta1CspAccountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
