package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifier for the network mapping
    name *string
    // Target network to be used for deployment
    network *string
}
// NewV1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping instantiates a new V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping()(*V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) {
    m := &V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Identifier for the network mapping
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) GetName()(*string) {
    return m.name
}
// GetNetwork gets the network property value. Target network to be used for deployment
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) GetNetwork()(*string) {
    return m.network
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
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
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Identifier for the network mapping
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) SetName(value *string)() {
    m.name = value
}
// SetNetwork sets the network property value. Target network to be used for deployment
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMapping) SetNetwork(value *string)() {
    m.network = value
}
type V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNetwork()(*string)
    SetName(value *string)()
    SetNetwork(value *string)()
}
