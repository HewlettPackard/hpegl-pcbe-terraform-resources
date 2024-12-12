package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics active network adapter information.
type V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Device name of the network adapter.
    device *string
}
// NewV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics instantiates a new V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics and sets the default values.
func NewV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics()(*V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) {
    m := &V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDevice gets the device property value. Device name of the network adapter.
// returns a *string when successful
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) GetDevice()(*string) {
    return m.device
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("device", m.GetDevice())
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
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDevice sets the device property value. Device name of the network adapter.
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNics) SetDevice(value *string)() {
    m.device = value
}
type V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevice()(*string)
    SetDevice(value *string)()
}
