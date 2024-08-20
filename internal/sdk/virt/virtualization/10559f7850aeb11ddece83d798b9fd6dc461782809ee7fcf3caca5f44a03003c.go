package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse represents a single instance of a hypervisor library image
type V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time at which the virtual machine image was created
    createdAt *string
    // The customer application identifier.
    customerId *string
    // Description given for the virtual machine image from the hypervisor image library
    description *string
    // Name of the virtual machine image
    displayName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // The hypervisorManagerInfo property
    hypervisorManagerInfo V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable
    // UUID string uniquely identifying the virtual machine image from the hypervisor image library
    id *string
    // Name of the hypervisor library
    libraryName *string
    // Virtual machine image name
    name *string
    // The 'self' reference for this resource.
    resourceUri *string
    // Size of the virtual machine image from the hypervisor manager image library in bytes
    size *int32
    // True if the image will be pulled from a remote hypervisor image library (ex. vCenter subscribed content library).
    subscribed *bool
    // Hypervisor object type
    typeEscaped *string
    // Hypervisor provided identifier of the virtual machine image from the hypervisor image library
    uid *string
    // Time at which the virtual machine image was updated
    updatedAt *string
}
// NewV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse instantiates a new V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse and sets the default values.
func NewV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse()(*V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) {
    m := &V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the createdAt property value. Time at which the virtual machine image was created
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetCreatedAt()(*string) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDescription gets the description property value. Description given for the virtual machine image from the hypervisor image library
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the virtual machine image
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["hypervisorManagerInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHypervisorManagerInfo(val.(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["libraryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLibraryName(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["subscribed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscribed(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetHypervisorManagerInfo gets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
// returns a V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetHypervisorManagerInfo()(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable) {
    return m.hypervisorManagerInfo
}
// GetId gets the id property value. UUID string uniquely identifying the virtual machine image from the hypervisor image library
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetId()(*string) {
    return m.id
}
// GetLibraryName gets the libraryName property value. Name of the hypervisor library
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetLibraryName()(*string) {
    return m.libraryName
}
// GetName gets the name property value. Virtual machine image name
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSize gets the size property value. Size of the virtual machine image from the hypervisor manager image library in bytes
// returns a *int32 when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetSize()(*int32) {
    return m.size
}
// GetSubscribed gets the subscribed property value. True if the image will be pulled from a remote hypervisor image library (ex. vCenter subscribed content library).
// returns a *bool when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetSubscribed()(*bool) {
    return m.subscribed
}
// GetTypeEscaped gets the type property value. Hypervisor object type
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUid gets the uid property value. Hypervisor provided identifier of the virtual machine image from the hypervisor image library
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. Time at which the virtual machine image was updated
// returns a *string when successful
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) GetUpdatedAt()(*string) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hypervisorManagerInfo", m.GetHypervisorManagerInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("libraryName", m.GetLibraryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("subscribed", m.GetSubscribed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uid", m.GetUid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updatedAt", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the createdAt property value. Time at which the virtual machine image was created
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetCreatedAt(value *string)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier.
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDescription sets the description property value. Description given for the virtual machine image from the hypervisor image library
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the virtual machine image
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetHypervisorManagerInfo sets the hypervisorManagerInfo property value. The hypervisorManagerInfo property
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetHypervisorManagerInfo(value V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable)() {
    m.hypervisorManagerInfo = value
}
// SetId sets the id property value. UUID string uniquely identifying the virtual machine image from the hypervisor image library
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetId(value *string)() {
    m.id = value
}
// SetLibraryName sets the libraryName property value. Name of the hypervisor library
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetLibraryName(value *string)() {
    m.libraryName = value
}
// SetName sets the name property value. Virtual machine image name
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSize sets the size property value. Size of the virtual machine image from the hypervisor manager image library in bytes
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetSize(value *int32)() {
    m.size = value
}
// SetSubscribed sets the subscribed property value. True if the image will be pulled from a remote hypervisor image library (ex. vCenter subscribed content library).
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetSubscribed(value *bool)() {
    m.subscribed = value
}
// SetTypeEscaped sets the type property value. Hypervisor object type
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUid sets the uid property value. Hypervisor provided identifier of the virtual machine image from the hypervisor image library
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. Time at which the virtual machine image was updated
func (m *V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse) SetUpdatedAt(value *string)() {
    m.updatedAt = value
}
type V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetCustomerId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetGeneration()(*int64)
    GetHypervisorManagerInfo()(V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable)
    GetId()(*string)
    GetLibraryName()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetSize()(*int32)
    GetSubscribed()(*bool)
    GetTypeEscaped()(*string)
    GetUid()(*string)
    GetUpdatedAt()(*string)
    SetCreatedAt(value *string)()
    SetCustomerId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetGeneration(value *int64)()
    SetHypervisorManagerInfo(value V1beta1HypervisorManagersItemHypervisorLibraryImagesItemHypervisorLibraryImageGetResponse_hypervisorManagerInfoable)()
    SetId(value *string)()
    SetLibraryName(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetSize(value *int32)()
    SetSubscribed(value *bool)()
    SetTypeEscaped(value *string)()
    SetUid(value *string)()
    SetUpdatedAt(value *string)()
}
