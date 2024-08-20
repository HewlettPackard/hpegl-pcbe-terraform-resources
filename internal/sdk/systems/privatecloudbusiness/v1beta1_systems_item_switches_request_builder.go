package privatecloudbusiness

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1beta1SystemsItemSwitchesRequestBuilder builds and executes requests for operations under \private-cloud-business\v1beta1\systems\{-id}\switches
type V1beta1SystemsItemSwitchesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1SystemsItemSwitchesRequestBuilderGetQueryParameters get switch details by system idRetrieving all of the properties for all switches in a system can take a long time because the amount of data  is large.  Use the 'select' query parameter to choose only the properties you want to retrieve for each switch. For example, to get the name and serial number for a specific switch use ?select=id,name,serialNumber
type V1beta1SystemsItemSwitchesRequestBuilderGetQueryParameters struct {
    // The expression to filter responses.
    Filter *string `uriparametername:"filter"`
    // Use limit in conjunction with offset for paging, e.g.: offset=30&&limit=10. Limit is the maximum number of items to include in the response.
    Limit *int32 `uriparametername:"limit"`
    // Use offset in conjunction with limit for paging, e.g.: offset=30&&limit=10. Offset is the number of items from the beginning of the result set to the first item included in the response.
    Offset *int32 `uriparametername:"offset"`
    // Query parameter listing the properties of Switch information to fetch.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a directionindicator ("asc" or "desc"). If no direction indicator is specified thedefault order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SystemsItemSwitchesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SystemsItemSwitchesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SystemsItemSwitchesRequestBuilderGetQueryParameters
}
// BySwitchId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/systems.privateCloudBusiness.v1beta1.systems.item.switches.item collection
// returns a *V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder when successful
func (m *V1beta1SystemsItemSwitchesRequestBuilder) BySwitchId(switchId string)(*V1beta1SystemsItemSwitchesWithSwitchItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if switchId != "" {
        urlTplParams["switchId"] = switchId
    }
    return NewV1beta1SystemsItemSwitchesWithSwitchItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SystemsItemSwitchesRequestBuilderInternal instantiates a new V1beta1SystemsItemSwitchesRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSwitchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSwitchesRequestBuilder) {
    m := &V1beta1SystemsItemSwitchesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/private-cloud-business/v1beta1/systems/{%2Did}/switches{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SystemsItemSwitchesRequestBuilder instantiates a new V1beta1SystemsItemSwitchesRequestBuilder and sets the default values.
func NewV1beta1SystemsItemSwitchesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SystemsItemSwitchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SystemsItemSwitchesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get switch details by system idRetrieving all of the properties for all switches in a system can take a long time because the amount of data  is large.  Use the 'select' query parameter to choose only the properties you want to retrieve for each switch. For example, to get the name and serial number for a specific switch use ?select=id,name,serialNumber
// Deprecated: This method is obsolete. Use GetAsSwitchesGetResponse instead.
// returns a V1beta1SystemsItemSwitchesResponseable when successful
// returns a V1beta1SystemsItemSwitches400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSwitches401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSwitches404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSwitches500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSwitchesRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SystemsItemSwitchesRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemSwitchesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSwitches400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSwitches401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSwitches404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSwitches500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSwitchesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSwitchesResponseable), nil
}
// GetAsSwitchesGetResponse get switch details by system idRetrieving all of the properties for all switches in a system can take a long time because the amount of data  is large.  Use the 'select' query parameter to choose only the properties you want to retrieve for each switch. For example, to get the name and serial number for a specific switch use ?select=id,name,serialNumber
// returns a V1beta1SystemsItemSwitchesGetResponseable when successful
// returns a V1beta1SystemsItemSwitches400Error error when the service returns a 400 status code
// returns a V1beta1SystemsItemSwitches401Error error when the service returns a 401 status code
// returns a V1beta1SystemsItemSwitches404Error error when the service returns a 404 status code
// returns a V1beta1SystemsItemSwitches500Error error when the service returns a 500 status code
func (m *V1beta1SystemsItemSwitchesRequestBuilder) GetAsSwitchesGetResponse(ctx context.Context, requestConfiguration *V1beta1SystemsItemSwitchesRequestBuilderGetRequestConfiguration)(V1beta1SystemsItemSwitchesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SystemsItemSwitches400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SystemsItemSwitches401ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SystemsItemSwitches404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SystemsItemSwitches500ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SystemsItemSwitchesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SystemsItemSwitchesGetResponseable), nil
}
// ToGetRequestInformation get switch details by system idRetrieving all of the properties for all switches in a system can take a long time because the amount of data  is large.  Use the 'select' query parameter to choose only the properties you want to retrieve for each switch. For example, to get the name and serial number for a specific switch use ?select=id,name,serialNumber
// returns a *RequestInformation when successful
func (m *V1beta1SystemsItemSwitchesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SystemsItemSwitchesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SystemsItemSwitchesRequestBuilder when successful
func (m *V1beta1SystemsItemSwitchesRequestBuilder) WithUrl(rawUrl string)(*V1beta1SystemsItemSwitchesRequestBuilder) {
    return NewV1beta1SystemsItemSwitchesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
