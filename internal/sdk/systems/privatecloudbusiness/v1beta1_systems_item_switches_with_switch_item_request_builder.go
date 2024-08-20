package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\switches\{switchId}
type V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetQueryParameters get switch details by system id and switch id Retrieving all of the properties for a single switch can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want for retrieve. For example, to get the switch id, name, and serial number, use ?select=id,name,serialNumber
type V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetQueryParameters struct {
    // Query parameter listing the properties of Switch information to fetch.
    Select *string `uriparametername:"select"`
}
// V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetQueryParameters
}
// NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderInternal instantiates a new V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) {
    m := &V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/switches/{switchId}{?select*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder instantiates a new V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get switch details by system id and switch id Retrieving all of the properties for a single switch can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want for retrieve. For example, to get the switch id, name, and serial number, use ?select=id,name,serialNumber
// Deprecated: This method is obsolete. Use GetAsWithSwitchGetResponse instead.
// returns a V1beta1SystemsItemSwitchesItemWithSwitchResponseable when successful
// returns a V1beta1SystemsItemSwitchesItemWithSwitch400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSwitchesItemWithSwitch401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSwitchesItemWithSwitch404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSwitchesItemWithSwitch500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemSwitchesItemWithSwitchResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSwitchesItemWithSwitch400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSwitchesItemWithSwitch401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSwitchesItemWithSwitch404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSwitchesItemWithSwitch500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSwitchesItemWithSwitchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSwitchesItemWithSwitchResponseable), nil
}
// GetAsWithSwitchGetResponse get switch details by system id and switch id Retrieving all of the properties for a single switch can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want for retrieve. For example, to get the switch id, name, and serial number, use ?select=id,name,serialNumber
// returns a V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable when successful
// returns a V1beta1SystemsItemSwitchesItemWithSwitch400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSwitchesItemWithSwitch401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSwitchesItemWithSwitch404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSwitchesItemWithSwitch500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) GetAsWithSwitchGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSwitchesItemWithSwitch400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSwitchesItemWithSwitch401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSwitchesItemWithSwitch404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSwitchesItemWithSwitch500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSwitchesItemWithSwitchGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable), nil
}
// ToGetRequestInformation get switch details by system id and switch id Retrieving all of the properties for a single switch can take a long time because the amount of data is large. Use the select query parameter to choose only the properties you want for retrieve. For example, to get the switch id, name, and serial number, use ?select=id,name,serialNumber
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder when successful
func (m *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) {
    return NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
