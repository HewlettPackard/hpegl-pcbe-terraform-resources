package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the csp machine image resource representing the cloud provider machine image
    imageId *string
    // Password to login to the virtual machine. This is needed for windows based virtual machines.
    password *string
    // Region on which CSP machine instance will be created
    region *string
    // An identifier for the resource group, usually a UUID.
    resourceGroupId *string
    // An identifier for the subscription, usually a UUID.
    subscriptionId *string
    // Username to login to the virtual machine. This is needed for windows based virtual machines.
    username *string
    // virtual machine size
    vmSize *string
}
// NewV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 instantiates a new V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2 and sets the default values.
func NewV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2()(*V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) {
    m := &V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["imageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["resourceGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceGroupId(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    res["vmSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmSize(val)
        }
        return nil
    }
    return res
}
// GetImageId gets the imageId property value. ID of the csp machine image resource representing the cloud provider machine image
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetImageId()(*string) {
    return m.imageId
}
// GetPassword gets the password property value. Password to login to the virtual machine. This is needed for windows based virtual machines.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetPassword()(*string) {
    return m.password
}
// GetRegion gets the region property value. Region on which CSP machine instance will be created
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetRegion()(*string) {
    return m.region
}
// GetResourceGroupId gets the resourceGroupId property value. An identifier for the resource group, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetResourceGroupId()(*string) {
    return m.resourceGroupId
}
// GetSubscriptionId gets the subscriptionId property value. An identifier for the subscription, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// GetUsername gets the username property value. Username to login to the virtual machine. This is needed for windows based virtual machines.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetUsername()(*string) {
    return m.username
}
// GetVmSize gets the vmSize property value. virtual machine size
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) GetVmSize()(*string) {
    return m.vmSize
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("imageId", m.GetImageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceGroupId", m.GetResourceGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vmSize", m.GetVmSize())
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
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImageId sets the imageId property value. ID of the csp machine image resource representing the cloud provider machine image
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetImageId(value *string)() {
    m.imageId = value
}
// SetPassword sets the password property value. Password to login to the virtual machine. This is needed for windows based virtual machines.
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetPassword(value *string)() {
    m.password = value
}
// SetRegion sets the region property value. Region on which CSP machine instance will be created
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetRegion(value *string)() {
    m.region = value
}
// SetResourceGroupId sets the resourceGroupId property value. An identifier for the resource group, usually a UUID.
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetResourceGroupId(value *string)() {
    m.resourceGroupId = value
}
// SetSubscriptionId sets the subscriptionId property value. An identifier for the subscription, usually a UUID.
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
// SetUsername sets the username property value. Username to login to the virtual machine. This is needed for windows based virtual machines.
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetUsername(value *string)() {
    m.username = value
}
// SetVmSize sets the vmSize property value. virtual machine size
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2) SetVmSize(value *string)() {
    m.vmSize = value
}
type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImageId()(*string)
    GetPassword()(*string)
    GetRegion()(*string)
    GetResourceGroupId()(*string)
    GetSubscriptionId()(*string)
    GetUsername()(*string)
    GetVmSize()(*string)
    SetImageId(value *string)()
    SetPassword(value *string)()
    SetRegion(value *string)()
    SetResourceGroupId(value *string)()
    SetSubscriptionId(value *string)()
    SetUsername(value *string)()
    SetVmSize(value *string)()
}
