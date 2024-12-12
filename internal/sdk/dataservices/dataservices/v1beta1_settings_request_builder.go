package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
)

// V1beta1SettingsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\settings
type V1beta1SettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SettingsRequestBuilderGetQueryParameters returns all settings values for the current account
type V1beta1SettingsRequestBuilderGetQueryParameters struct {
    // The expression to use for filtering responses. You can filter on the following properties: customerId, id, name, possibleValues, currentValue, settingDescription, lastUpdatedBy, lastUpdatedAt, externalApplicationName. You can combine multiple comparison operators using AND. The comparisons supported are the following:“eq” : Valid for number, boolean and string properties.“gt” : Valid for number or string timestamp properties.“lt” :  Valid for number or string timestamp properties“in” : Valid for an array of stringsSyntax:“eq” : filter=<property> eq <value> {host:port}/data-services/v1beta1/settings?filter=<property> eq <value>“gt” : filter=<property> gt <value> {host:port}/data-services/v1beta1/settings?filter=<property> gt <value>“lt” : filter=<property> lt <value> {host:port}/data-services/v1beta1/settings?filter=<property> lt <value>“in” : filter=<property> in <value> {host:port}/data-services/v1beta1/settings?filter=<property> in <value>* Use AND to filter on multiple properties: {host:port}/data-services/v1beta1/settings?filter=<property1> eq <value1> and <property2> lt <value2>* To filter multiple values on one property e.g. filter=name in ('foo','bar') {host:port}/data-services/v1beta1/settings?filter=foo%bar%20in%20severityExamples:GET /data-services/v1beta1/settings?filter=name eq 'SETTINGNAME'GET /data-services/v1beta1/settings?filter=name eq 'SETTINGNAME' and lastUpdatedBy eq 'CREATED'GET /data-services/v1beta1/settings?filter=relatedObjectType in ('NIMBLE-VOLUME')Filters are supported on following attributes:customerId,id,name,possibleValues,currentValue,settingDescription,lastUpdatedBy,lastUpdatedAt,externalApplicationName
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging. The limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // Use offset in conjunction with limit for paging. The offset is the number of items from the beginning of the result set to the first item included in the response.
    Offset *int32 `uriparametername:"offset"`
    // A comma-separated list of properties to include in response. If this is omitted, all properties are returned.
    Select *string `uriparametername:"select"`
    // The property to sort by followed by a direction indicator ("asc" or "desc"). If no direction indicator is specified the default order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SettingsRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.settings.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1SettingsSettingsItemRequestBuilder when successful
func (m *V1beta1SettingsRequestBuilder) ById(id string)(*V1beta1SettingsSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1SettingsSettingsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.settings.item collection
// returns a *V1beta1SettingsSettingsItemRequestBuilder when successful
func (m *V1beta1SettingsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1SettingsSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1SettingsSettingsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SettingsRequestBuilderInternal instantiates a new V1beta1SettingsRequestBuilder and sets the default values.
func NewV1beta1SettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SettingsRequestBuilder) {
    m := &V1beta1SettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/settings{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SettingsRequestBuilder instantiates a new V1beta1SettingsRequestBuilder and sets the default values.
func NewV1beta1SettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all settings values for the current account
// Deprecated: This method is obsolete. Use GetAsSettingsGetResponse instead.
// returns a V1beta1SettingsResponseable when successful
// returns a V1beta1Settings400Error error when the service returns a 400 status code
// returns a V1beta1Settings401Error error when the service returns a 401 status code
// returns a V1beta1Settings403Error error when the service returns a 403 status code
// returns a V1beta1Settings404Error error when the service returns a 404 status code
// returns a V1beta1Settings422Error error when the service returns a 422 status code
// returns a V1beta1Settings500Error error when the service returns a 500 status code
// returns a V1beta1Settings503Error error when the service returns a 503 status code
func (m *V1beta1SettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SettingsRequestBuilderGetRequestConfiguration)(V1beta1SettingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Settings400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Settings401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Settings403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Settings404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Settings422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Settings500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Settings503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SettingsResponseable), nil
}
// GetAsSettingsGetResponse returns all settings values for the current account
// returns a V1beta1SettingsGetResponseable when successful
// returns a V1beta1Settings400Error error when the service returns a 400 status code
// returns a V1beta1Settings401Error error when the service returns a 401 status code
// returns a V1beta1Settings403Error error when the service returns a 403 status code
// returns a V1beta1Settings404Error error when the service returns a 404 status code
// returns a V1beta1Settings422Error error when the service returns a 422 status code
// returns a V1beta1Settings500Error error when the service returns a 500 status code
// returns a V1beta1Settings503Error error when the service returns a 503 status code
func (m *V1beta1SettingsRequestBuilder) GetAsSettingsGetResponse(ctx context.Context, requestConfiguration *V1beta1SettingsRequestBuilderGetRequestConfiguration)(V1beta1SettingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Settings400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Settings401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Settings403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Settings404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Settings422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Settings500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Settings503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SettingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SettingsGetResponseable), nil
}
// ToGetRequestInformation returns all settings values for the current account
// returns a *RequestInformation when successful
func (m *V1beta1SettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SettingsRequestBuilder when successful
func (m *V1beta1SettingsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SettingsRequestBuilder) {
    return NewV1beta1SettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
