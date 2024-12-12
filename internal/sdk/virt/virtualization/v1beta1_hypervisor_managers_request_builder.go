package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers
type V1beta1HypervisorManagersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HypervisorManagersPostRequestBody composed type wrapper for classes V1beta1HypervisorManagersPostRequestBodyMember1able, V1beta1HypervisorManagersPostRequestBodyMember2able
type HypervisorManagersPostRequestBody struct {
    // Composed type representation for type V1beta1HypervisorManagersPostRequestBodyMember1able
    v1beta1HypervisorManagersPostRequestBodyMember1 V1beta1HypervisorManagersPostRequestBodyMember1able
    // Composed type representation for type V1beta1HypervisorManagersPostRequestBodyMember2able
    v1beta1HypervisorManagersPostRequestBodyMember2 V1beta1HypervisorManagersPostRequestBodyMember2able
}
// NewHypervisorManagersPostRequestBody instantiates a new HypervisorManagersPostRequestBody and sets the default values.
func NewHypervisorManagersPostRequestBody()(*HypervisorManagersPostRequestBody) {
    m := &HypervisorManagersPostRequestBody{
    }
    return m
}
// CreateHypervisorManagersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHypervisorManagersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewHypervisorManagersPostRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HypervisorManagersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetV1beta1HypervisorManagersPostRequestBodyMember1() != nil {
        return m.GetV1beta1HypervisorManagersPostRequestBodyMember1().GetFieldDeserializers()
    } else if m.GetV1beta1HypervisorManagersPostRequestBodyMember2() != nil {
        return m.GetV1beta1HypervisorManagersPostRequestBodyMember2().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *HypervisorManagersPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1HypervisorManagersPostRequestBodyMember1 gets the V1beta1HypervisorManagersPostRequestBodyMember1 property value. Composed type representation for type V1beta1HypervisorManagersPostRequestBodyMember1able
// returns a V1beta1HypervisorManagersPostRequestBodyMember1able when successful
func (m *HypervisorManagersPostRequestBody) GetV1beta1HypervisorManagersPostRequestBodyMember1()(V1beta1HypervisorManagersPostRequestBodyMember1able) {
    return m.v1beta1HypervisorManagersPostRequestBodyMember1
}
// GetV1beta1HypervisorManagersPostRequestBodyMember2 gets the V1beta1HypervisorManagersPostRequestBodyMember2 property value. Composed type representation for type V1beta1HypervisorManagersPostRequestBodyMember2able
// returns a V1beta1HypervisorManagersPostRequestBodyMember2able when successful
func (m *HypervisorManagersPostRequestBody) GetV1beta1HypervisorManagersPostRequestBodyMember2()(V1beta1HypervisorManagersPostRequestBodyMember2able) {
    return m.v1beta1HypervisorManagersPostRequestBodyMember2
}
// Serialize serializes information the current object
func (m *HypervisorManagersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1HypervisorManagersPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1HypervisorManagersPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1HypervisorManagersPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1HypervisorManagersPostRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1HypervisorManagersPostRequestBodyMember1 sets the V1beta1HypervisorManagersPostRequestBodyMember1 property value. Composed type representation for type V1beta1HypervisorManagersPostRequestBodyMember1able
func (m *HypervisorManagersPostRequestBody) SetV1beta1HypervisorManagersPostRequestBodyMember1(value V1beta1HypervisorManagersPostRequestBodyMember1able)() {
    m.v1beta1HypervisorManagersPostRequestBodyMember1 = value
}
// SetV1beta1HypervisorManagersPostRequestBodyMember2 sets the V1beta1HypervisorManagersPostRequestBodyMember2 property value. Composed type representation for type V1beta1HypervisorManagersPostRequestBodyMember2able
func (m *HypervisorManagersPostRequestBody) SetV1beta1HypervisorManagersPostRequestBodyMember2(value V1beta1HypervisorManagersPostRequestBodyMember2able)() {
    m.v1beta1HypervisorManagersPostRequestBodyMember2 = value
}
// V1beta1HypervisorManagersRequestBuilderGetQueryParameters list all the registered hypervisor managers.
type V1beta1HypervisorManagersRequestBuilderGetQueryParameters struct {
    // The filter query parameter is used to filter the set of resources returned in the response.The returned set of resources must match the criteria in the filter query parameter.A comparison compares a property name to a literal. The comparisons supported are the following:* “eq” : Is a property equal to value. Valid for number, boolean and string properties.* “ne” : Is a property not equal to value. Valid for number, boolean and string properties.* “gt” : Is a property greater than a value. Valid for number or string timestamp properties.* “lt” : Is a property less than a value. Valid for number or string timestamp properties* “in” : Is a value in a property (that is an array of strings)Filters are supported on the following attributes:* hypervisorManagerType* appType* state* status* releaseVersion* createdAt* name* services* dataOrchestratorInfo/id* username* networkAddress* displayName
    Filter *string `uriparametername:"filter"`
    // The numbers of items to return
    Limit *int32 `uriparametername:"limit"`
    // The number of items to skip before starting to collect the result set
    Offset *int32 `uriparametername:"offset"`
    // The select query parameter is used to limit the properties returned with a resource or collection-level GET.Multiple properties can be listed to be returned. The server must only return the set of properties requested by the client.The property “select” is the name of the select query parameter; its value is the list of properties to return separated by commas.
    Select *string `uriparametername:"select"`
    // A comma separated list of properties to sort by, followed by a directionindicator ("asc" or "desc"). If no direction indicator is specified, thedefault order is ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1HypervisorManagersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1HypervisorManagersRequestBuilderGetQueryParameters
}
// V1beta1HypervisorManagersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
type HypervisorManagersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1HypervisorManagersPostRequestBodyMember1()(V1beta1HypervisorManagersPostRequestBodyMember1able)
    GetV1beta1HypervisorManagersPostRequestBodyMember2()(V1beta1HypervisorManagersPostRequestBodyMember2able)
    SetV1beta1HypervisorManagersPostRequestBodyMember1(value V1beta1HypervisorManagersPostRequestBodyMember1able)()
    SetV1beta1HypervisorManagersPostRequestBodyMember2(value V1beta1HypervisorManagersPostRequestBodyMember2able)()
}
// ByHypervisorId gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/virt.virtualization.v1beta1.hypervisorManagers.item collection
// returns a *V1beta1HypervisorManagersHypervisorItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersRequestBuilder) ByHypervisorId(hypervisorId string)(*V1beta1HypervisorManagersHypervisorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hypervisorId != "" {
        urlTplParams["hypervisor%2Did"] = hypervisorId
    }
    return NewV1beta1HypervisorManagersHypervisorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1HypervisorManagersRequestBuilderInternal instantiates a new V1beta1HypervisorManagersRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersRequestBuilder) {
    m := &V1beta1HypervisorManagersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers{?filter*,limit*,offset*,select*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersRequestBuilder instantiates a new V1beta1HypervisorManagersRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the registered hypervisor managers.
// Deprecated: This method is obsolete. Use GetAsHypervisorManagersGetResponse instead.
// returns a V1beta1HypervisorManagersResponseable when successful
// returns a V1beta1HypervisorManagers400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagers401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagers403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagers500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagers503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagers403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagers500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagers503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersResponseable), nil
}
// GetAsHypervisorManagersGetResponse list all the registered hypervisor managers.
// returns a V1beta1HypervisorManagersGetResponseable when successful
// returns a V1beta1HypervisorManagers400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagers401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagers403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagers500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagers503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersRequestBuilder) GetAsHypervisorManagersGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagers403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagers500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagers503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersGetResponseable), nil
}
// Post this API is used to register a new Hypervisor Manager from Data Services Connector with the DSCC.
// returns a UntypedNodeable when successful
// returns a V1beta1HypervisorManagers400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagers401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagers403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagers500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagers503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersRequestBuilder) Post(ctx context.Context, body HypervisorManagersPostRequestBodyable, requestConfiguration *V1beta1HypervisorManagersRequestBuilderPostRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagers400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagers401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagers403ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagers500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagers503ErrorFromDiscriminatorValue,
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
// ToGetRequestInformation list all the registered hypervisor managers.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation this API is used to register a new Hypervisor Manager from Data Services Connector with the DSCC.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersRequestBuilder) ToPostRequestInformation(ctx context.Context, body HypervisorManagersPostRequestBodyable, requestConfiguration *V1beta1HypervisorManagersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1HypervisorManagersRequestBuilder when successful
func (m *V1beta1HypervisorManagersRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersRequestBuilder) {
    return NewV1beta1HypervisorManagersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
