package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody customizes guest operating system attributes. Supported operating system are linux. Must be one of - linuxConfig, cloudInit.
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Cloud init configuration for the guest operating system
    cloudInit V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable
    // Configurations for linux operating system.
    linuxConfig V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable
    // Maximum interval time in minutes to wait for the cloud-init to be applied on the guest virtual machine. Maximum allowed time is 90 minutes. If the value is set to 0, cloudInit configurations are applied, but not verfied. If value is greater than 0, after cloudInit configurations are applied, the virtual machine is powered on and waits for completion for provided interval of time, to validate the status of the cloud-init configuration that was applied.
    waitIntervalInMinutes *int32
}
// NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody instantiates a new V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody and sets the default values.
func NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody()(*V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) {
    m := &V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudInit gets the cloudInit property value. Cloud init configuration for the guest operating system
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) GetCloudInit()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable) {
    return m.cloudInit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudInit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudInit(val.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable))
        }
        return nil
    }
    res["linuxConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinuxConfig(val.(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable))
        }
        return nil
    }
    res["waitIntervalInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitIntervalInMinutes(val)
        }
        return nil
    }
    return res
}
// GetLinuxConfig gets the linuxConfig property value. Configurations for linux operating system.
// returns a V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) GetLinuxConfig()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable) {
    return m.linuxConfig
}
// GetWaitIntervalInMinutes gets the waitIntervalInMinutes property value. Maximum interval time in minutes to wait for the cloud-init to be applied on the guest virtual machine. Maximum allowed time is 90 minutes. If the value is set to 0, cloudInit configurations are applied, but not verfied. If value is greater than 0, after cloudInit configurations are applied, the virtual machine is powered on and waits for completion for provided interval of time, to validate the status of the cloud-init configuration that was applied.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) GetWaitIntervalInMinutes()(*int32) {
    return m.waitIntervalInMinutes
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cloudInit", m.GetCloudInit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("linuxConfig", m.GetLinuxConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("waitIntervalInMinutes", m.GetWaitIntervalInMinutes())
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
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudInit sets the cloudInit property value. Cloud init configuration for the guest operating system
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) SetCloudInit(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable)() {
    m.cloudInit = value
}
// SetLinuxConfig sets the linuxConfig property value. Configurations for linux operating system.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) SetLinuxConfig(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable)() {
    m.linuxConfig = value
}
// SetWaitIntervalInMinutes sets the waitIntervalInMinutes property value. Maximum interval time in minutes to wait for the cloud-init to be applied on the guest virtual machine. Maximum allowed time is 90 minutes. If the value is set to 0, cloudInit configurations are applied, but not verfied. If value is greater than 0, after cloudInit configurations are applied, the virtual machine is powered on and waits for completion for provided interval of time, to validate the status of the cloud-init configuration that was applied.
func (m *V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody) SetWaitIntervalInMinutes(value *int32)() {
    m.waitIntervalInMinutes = value
}
type V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudInit()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable)
    GetLinuxConfig()(V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable)
    GetWaitIntervalInMinutes()(*int32)
    SetCloudInit(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_cloudInitable)()
    SetLinuxConfig(value V1beta1VirtualMachinesItemCustomizeGuestOsPostRequestBody_linuxConfigable)()
    SetWaitIntervalInMinutes(value *int32)()
}
