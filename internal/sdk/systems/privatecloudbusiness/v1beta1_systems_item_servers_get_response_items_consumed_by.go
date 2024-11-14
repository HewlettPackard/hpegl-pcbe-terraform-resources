package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SystemsItemServersGetResponse_items_consumedBy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique identifier (an UUID) of the consumer resource (e.g. an AI K8S cluster or a vmware ESXi hypervisor cluster)
    consumerId *string
    // Identifier of the server as maintained by consuming platform service e.g. DSCC EZ k8s service
    serverIdInPlatformService *string
}
// NewV1beta1SystemsItemServersGetResponse_items_consumedBy instantiates a new V1beta1SystemsItemServersGetResponse_items_consumedBy and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_consumedBy()(*V1beta1SystemsItemServersGetResponse_items_consumedBy) {
    m := &V1beta1SystemsItemServersGetResponse_items_consumedBy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_consumedByFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_consumedByFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_consumedBy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsumerId gets the consumerId property value. Unique identifier (an UUID) of the consumer resource (e.g. an AI K8S cluster or a vmware ESXi hypervisor cluster)
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) GetConsumerId()(*string) {
    return m.consumerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["consumerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsumerId(val)
        }
        return nil
    }
    res["serverIdInPlatformService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerIdInPlatformService(val)
        }
        return nil
    }
    return res
}
// GetServerIdInPlatformService gets the serverIdInPlatformService property value. Identifier of the server as maintained by consuming platform service e.g. DSCC EZ k8s service
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) GetServerIdInPlatformService()(*string) {
    return m.serverIdInPlatformService
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("consumerId", m.GetConsumerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serverIdInPlatformService", m.GetServerIdInPlatformService())
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
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsumerId sets the consumerId property value. Unique identifier (an UUID) of the consumer resource (e.g. an AI K8S cluster or a vmware ESXi hypervisor cluster)
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) SetConsumerId(value *string)() {
    m.consumerId = value
}
// SetServerIdInPlatformService sets the serverIdInPlatformService property value. Identifier of the server as maintained by consuming platform service e.g. DSCC EZ k8s service
func (m *V1beta1SystemsItemServersGetResponse_items_consumedBy) SetServerIdInPlatformService(value *string)() {
    m.serverIdInPlatformService = value
}
type V1beta1SystemsItemServersGetResponse_items_consumedByable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsumerId()(*string)
    GetServerIdInPlatformService()(*string)
    SetConsumerId(value *string)()
    SetServerIdInPlatformService(value *string)()
}
