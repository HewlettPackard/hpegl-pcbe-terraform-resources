package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies if the network has to be connected at power on.
    connectAtPowerOn *bool
    // A user-friendly name that identifies the hypervisor network. This will always be same as name since adding or updating hypervisor networks is not supported when managed from a manager, such as the vCenter.
    displayName *string
    // UUID string uniquely identifying the hypervisor network resource.
    id *string
    // Name of the network as reported by the hypervisor manager.
    name *string
    // The URI reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails instantiates a new V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails and sets the default values.
func NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails()(*V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) {
    m := &V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConnectAtPowerOn gets the connectAtPowerOn property value. Specifies if the network has to be connected at power on.
// returns a *bool when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetConnectAtPowerOn()(*bool) {
    return m.connectAtPowerOn
}
// GetDisplayName gets the displayName property value. A user-friendly name that identifies the hypervisor network. This will always be same as name since adding or updating hypervisor networks is not supported when managed from a manager, such as the vCenter.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connectAtPowerOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectAtPowerOn(val)
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
    return res
}
// GetId gets the id property value. UUID string uniquely identifying the hypervisor network resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the network as reported by the hypervisor manager.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The URI reference for this resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("connectAtPowerOn", m.GetConnectAtPowerOn())
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
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConnectAtPowerOn sets the connectAtPowerOn property value. Specifies if the network has to be connected at power on.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetConnectAtPowerOn(value *bool)() {
    m.connectAtPowerOn = value
}
// SetDisplayName sets the displayName property value. A user-friendly name that identifies the hypervisor network. This will always be same as name since adding or updating hypervisor networks is not supported when managed from a manager, such as the vCenter.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. UUID string uniquely identifying the hypervisor network resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the network as reported by the hypervisor manager.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The URI reference for this resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetails) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1VirtualMachinesItemVirtualMachinesGetResponse_networkAdapters_networkDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAtPowerOn()(*bool)
    GetDisplayName()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetConnectAtPowerOn(value *bool)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
