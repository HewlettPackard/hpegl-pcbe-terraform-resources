package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem host network system information.
type V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of port group information.
    portGroups []V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable
}
// NewV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem instantiates a new V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem and sets the default values.
func NewV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem()(*V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) {
    m := &V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsGetResponse_items_hostNetworkSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsGetResponse_items_hostNetworkSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["portGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable)
                }
            }
            m.SetPortGroups(res)
        }
        return nil
    }
    return res
}
// GetPortGroups gets the portGroups property value. List of port group information.
// returns a []V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable when successful
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) GetPortGroups()([]V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable) {
    return m.portGroups
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPortGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPortGroups()))
        for i, v := range m.GetPortGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("portGroups", cast)
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
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPortGroups sets the portGroups property value. List of port group information.
func (m *V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem) SetPortGroups(value []V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable)() {
    m.portGroups = value
}
type V1beta1HypervisorHostsGetResponse_items_hostNetworkSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPortGroups()([]V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable)
    SetPortGroups(value []V1beta1HypervisorHostsGetResponse_items_hostNetworkSystem_portGroupsable)()
}
