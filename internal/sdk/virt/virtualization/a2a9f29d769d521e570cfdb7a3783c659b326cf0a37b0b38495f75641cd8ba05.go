package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards the network cards for the instance type.
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The maximum number of network interfaces for the network card.
    maximumNetworkInterfaces *int32
    // The index of the network card.
    networkCardIndex *int32
    // The network performance of the network card.
    networkPerformance *string
}
// NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards instantiates a new V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards and sets the default values.
func NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards()(*V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) {
    m := &V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maximumNetworkInterfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumNetworkInterfaces(val)
        }
        return nil
    }
    res["networkCardIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkCardIndex(val)
        }
        return nil
    }
    res["networkPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkPerformance(val)
        }
        return nil
    }
    return res
}
// GetMaximumNetworkInterfaces gets the maximumNetworkInterfaces property value. The maximum number of network interfaces for the network card.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) GetMaximumNetworkInterfaces()(*int32) {
    return m.maximumNetworkInterfaces
}
// GetNetworkCardIndex gets the networkCardIndex property value. The index of the network card.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) GetNetworkCardIndex()(*int32) {
    return m.networkCardIndex
}
// GetNetworkPerformance gets the networkPerformance property value. The network performance of the network card.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) GetNetworkPerformance()(*string) {
    return m.networkPerformance
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maximumNetworkInterfaces", m.GetMaximumNetworkInterfaces())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("networkCardIndex", m.GetNetworkCardIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkPerformance", m.GetNetworkPerformance())
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
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaximumNetworkInterfaces sets the maximumNetworkInterfaces property value. The maximum number of network interfaces for the network card.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) SetMaximumNetworkInterfaces(value *int32)() {
    m.maximumNetworkInterfaces = value
}
// SetNetworkCardIndex sets the networkCardIndex property value. The index of the network card.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) SetNetworkCardIndex(value *int32)() {
    m.networkCardIndex = value
}
// SetNetworkPerformance sets the networkPerformance property value. The network performance of the network card.
func (m *V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCards) SetNetworkPerformance(value *string)() {
    m.networkPerformance = value
}
type V1beta1CspMachineInstanceTypesGetResponse_items_cspInfoMember1_networkInfo_networkCardsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumNetworkInterfaces()(*int32)
    GetNetworkCardIndex()(*int32)
    GetNetworkPerformance()(*string)
    SetMaximumNetworkInterfaces(value *int32)()
    SetNetworkCardIndex(value *int32)()
    SetNetworkPerformance(value *string)()
}
