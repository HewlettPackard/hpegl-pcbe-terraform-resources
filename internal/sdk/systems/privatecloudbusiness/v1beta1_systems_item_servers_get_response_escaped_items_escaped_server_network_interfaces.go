package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces network interface information
type V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // MAC address
    macAddress *string
    // speed
    speedGbps *float64
}
// NewV1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces instantiates a new V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces()(*V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) {
    m := &V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    res["speedGbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeedGbps(val)
        }
        return nil
    }
    return res
}
// GetMacAddress gets the macAddress property value. MAC address
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) GetMacAddress()(*string) {
    return m.macAddress
}
// GetSpeedGbps gets the speedGbps property value. speed
// returns a *float64 when successful
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) GetSpeedGbps()(*float64) {
    return m.speedGbps
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("speedGbps", m.GetSpeedGbps())
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
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMacAddress sets the macAddress property value. MAC address
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetSpeedGbps sets the speedGbps property value. speed
func (m *V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfaces) SetSpeedGbps(value *float64)() {
    m.speedGbps = value
}
type V1beta1SystemsItemServersGetResponse_items_serverNetworkInterfacesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMacAddress()(*string)
    GetSpeedGbps()(*float64)
    SetMacAddress(value *string)()
    SetSpeedGbps(value *float64)()
}
