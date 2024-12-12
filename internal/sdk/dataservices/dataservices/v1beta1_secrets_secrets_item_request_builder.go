package dataservices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SecretsSecretsItemRequestBuilder builds and executes requests for operations under \data-services\v1beta1\secrets\{id}
type V1beta1SecretsSecretsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecretsPatchRequestBody composed type wrapper for classes V1beta1SecretsItemSecretsPatchRequestBodyMember1able, V1beta1SecretsItemSecretsPatchRequestBodyMember2able
type SecretsPatchRequestBody struct {
    // Composed type representation for type V1beta1SecretsItemSecretsPatchRequestBodyMember1able
    v1beta1SecretsItemSecretsPatchRequestBodyMember1 V1beta1SecretsItemSecretsPatchRequestBodyMember1able
    // Composed type representation for type V1beta1SecretsItemSecretsPatchRequestBodyMember2able
    v1beta1SecretsItemSecretsPatchRequestBodyMember2 V1beta1SecretsItemSecretsPatchRequestBodyMember2able
}
// NewSecretsPatchRequestBody instantiates a new SecretsPatchRequestBody and sets the default values.
func NewSecretsPatchRequestBody()(*SecretsPatchRequestBody) {
    m := &SecretsPatchRequestBody{
    }
    return m
}
// CreateSecretsPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretsPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSecretsPatchRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchRequestBodyMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsItemSecretsPatchRequestBodyMember1able); ok {
                result.SetV1beta1SecretsItemSecretsPatchRequestBodyMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateV1beta1SecretsItemSecretsPatchRequestBodyMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(V1beta1SecretsItemSecretsPatchRequestBodyMember2able); ok {
                result.SetV1beta1SecretsItemSecretsPatchRequestBodyMember2(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretsPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *SecretsPatchRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1SecretsItemSecretsPatchRequestBodyMember1 gets the V1beta1SecretsItemSecretsPatchRequestBodyMember1 property value. Composed type representation for type V1beta1SecretsItemSecretsPatchRequestBodyMember1able
// returns a V1beta1SecretsItemSecretsPatchRequestBodyMember1able when successful
func (m *SecretsPatchRequestBody) GetV1beta1SecretsItemSecretsPatchRequestBodyMember1()(V1beta1SecretsItemSecretsPatchRequestBodyMember1able) {
    return m.v1beta1SecretsItemSecretsPatchRequestBodyMember1
}
// GetV1beta1SecretsItemSecretsPatchRequestBodyMember2 gets the V1beta1SecretsItemSecretsPatchRequestBodyMember2 property value. Composed type representation for type V1beta1SecretsItemSecretsPatchRequestBodyMember2able
// returns a V1beta1SecretsItemSecretsPatchRequestBodyMember2able when successful
func (m *SecretsPatchRequestBody) GetV1beta1SecretsItemSecretsPatchRequestBodyMember2()(V1beta1SecretsItemSecretsPatchRequestBodyMember2able) {
    return m.v1beta1SecretsItemSecretsPatchRequestBodyMember2
}
// Serialize serializes information the current object
func (m *SecretsPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1SecretsItemSecretsPatchRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsItemSecretsPatchRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1SecretsItemSecretsPatchRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1SecretsItemSecretsPatchRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1SecretsItemSecretsPatchRequestBodyMember1 sets the V1beta1SecretsItemSecretsPatchRequestBodyMember1 property value. Composed type representation for type V1beta1SecretsItemSecretsPatchRequestBodyMember1able
func (m *SecretsPatchRequestBody) SetV1beta1SecretsItemSecretsPatchRequestBodyMember1(value V1beta1SecretsItemSecretsPatchRequestBodyMember1able)() {
    m.v1beta1SecretsItemSecretsPatchRequestBodyMember1 = value
}
// SetV1beta1SecretsItemSecretsPatchRequestBodyMember2 sets the V1beta1SecretsItemSecretsPatchRequestBodyMember2 property value. Composed type representation for type V1beta1SecretsItemSecretsPatchRequestBodyMember2able
func (m *SecretsPatchRequestBody) SetV1beta1SecretsItemSecretsPatchRequestBodyMember2(value V1beta1SecretsItemSecretsPatchRequestBodyMember2able)() {
    m.v1beta1SecretsItemSecretsPatchRequestBodyMember2 = value
}
// V1beta1SecretsSecretsItemRequestBuilderDeleteQueryParameters removes the specified secret. All associated assignments will also be removed.
type V1beta1SecretsSecretsItemRequestBuilderDeleteQueryParameters struct {
    // Enable delete-lock safety checking
    Safe *bool `uriparametername:"safe"`
}
// V1beta1SecretsSecretsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretsSecretsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *V1beta1SecretsSecretsItemRequestBuilderDeleteQueryParameters
}
// V1beta1SecretsSecretsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretsSecretsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1SecretsSecretsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1SecretsSecretsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
type SecretsPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1SecretsItemSecretsPatchRequestBodyMember1()(V1beta1SecretsItemSecretsPatchRequestBodyMember1able)
    GetV1beta1SecretsItemSecretsPatchRequestBodyMember2()(V1beta1SecretsItemSecretsPatchRequestBodyMember2able)
    SetV1beta1SecretsItemSecretsPatchRequestBodyMember1(value V1beta1SecretsItemSecretsPatchRequestBodyMember1able)()
    SetV1beta1SecretsItemSecretsPatchRequestBodyMember2(value V1beta1SecretsItemSecretsPatchRequestBodyMember2able)()
}
// NewV1beta1SecretsSecretsItemRequestBuilderInternal instantiates a new V1beta1SecretsSecretsItemRequestBuilder and sets the default values.
func NewV1beta1SecretsSecretsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretsSecretsItemRequestBuilder) {
    m := &V1beta1SecretsSecretsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/data-services/v1beta1/secrets/{id}{?safe*}", pathParameters),
    }
    return m
}
// NewV1beta1SecretsSecretsItemRequestBuilder instantiates a new V1beta1SecretsSecretsItemRequestBuilder and sets the default values.
func NewV1beta1SecretsSecretsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1SecretsSecretsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1SecretsSecretsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes the specified secret. All associated assignments will also be removed.
// returns a V1beta1SecretsItemSecrets400Error error when the service returns a 400 status code
// returns a V1beta1SecretsItemSecrets401Error error when the service returns a 401 status code
// returns a V1beta1SecretsItemSecrets403Error error when the service returns a 403 status code
// returns a V1beta1SecretsItemSecrets404Error error when the service returns a 404 status code
// returns a V1beta1SecretsItemSecrets409Error error when the service returns a 409 status code
// returns a V1beta1SecretsItemSecrets412Error error when the service returns a 412 status code
// returns a V1beta1SecretsItemSecrets422Error error when the service returns a 422 status code
// returns a V1beta1SecretsItemSecrets500Error error when the service returns a 500 status code
// returns a V1beta1SecretsItemSecrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsSecretsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretsItemSecrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretsItemSecrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretsItemSecrets404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SecretsItemSecrets409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1SecretsItemSecrets412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretsItemSecrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretsItemSecrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretsItemSecrets503ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get reports the attributes and properties of the specified secret.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// returns a V1beta1SecretsItemSecretsResponseable when successful
// returns a V1beta1SecretsItemSecrets400Error error when the service returns a 400 status code
// returns a V1beta1SecretsItemSecrets401Error error when the service returns a 401 status code
// returns a V1beta1SecretsItemSecrets403Error error when the service returns a 403 status code
// returns a V1beta1SecretsItemSecrets404Error error when the service returns a 404 status code
// returns a V1beta1SecretsItemSecrets422Error error when the service returns a 422 status code
// returns a V1beta1SecretsItemSecrets500Error error when the service returns a 500 status code
// returns a V1beta1SecretsItemSecrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsSecretsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderGetRequestConfiguration)(V1beta1SecretsItemSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretsItemSecrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretsItemSecrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretsItemSecrets404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretsItemSecrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretsItemSecrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretsItemSecrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsItemSecretsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsItemSecretsResponseable), nil
}
// GetAsSecretsGetResponse reports the attributes and properties of the specified secret.
// returns a V1beta1SecretsItemSecretsGetResponseable when successful
// returns a V1beta1SecretsItemSecrets400Error error when the service returns a 400 status code
// returns a V1beta1SecretsItemSecrets401Error error when the service returns a 401 status code
// returns a V1beta1SecretsItemSecrets403Error error when the service returns a 403 status code
// returns a V1beta1SecretsItemSecrets404Error error when the service returns a 404 status code
// returns a V1beta1SecretsItemSecrets422Error error when the service returns a 422 status code
// returns a V1beta1SecretsItemSecrets500Error error when the service returns a 500 status code
// returns a V1beta1SecretsItemSecrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsSecretsItemRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderGetRequestConfiguration)(V1beta1SecretsItemSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretsItemSecrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretsItemSecrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretsItemSecrets404ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretsItemSecrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretsItemSecrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretsItemSecrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsItemSecretsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsItemSecretsGetResponseable), nil
}
// Patch changes an existing secret using the provided redefinition.
// Deprecated: This method is obsolete. Use PatchAsSecretsPatchResponse instead.
// returns a V1beta1SecretsItemSecretsResponseable when successful
// returns a V1beta1SecretsItemSecrets400Error error when the service returns a 400 status code
// returns a V1beta1SecretsItemSecrets401Error error when the service returns a 401 status code
// returns a V1beta1SecretsItemSecrets403Error error when the service returns a 403 status code
// returns a V1beta1SecretsItemSecrets404Error error when the service returns a 404 status code
// returns a V1beta1SecretsItemSecrets409Error error when the service returns a 409 status code
// returns a V1beta1SecretsItemSecrets412Error error when the service returns a 412 status code
// returns a V1beta1SecretsItemSecrets422Error error when the service returns a 422 status code
// returns a V1beta1SecretsItemSecrets500Error error when the service returns a 500 status code
// returns a V1beta1SecretsItemSecrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsSecretsItemRequestBuilder) Patch(ctx context.Context, body SecretsPatchRequestBodyable, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderPatchRequestConfiguration)(V1beta1SecretsItemSecretsResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretsItemSecrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretsItemSecrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretsItemSecrets404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SecretsItemSecrets409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1SecretsItemSecrets412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretsItemSecrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretsItemSecrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretsItemSecrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsItemSecretsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsItemSecretsResponseable), nil
}
// PatchAsSecretsPatchResponse changes an existing secret using the provided redefinition.
// returns a V1beta1SecretsItemSecretsPatchResponseable when successful
// returns a V1beta1SecretsItemSecrets400Error error when the service returns a 400 status code
// returns a V1beta1SecretsItemSecrets401Error error when the service returns a 401 status code
// returns a V1beta1SecretsItemSecrets403Error error when the service returns a 403 status code
// returns a V1beta1SecretsItemSecrets404Error error when the service returns a 404 status code
// returns a V1beta1SecretsItemSecrets409Error error when the service returns a 409 status code
// returns a V1beta1SecretsItemSecrets412Error error when the service returns a 412 status code
// returns a V1beta1SecretsItemSecrets422Error error when the service returns a 422 status code
// returns a V1beta1SecretsItemSecrets500Error error when the service returns a 500 status code
// returns a V1beta1SecretsItemSecrets503Error error when the service returns a 503 status code
func (m *V1beta1SecretsSecretsItemRequestBuilder) PatchAsSecretsPatchResponse(ctx context.Context, body SecretsPatchRequestBodyable, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderPatchRequestConfiguration)(V1beta1SecretsItemSecretsPatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1SecretsItemSecrets401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1SecretsItemSecrets403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1SecretsItemSecrets404ErrorFromDiscriminatorValue,
        "409": CreateV1beta1SecretsItemSecrets409ErrorFromDiscriminatorValue,
        "412": CreateV1beta1SecretsItemSecrets412ErrorFromDiscriminatorValue,
        "422": CreateV1beta1SecretsItemSecrets422ErrorFromDiscriminatorValue,
        "500": CreateV1beta1SecretsItemSecrets500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1SecretsItemSecrets503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1SecretsItemSecretsPatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1SecretsItemSecretsPatchResponseable), nil
}
// ToDeleteRequestInformation removes the specified secret. All associated assignments will also be removed.
// returns a *RequestInformation when successful
func (m *V1beta1SecretsSecretsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation reports the attributes and properties of the specified secret.
// returns a *RequestInformation when successful
func (m *V1beta1SecretsSecretsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation changes an existing secret using the provided redefinition.
// returns a *RequestInformation when successful
func (m *V1beta1SecretsSecretsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body SecretsPatchRequestBodyable, requestConfiguration *V1beta1SecretsSecretsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *V1beta1SecretsSecretsItemRequestBuilder when successful
func (m *V1beta1SecretsSecretsItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1SecretsSecretsItemRequestBuilder) {
    return NewV1beta1SecretsSecretsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
