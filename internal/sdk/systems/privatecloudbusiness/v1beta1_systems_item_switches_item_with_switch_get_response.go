package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemSwitchesItemWithSwitchGetResponse details of the Switch with select information.
type V1beta1SystemsItemSwitchesItemWithSwitchGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The switch firmware version.
    firmwareVersion *string
    // Switch Health Information
    health V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable
    // IP address of the switch.
    ipAddress *string
    // Rack name of the switch.
    rackName *string
    // Switch serial number.
    serialNumber *string
    // Unique Identifier of the system, usually a UUID.
    systemId *string
}
// NewV1beta1SystemsItemSwitchesItemWithSwitchGetResponse instantiates a new V1beta1SystemsItemSwitchesItemWithSwitchGetResponse and sets the default values.
func NewV1beta1SystemsItemSwitchesItemWithSwitchGetResponse()(*V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) {
    m := &V1beta1SystemsItemSwitchesItemWithSwitchGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemSwitchesItemWithSwitchGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSwitchesItemWithSwitchGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSwitchesItemWithSwitchGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["firmwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareVersion(val)
        }
        return nil
    }
    res["health"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealth(val.(V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable))
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["rackName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRackName(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["systemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemId(val)
        }
        return nil
    }
    return res
}
// GetFirmwareVersion gets the firmwareVersion property value. The switch firmware version.
// returns a *string when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetFirmwareVersion()(*string) {
    return m.firmwareVersion
}
// GetHealth gets the health property value. Switch Health Information
// returns a V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetHealth()(V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable) {
    return m.health
}
// GetIpAddress gets the ipAddress property value. IP address of the switch.
// returns a *string when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetRackName gets the rackName property value. Rack name of the switch.
// returns a *string when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetRackName()(*string) {
    return m.rackName
}
// GetSerialNumber gets the serialNumber property value. Switch serial number.
// returns a *string when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSystemId gets the systemId property value. Unique Identifier of the system, usually a UUID.
// returns a *string when successful
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) GetSystemId()(*string) {
    return m.systemId
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("firmwareVersion", m.GetFirmwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("health", m.GetHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rackName", m.GetRackName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("systemId", m.GetSystemId())
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
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFirmwareVersion sets the firmwareVersion property value. The switch firmware version.
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetFirmwareVersion(value *string)() {
    m.firmwareVersion = value
}
// SetHealth sets the health property value. Switch Health Information
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetHealth(value V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable)() {
    m.health = value
}
// SetIpAddress sets the ipAddress property value. IP address of the switch.
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetRackName sets the rackName property value. Rack name of the switch.
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetRackName(value *string)() {
    m.rackName = value
}
// SetSerialNumber sets the serialNumber property value. Switch serial number.
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSystemId sets the systemId property value. Unique Identifier of the system, usually a UUID.
func (m *V1beta1SystemsItemSwitchesItemWithSwitchGetResponse) SetSystemId(value *string)() {
    m.systemId = value
}
type V1beta1SystemsItemSwitchesItemWithSwitchGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirmwareVersion()(*string)
    GetHealth()(V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable)
    GetIpAddress()(*string)
    GetRackName()(*string)
    GetSerialNumber()(*string)
    GetSystemId()(*string)
    SetFirmwareVersion(value *string)()
    SetHealth(value V1beta1SystemsItemSwitchesItemWithSwitchGetResponse_healthable)()
    SetIpAddress(value *string)()
    SetRackName(value *string)()
    SetSerialNumber(value *string)()
    SetSystemId(value *string)()
}
