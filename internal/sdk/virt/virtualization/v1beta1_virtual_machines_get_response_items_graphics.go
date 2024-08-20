package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesGetResponse_items_graphics graphics configuration for the virtual machine.
type V1beta1VirtualMachinesGetResponse_items_graphics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Port number for the graphics adapter.
    port *int32
    // Port number for the web socket.
    webSocketPort *int32
}
// NewV1beta1VirtualMachinesGetResponse_items_graphics instantiates a new V1beta1VirtualMachinesGetResponse_items_graphics and sets the default values.
func NewV1beta1VirtualMachinesGetResponse_items_graphics()(*V1beta1VirtualMachinesGetResponse_items_graphics) {
    m := &V1beta1VirtualMachinesGetResponse_items_graphics{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesGetResponse_items_graphicsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesGetResponse_items_graphicsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesGetResponse_items_graphics(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["webSocketPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebSocketPort(val)
        }
        return nil
    }
    return res
}
// GetPort gets the port property value. Port number for the graphics adapter.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) GetPort()(*int32) {
    return m.port
}
// GetWebSocketPort gets the webSocketPort property value. Port number for the web socket.
// returns a *int32 when successful
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) GetWebSocketPort()(*int32) {
    return m.webSocketPort
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("webSocketPort", m.GetWebSocketPort())
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
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPort sets the port property value. Port number for the graphics adapter.
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) SetPort(value *int32)() {
    m.port = value
}
// SetWebSocketPort sets the webSocketPort property value. Port number for the web socket.
func (m *V1beta1VirtualMachinesGetResponse_items_graphics) SetWebSocketPort(value *int32)() {
    m.webSocketPort = value
}
type V1beta1VirtualMachinesGetResponse_items_graphicsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPort()(*int32)
    GetWebSocketPort()(*int32)
    SetPort(value *int32)()
    SetWebSocketPort(value *int32)()
}
