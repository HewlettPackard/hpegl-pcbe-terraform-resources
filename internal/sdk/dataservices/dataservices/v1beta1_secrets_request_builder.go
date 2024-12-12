package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsRequestBuilder builds and executes requests for operations under \data-services\v1beta1\secrets
type V1beta1SecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecretsPostRequestBody composed type wrapper for classes V1beta1SecretsPostRequestBodyMember1able, V1beta1SecretsPostRequestBodyMember2able
type SecretsPostRequestBody struct {
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember1able
    v1beta1SecretsPostRequestBodyMember1 V1beta1SecretsPostRequestBodyMember1able
    // Composed type representation for type V1beta1SecretsPostRequestBodyMember2able
    v1beta1SecretsPostRequestBodyMember2 V1beta1SecretsPostRequestBodyMember2able
}
// NewSecretsPostRequestBody instantiates a new SecretsPostRequestBody and sets the default values.
func NewSecretsPostRequestBody()(*SecretsPostRequestBody) {
    m := &SecretsPostRequestBody{
    }
    return m
}
// CreateSecretsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSecretsPostRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember1able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsPostRequestBodyMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsPostRequestBodyMember2able); ok {
                result.SetV1beta1SecretsPostRequestBodyMember2(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *SecretsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1SecretsPostRequestBodyMember1 gets the V1beta1SecretsPostRequestBodyMember1 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember1able
// returns a V1beta1SecretsPostRequestBodyMember1able when successful
func (m *SecretsPostRequestBody) GetV1beta1SecretsPostRequestBodyMember1()(V1beta1SecretsPostRequestBodyMember1able) {
    return m.v1beta1SecretsPostRequestBodyMember1
}
// GetV1beta1SecretsPostRequestBodyMember2 gets the V1beta1SecretsPostRequestBodyMember2 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2able
// returns a V1beta1SecretsPostRequestBodyMember2able when successful
func (m *SecretsPostRequestBody) GetV1beta1SecretsPostRequestBodyMember2()(V1beta1SecretsPostRequestBodyMember2able) {
    return m.v1beta1SecretsPostRequestBodyMember2
}
// Serialize serializes information the current object
func (m *SecretsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1SecretsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1SecretsPostRequestBodyMember1 sets the V1beta1SecretsPostRequestBodyMember1 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember1able
func (m *SecretsPostRequestBody) SetV1beta1SecretsPostRequestBodyMember1(value V1beta1SecretsPostRequestBodyMember1able)() {
    m.v1beta1SecretsPostRequestBodyMember1 = value
}
// SetV1beta1SecretsPostRequestBodyMember2 sets the V1beta1SecretsPostRequestBodyMember2 property value. Composed type representation for type V1beta1SecretsPostRequestBodyMember2able
func (m *SecretsPostRequestBody) SetV1beta1SecretsPostRequestBodyMember2(value V1beta1SecretsPostRequestBodyMember2able)() {
    m.v1beta1SecretsPostRequestBodyMember2 = value
}
// V1beta1SecretsRequestBuilderGetQueryParameters reports the attributes and properties of the secrets owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
type V1beta1SecretsRequestBuilderGetQueryParameters struct {
    // An OData expression to filter responses by attribute. The OData logical operator "eq" is case-sensitive and supported for attributes "classifier", "label", "name", "policy", "prototype", "service", "status", and "subclassifier". The OData function "contains()" is not case-sensitive and supported for attributes "label", "name", "policy", and "service". The OData logical operator "and" is supported for all attributes.
    Filter *string `uriparametername:"filter"`
    // The limit query parameter should be used in conjunction with offset for paging within a batched result set. The limit is the maximum number of items to include in the response. Example: offset=30&limit=10
    Limit *int32 `uriparametername:"limit"`
    // The offset query parameter should be used in conjunction with limit for paging within a batched result set. The offset is the number of items from the beginning of the batched result set to the first item included in the response. Example: offset=30&limit=10
    Offset *int32 `uriparametername:"offset"`
    // A response attribute to sort by, followed by a direction indicator ("asc" or "desc"). The attribute may be one of "assignmentsCount", "classifier", "createdAt", "createdBy", "id", "label", "lastUpdatedBy", "name", "policy", "prototype", "service", "status", "subclassifier", or "updatedAt". Default: ascending.
    Sort *string `uriparametername:"sort"`
}
// V1beta1SecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SecretsRequestBuilderGetQueryParameters
}
// V1beta1SecretsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
type SecretsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1SecretsPostRequestBodyMember1()(V1beta1SecretsPostRequestBodyMember1able)
    GetV1beta1SecretsPostRequestBodyMember2()(V1beta1SecretsPostRequestBodyMember2able)
    SetV1beta1SecretsPostRequestBodyMember1(value V1beta1SecretsPostRequestBodyMember1able)()
    SetV1beta1SecretsPostRequestBodyMember2(value V1beta1SecretsPostRequestBodyMember2able)()
}
// ById gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.secrets.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *V1beta1SecretsSecretsItemRequestBuilder when successful
func (m *V1beta1SecretsRequestBuilder) ById(id string)(*V1beta1SecretsSecretsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewV1beta1SecretsSecretsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/dataservices.dataServices.v1beta1.secrets.item collection
// returns a *V1beta1SecretsSecretsItemRequestBuilder when successful
func (m *V1beta1SecretsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*V1beta1SecretsSecretsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewV1beta1SecretsSecretsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1beta1SecretsRequestBuilderInternal instantiates a new V1beta1SecretsRequestBuilder and sets the default values.
func NewV1beta1SecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretsRequestBuilder) {
    m := &V1beta1SecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/secrets{?filter*,limit*,offset*,sort*}", pathParameters),
    }
    return m
}
// NewV1beta1SecretsRequestBuilder instantiates a new V1beta1SecretsRequestBuilder and sets the default values.
func NewV1beta1SecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get reports the attributes and properties of the secrets owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// returns a V1beta1SecretsResponseable when successful
// returns a V1beta1Secrets400Error error when the service returns a 400 status code
// returns a V1beta1Secrets401Error error when the service returns a 401 status code
// returns a V1beta1Secrets403Error error when the service returns a 403 status code
// returns a V1beta1Secrets404Error error when the service returns a 404 status code
// returns a V1beta1Secrets422Error error when the service returns a 422 status code
// returns a V1beta1Secrets500Error error when the service returns a 500 status code
// returns a V1beta1Secrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SecretsRequestBuilderGetRequestConfiguration)(V1beta1SecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Secrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Secrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Secrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Secrets404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Secrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Secrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Secrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsResponseable), nil
}
// GetAsSecretsGetResponse reports the attributes and properties of the secrets owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
// returns a V1beta1SecretsGetResponseable when successful
// returns a V1beta1Secrets400Error error when the service returns a 400 status code
// returns a V1beta1Secrets401Error error when the service returns a 401 status code
// returns a V1beta1Secrets403Error error when the service returns a 403 status code
// returns a V1beta1Secrets404Error error when the service returns a 404 status code
// returns a V1beta1Secrets422Error error when the service returns a 422 status code
// returns a V1beta1Secrets500Error error when the service returns a 500 status code
// returns a V1beta1Secrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *V1beta1SecretsRequestBuilderGetRequestConfiguration)(V1beta1SecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Secrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Secrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Secrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Secrets404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Secrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Secrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Secrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsGetResponseable), nil
}
// Post adds a new secret using the provided specification.
// Deprecated: This method is obsolete. Use PostAsSecretsPostResponse instead.
// returns a V1beta1SecretsResponseable when successful
// returns a V1beta1Secrets400Error error when the service returns a 400 status code
// returns a V1beta1Secrets401Error error when the service returns a 401 status code
// returns a V1beta1Secrets403Error error when the service returns a 403 status code
// returns a V1beta1Secrets404Error error when the service returns a 404 status code
// returns a V1beta1Secrets409Error error when the service returns a 409 status code
// returns a V1beta1Secrets412Error error when the service returns a 412 status code
// returns a V1beta1Secrets422Error error when the service returns a 422 status code
// returns a V1beta1Secrets500Error error when the service returns a 500 status code
// returns a V1beta1Secrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsRequestBuilder) Post(ctx context.Context, body SecretsPostRequestBodyable, requestConfiguration *V1beta1SecretsRequestBuilderPostRequestConfiguration)(V1beta1SecretsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Secrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Secrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Secrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Secrets404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1Secrets409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1Secrets412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Secrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Secrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Secrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsResponseable), nil
}
// PostAsSecretsPostResponse adds a new secret using the provided specification.
// returns a V1beta1SecretsPostResponseable when successful
// returns a V1beta1Secrets400Error error when the service returns a 400 status code
// returns a V1beta1Secrets401Error error when the service returns a 401 status code
// returns a V1beta1Secrets403Error error when the service returns a 403 status code
// returns a V1beta1Secrets404Error error when the service returns a 404 status code
// returns a V1beta1Secrets409Error error when the service returns a 409 status code
// returns a V1beta1Secrets412Error error when the service returns a 412 status code
// returns a V1beta1Secrets422Error error when the service returns a 422 status code
// returns a V1beta1Secrets500Error error when the service returns a 500 status code
// returns a V1beta1Secrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsRequestBuilder) PostAsSecretsPostResponse(ctx context.Context, body SecretsPostRequestBodyable, requestConfiguration *V1beta1SecretsRequestBuilderPostRequestConfiguration)(V1beta1SecretsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1Secrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1Secrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1Secrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1Secrets404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1Secrets409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1Secrets412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1Secrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1Secrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1Secrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsPostResponseable), nil
}
// ToGetRequestInformation reports the attributes and properties of the secrets owned by the customer. The response can be paged by using the limit and offset query parameters and filtered and sorted by using the filter and sort query parameters.
// returns a *RequestInformation when successful
func (m *V1beta1SecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation adds a new secret using the provided specification.
// returns a *RequestInformation when successful
func (m *V1beta1SecretsRequestBuilder) ToPostRequestInformation(ctx context.Context, body SecretsPostRequestBodyable, requestConfiguration *V1beta1SecretsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V1beta1SecretsRequestBuilder when successful
func (m *V1beta1SecretsRequestBuilder) WithUrl(rawUrl string)(*V1beta1SecretsRequestBuilder) {
    return NewV1beta1SecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
