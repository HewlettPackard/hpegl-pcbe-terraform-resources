package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\data-services-connectors
type V1beta1DataServicesConnectorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1beta1DataServicesConnectorsRequestBuilderGetQueryParameters returns a list of Data Services Connectors according to the given queryparameters for paging, filtering, and sorting.
type V1beta1DataServicesConnectorsRequestBuilderGetQueryParameters struct {
    // An expression by which to filter the results.The returned set of resources will match the criteria in the filter query parameter.Filtering is supported with following attributes: * id * name * hostname * serialNumber * status * state * connectionState  * softwareVersion * persona * deploymentType * storageSystemSerialNumber
    Filter *string `uriparametername:"filter"`
    // The maximum number of items to include in the response.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Limit *int32 `uriparametername:"limit"`
    // The number of items to omit from the beginning of the result set.The offset and limit query parameters are used in conjunction for pagination,for example "offset=30&limit=10" indicates the fourth page of 10 items.
    Offset *int32 `uriparametername:"offset"`
    // A list of properties to include in the response.
    Select *string `uriparametername:"select"`
    // A resource property by which to sort, followed by an optional directionindicator ("asc" or "desc"). If no direction indicator is specified thedefault order is ascending.Sorting is supported on the following properties:* name* generation* connectionState* serialNumber* softwareVersion* status* state* stateReason* id* platform* vCpu* totalMemoryinGb* poweredOnAt* createdAt* interfaces.network.defaultGateway* dateTime.methodDateTimeSet* ntp.status* ntp.state* ntp.stateReason* persona* deploymentType* storageSystemSerialNumber
    Sort *string `uriparametername:"sort"`
}
// V1beta1DataServicesConnectorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1DataServicesConnectorsRequestBuilderGetQueryParameters
}
// V1beta1DataServicesConnectorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1DataServicesConnectorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.dataServicesConnectors.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsRequestBuilder) ById(id string)(*V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.dataServicesConnectors.item collection
// returns a *V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1DataServicesConnectorsDataServicesConnectorsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1DataServicesConnectorsRequestBuilderInternal instantiates a new V1beta1DataServicesConnectorsRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsRequestBuilder) {
    m := &V1beta1DataServicesConnectorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/data-services-connectors{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1DataServicesConnectorsRequestBuilder instantiates a new V1beta1DataServicesConnectorsRequestBuilder and sets the default values.
func NewV1beta1DataServicesConnectorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1DataServicesConnectorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1DataServicesConnectorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of Data Services Connectors according to the given queryparameters for paging, filtering, and sorting.
// Deprecated: This method is obsolete. Use GetAsDataServicesConnectorsGetResponse instead.
// returns a V1beta1DataServicesConnectorsResponseable when successful
// returns a V1beta1DataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsRequestBuilderGetRequestConfiguration)(V1beta1DataServicesConnectorsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectors403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectors503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DataServicesConnectorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DataServicesConnectorsResponseable), nil
}
// GetAsDataServicesConnectorsGetResponse returns a list of Data Services Connectors according to the given queryparameters for paging, filtering, and sorting.
// returns a V1beta1DataServicesConnectorsGetResponseable when successful
// returns a V1beta1DataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsRequestBuilder) GetAsDataServicesConnectorsGetResponse(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsRequestBuilderGetRequestConfiguration)(V1beta1DataServicesConnectorsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectors403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectors503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1DataServicesConnectorsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1DataServicesConnectorsGetResponseable), nil
}
// Post this API is used to register a new Data Services Connector with the DSCC customer account. A unique Activation Code must be provided to register the Data Services Connector.
// returns a UntypedNodeable when successful
// returns a V1beta1DataServicesConnectors400Error error when the service returns a 400 status code
// returns a V1beta1DataServicesConnectors401Error error when the service returns a 401 status code
// returns a V1beta1DataServicesConnectors403Error error when the service returns a 403 status code
// returns a V1beta1DataServicesConnectors500Error error when the service returns a 500 status code
// returns a V1beta1DataServicesConnectors503Error error when the service returns a 503 status code
func (m *V1beta1DataServicesConnectorsRequestBuilder) Post(ctx context.Context, body V1beta1DataServicesConnectorsPostRequestBodyable, requestConfiguration *V1beta1DataServicesConnectorsRequestBuilderPostRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1DataServicesConnectors400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1DataServicesConnectors401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1DataServicesConnectors403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1DataServicesConnectors500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1DataServicesConnectors503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable), nil
}
// ToGetRequestInformation returns a list of Data Services Connectors according to the given queryparameters for paging, filtering, and sorting.
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1DataServicesConnectorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation this API is used to register a new Data Services Connector with the DSCC customer account. A unique Activation Code must be provided to register the Data Services Connector.
// returns a *RequestInformation when successful
func (m *V1beta1DataServicesConnectorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1beta1DataServicesConnectorsPostRequestBodyable, requestConfiguration *V1beta1DataServicesConnectorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1DataServicesConnectorsRequestBuilder when successful
func (m *V1beta1DataServicesConnectorsRequestBuilder) WithUrl(rawUrl string)(*V1beta1DataServicesConnectorsRequestBuilder) {
    return NewV1beta1DataServicesConnectorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
