package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsGetResponse_items_interfaces struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The network property
    network V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable
}
// NewV1beta1DataServicesConnectorsGetResponse_items_interfaces instantiates a new V1beta1DataServicesConnectorsGetResponse_items_interfaces and sets the default values.
func NewV1beta1DataServicesConnectorsGetResponse_items_interfaces()(*V1beta1DataServicesConnectorsGetResponse_items_interfaces) {
    m := &V1beta1DataServicesConnectorsGetResponse_items_interfaces{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsGetResponse_items_interfacesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsGetResponse_items_interfacesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsGetResponse_items_interfaces(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_networkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val.(V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable))
        }
        return nil
    }
    return res
}
// GetNetwork gets the network property value. The network property
// returns a V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces) GetNetwork()(V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable) {
    return m.network
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("network", m.GetNetwork())
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
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetwork sets the network property value. The network property
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces) SetNetwork(value V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable)() {
    m.network = value
}
type V1beta1DataServicesConnectorsGetResponse_items_interfacesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetwork()(V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable)
    SetNetwork(value V1beta1DataServicesConnectorsGetResponse_items_interfaces_networkable)()
}
