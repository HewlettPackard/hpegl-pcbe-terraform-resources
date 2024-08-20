package virtualization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersHypervisorItemRequestBuilder builds and executes requests for operations under \virtualization\v1beta1\hypervisor-managers\{hypervisor-id}
type V1beta1HypervisorManagersHypervisorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HypervisorPatchRequestBody composed type wrapper for classes V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able, V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
type HypervisorPatchRequestBody struct {
    // Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
    hypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
    // Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
    hypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
    // Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
    v1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
    // Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
    v1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
}
// NewHypervisorPatchRequestBody instantiates a new HypervisorPatchRequestBody and sets the default values.
func NewHypervisorPatchRequestBody()(*HypervisorPatchRequestBody) {
    m := &HypervisorPatchRequestBody{
    }
    return m
}
// CreateHypervisorPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHypervisorPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewHypervisorPatchRequestBody()
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
func (m *HypervisorPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 gets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
// returns a V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able when successful
func (m *HypervisorPatchRequestBody) GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able) {
    return m.hypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1
}
// GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 gets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
// returns a V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able when successful
func (m *HypervisorPatchRequestBody) GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able) {
    return m.hypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *HypervisorPatchRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 gets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
// returns a V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able when successful
func (m *HypervisorPatchRequestBody) GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able) {
    return m.v1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1
}
// GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 gets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
// returns a V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able when successful
func (m *HypervisorPatchRequestBody) GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able) {
    return m.v1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2
}
// Serialize serializes information the current object
func (m *HypervisorPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 sets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
func (m *HypervisorPatchRequestBody) SetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able)() {
    m.hypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 = value
}
// SetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 sets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
func (m *HypervisorPatchRequestBody) SetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able)() {
    m.hypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 = value
}
// SetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 sets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able
func (m *HypervisorPatchRequestBody) SetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able)() {
    m.v1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1 = value
}
// SetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 sets the V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 property value. Composed type representation for type V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able
func (m *HypervisorPatchRequestBody) SetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able)() {
    m.v1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2 = value
}
// V1beta1HypervisorManagersHypervisorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersHypervisorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1HypervisorManagersHypervisorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersHypervisorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// V1beta1HypervisorManagersHypervisorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1beta1HypervisorManagersHypervisorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
type HypervisorPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able)
    GetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able)
    GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able)
    GetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2()(V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able)
    SetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able)()
    SetHypervisorPatchRequestBodyV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able)()
    SetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember1able)()
    SetV1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2(value V1beta1HypervisorManagersItemHypervisorPatchRequestBodyMember2able)()
}
// NewV1beta1HypervisorManagersHypervisorItemRequestBuilderInternal instantiates a new V1beta1HypervisorManagersHypervisorItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersHypervisorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersHypervisorItemRequestBuilder) {
    m := &V1beta1HypervisorManagersHypervisorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/virtualization/v1beta1/hypervisor-managers/{hypervisor%2Did}", pathParameters),
    }
    return m
}
// NewV1beta1HypervisorManagersHypervisorItemRequestBuilder instantiates a new V1beta1HypervisorManagersHypervisorItemRequestBuilder and sets the default values.
func NewV1beta1HypervisorManagersHypervisorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1beta1HypervisorManagersHypervisorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1beta1HypervisorManagersHypervisorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete this API will unregister the hypervisor manager from Data Services Connector and Data Services Cloud Console.
// returns a UntypedNodeable when successful
// returns a V1beta1HypervisorManagersItemHypervisor400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemHypervisor401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisor403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisor404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemHypervisor500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagersItemHypervisor503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderDeleteRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemHypervisor400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemHypervisor401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisor403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemHypervisor404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisor500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagersItemHypervisor503ErrorFromDiscriminatorValue,
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
// Folders the folders property
// returns a *V1beta1HypervisorManagersItemFoldersRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Folders()(*V1beta1HypervisorManagersItemFoldersRequestBuilder) {
    return NewV1beta1HypervisorManagersItemFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get detailed information for a registered hypervisor manager qualified by hypervisor-id.
// Deprecated: This method is obsolete. Use GetAsHypervisorGetResponse instead.
// returns a V1beta1HypervisorManagersItemHypervisorResponseable when successful
// returns a V1beta1HypervisorManagersItemHypervisor401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisor403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisor404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemHypervisor500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagersItemHypervisor503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemHypervisorResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemHypervisor401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisor403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemHypervisor404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisor500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagersItemHypervisor503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemHypervisorResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemHypervisorResponseable), nil
}
// GetAsHypervisorGetResponse get detailed information for a registered hypervisor manager qualified by hypervisor-id.
// returns a V1beta1HypervisorManagersItemHypervisorGetResponseable when successful
// returns a V1beta1HypervisorManagersItemHypervisor401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisor403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisor404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemHypervisor500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagersItemHypervisor503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) GetAsHypervisorGetResponse(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderGetRequestConfiguration)(V1beta1HypervisorManagersItemHypervisorGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": CreateV1beta1HypervisorManagersItemHypervisor401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisor403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemHypervisor404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisor500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagersItemHypervisor503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1beta1HypervisorManagersItemHypervisorGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1beta1HypervisorManagersItemHypervisorGetResponseable), nil
}
// HypervisorLibraryImages the hypervisorLibraryImages property
// returns a *V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) HypervisorLibraryImages()(*V1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilder) {
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
// returns a *V1beta1HypervisorManagersItemNetworksRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Networks()(*V1beta1HypervisorManagersItemNetworksRequestBuilder) {
    return NewV1beta1HypervisorManagersItemNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update attributes of the specified Hypervisor Manager from Data Services Connector
// returns a UntypedNodeable when successful
// returns a V1beta1HypervisorManagersItemHypervisor400Error error when the service returns a 400 status code
// returns a V1beta1HypervisorManagersItemHypervisor401Error error when the service returns a 401 status code
// returns a V1beta1HypervisorManagersItemHypervisor403Error error when the service returns a 403 status code
// returns a V1beta1HypervisorManagersItemHypervisor404Error error when the service returns a 404 status code
// returns a V1beta1HypervisorManagersItemHypervisor500Error error when the service returns a 500 status code
// returns a V1beta1HypervisorManagersItemHypervisor503Error error when the service returns a 503 status code
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Patch(ctx context.Context, body HypervisorPatchRequestBodyable, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderPatchRequestConfiguration)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateV1beta1HypervisorManagersItemHypervisor400ErrorFromDiscriminatorValue,
        "401": CreateV1beta1HypervisorManagersItemHypervisor401ErrorFromDiscriminatorValue,
        "403": CreateV1beta1HypervisorManagersItemHypervisor403ErrorFromDiscriminatorValue,
        "404": CreateV1beta1HypervisorManagersItemHypervisor404ErrorFromDiscriminatorValue,
        "500": CreateV1beta1HypervisorManagersItemHypervisor500ErrorFromDiscriminatorValue,
        "503": CreateV1beta1HypervisorManagersItemHypervisor503ErrorFromDiscriminatorValue,
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
// Refresh the refresh property
// returns a *V1beta1HypervisorManagersItemRefreshRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Refresh()(*V1beta1HypervisorManagersItemRefreshRequestBuilder) {
    return NewV1beta1HypervisorManagersItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourcePools the resourcePools property
// returns a *V1beta1HypervisorManagersItemResourcePoolsRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) ResourcePools()(*V1beta1HypervisorManagersItemResourcePoolsRequestBuilder) {
    return NewV1beta1HypervisorManagersItemResourcePoolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags the tags property
// returns a *V1beta1HypervisorManagersItemTagsRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) Tags()(*V1beta1HypervisorManagersItemTagsRequestBuilder) {
    return NewV1beta1HypervisorManagersItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation this API will unregister the hypervisor manager from Data Services Connector and Data Services Cloud Console.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get detailed information for a registered hypervisor manager qualified by hypervisor-id.
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update attributes of the specified Hypervisor Manager from Data Services Connector
// returns a *RequestInformation when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body HypervisorPatchRequestBodyable, requestConfiguration *V1beta1HypervisorManagersHypervisorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1beta1HypervisorManagersHypervisorItemRequestBuilder when successful
func (m *V1beta1HypervisorManagersHypervisorItemRequestBuilder) WithUrl(rawUrl string)(*V1beta1HypervisorManagersHypervisorItemRequestBuilder) {
    return NewV1beta1HypervisorManagersHypervisorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
