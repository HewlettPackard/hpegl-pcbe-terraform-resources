package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody_networkConfig specifies name and the target network to use for deployment
type V1beta1VirtualMachinesPostRequestBody_networkConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies name and the target network to be used for deployment
    networkMapping []V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable
}
// NewV1beta1VirtualMachinesPostRequestBody_networkConfig instantiates a new V1beta1VirtualMachinesPostRequestBody_networkConfig and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_networkConfig()(*V1beta1VirtualMachinesPostRequestBody_networkConfig) {
    m := &V1beta1VirtualMachinesPostRequestBody_networkConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_networkConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_networkConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_networkConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["networkMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable)
                }
            }
            m.SetNetworkMapping(res)
        }
        return nil
    }
    return res
}
// GetNetworkMapping gets the networkMapping property value. Specifies name and the target network to be used for deployment
// returns a []V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable when successful
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig) GetNetworkMapping()([]V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable) {
    return m.networkMapping
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNetworkMapping() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkMapping()))
        for i, v := range m.GetNetworkMapping() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkMapping", cast)
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
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetworkMapping sets the networkMapping property value. Specifies name and the target network to be used for deployment
func (m *V1beta1VirtualMachinesPostRequestBody_networkConfig) SetNetworkMapping(value []V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable)() {
    m.networkMapping = value
}
type V1beta1VirtualMachinesPostRequestBody_networkConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkMapping()([]V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable)
    SetNetworkMapping(value []V1beta1VirtualMachinesPostRequestBody_networkConfig_networkMappingable)()
}
