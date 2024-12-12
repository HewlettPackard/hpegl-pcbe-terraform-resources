package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit cloud init configuration for the guest operating system
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The metadata information which is needed to configure the virtual machine before booting of the operation system.
    metadata *string
    // User customized content that will be used by the cloud-init.
    userdata *string
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    res["userdata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserdata(val)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. The metadata information which is needed to configure the virtual machine before booting of the operation system.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) GetMetadata()(*string) {
    return m.metadata
}
// GetUserdata gets the userdata property value. User customized content that will be used by the cloud-init.
// returns a *string when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) GetUserdata()(*string) {
    return m.userdata
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userdata", m.GetUserdata())
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMetadata sets the metadata property value. The metadata information which is needed to configure the virtual machine before booting of the operation system.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) SetMetadata(value *string)() {
    m.metadata = value
}
// SetUserdata sets the userdata property value. User customized content that will be used by the cloud-init.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInit) SetUserdata(value *string)() {
    m.userdata = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMetadata()(*string)
    GetUserdata()(*string)
    SetMetadata(value *string)()
    SetUserdata(value *string)()
}
