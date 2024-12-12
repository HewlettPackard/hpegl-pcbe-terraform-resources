package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups port group information.
type V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The linkable identifier for the port group.
    key *string
    // Name of the port group.
    name *string
    // The network adapter teaming policy information.
    nicTeamingPolicy V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable
    // VLAN ID associated with the port group.
    vlanId *int32
    // Virtual switch that contains the port group.
    vswitch *string
    // Virtual switch name.
    vswitchName *string
}
// NewV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups instantiates a new V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups and sets the default values.
func NewV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups()(*V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) {
    m := &V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
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
    res["nicTeamingPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNicTeamingPolicy(val.(V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable))
        }
        return nil
    }
    res["vlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVlanId(val)
        }
        return nil
    }
    res["vswitch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVswitch(val)
        }
        return nil
    }
    res["vswitchName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVswitchName(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The linkable identifier for the port group.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetKey()(*string) {
    return m.key
}
// GetName gets the name property value. Name of the port group.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetName()(*string) {
    return m.name
}
// GetNicTeamingPolicy gets the nicTeamingPolicy property value. The network adapter teaming policy information.
// returns a V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetNicTeamingPolicy()(V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable) {
    return m.nicTeamingPolicy
}
// GetVlanId gets the vlanId property value. VLAN ID associated with the port group.
// returns a *int32 when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetVlanId()(*int32) {
    return m.vlanId
}
// GetVswitch gets the vswitch property value. Virtual switch that contains the port group.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetVswitch()(*string) {
    return m.vswitch
}
// GetVswitchName gets the vswitchName property value. Virtual switch name.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) GetVswitchName()(*string) {
    return m.vswitchName
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("nicTeamingPolicy", m.GetNicTeamingPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("vlanId", m.GetVlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vswitch", m.GetVswitch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vswitchName", m.GetVswitchName())
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
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. The linkable identifier for the port group.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetKey(value *string)() {
    m.key = value
}
// SetName sets the name property value. Name of the port group.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetName(value *string)() {
    m.name = value
}
// SetNicTeamingPolicy sets the nicTeamingPolicy property value. The network adapter teaming policy information.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetNicTeamingPolicy(value V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable)() {
    m.nicTeamingPolicy = value
}
// SetVlanId sets the vlanId property value. VLAN ID associated with the port group.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetVlanId(value *int32)() {
    m.vlanId = value
}
// SetVswitch sets the vswitch property value. Virtual switch that contains the port group.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetVswitch(value *string)() {
    m.vswitch = value
}
// SetVswitchName sets the vswitchName property value. Virtual switch name.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups) SetVswitchName(value *string)() {
    m.vswitchName = value
}
type V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetName()(*string)
    GetNicTeamingPolicy()(V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable)
    GetVlanId()(*int32)
    GetVswitch()(*string)
    GetVswitchName()(*string)
    SetKey(value *string)()
    SetName(value *string)()
    SetNicTeamingPolicy(value V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable)()
    SetVlanId(value *int32)()
    SetVswitch(value *string)()
    SetVswitchName(value *string)()
}
