package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy the network adapter teaming policy information.
type V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy struct {
    // List of active network adapter information.
    activeNics []V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Flag to specify whether or not to notify the switch if a link fails.
    notifySwitches *bool
    // Name of the network adapter teaming policy.
    policyName *string
    // Flag to specify whether or not to use a rolling policy when restoring links.
    rollingOrder *bool
    // Flag to specify whether or not to use beacon probing.
    useBeaconProbing *bool
}
// NewV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy instantiates a new V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy and sets the default values.
func NewV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy()(*V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) {
    m := &V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy(), nil
}
// GetActiveNics gets the activeNics property value. List of active network adapter information.
// returns a []V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetActiveNics()([]V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable) {
    return m.activeNics
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activeNics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable)
                }
            }
            m.SetActiveNics(res)
        }
        return nil
    }
    res["notifySwitches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifySwitches(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["rollingOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRollingOrder(val)
        }
        return nil
    }
    res["useBeaconProbing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseBeaconProbing(val)
        }
        return nil
    }
    return res
}
// GetNotifySwitches gets the notifySwitches property value. Flag to specify whether or not to notify the switch if a link fails.
// returns a *bool when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetNotifySwitches()(*bool) {
    return m.notifySwitches
}
// GetPolicyName gets the policyName property value. Name of the network adapter teaming policy.
// returns a *string when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetPolicyName()(*string) {
    return m.policyName
}
// GetRollingOrder gets the rollingOrder property value. Flag to specify whether or not to use a rolling policy when restoring links.
// returns a *bool when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetRollingOrder()(*bool) {
    return m.rollingOrder
}
// GetUseBeaconProbing gets the useBeaconProbing property value. Flag to specify whether or not to use beacon probing.
// returns a *bool when successful
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) GetUseBeaconProbing()(*bool) {
    return m.useBeaconProbing
}
// Serialize serializes information the current object
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActiveNics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActiveNics()))
        for i, v := range m.GetActiveNics() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("activeNics", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("notifySwitches", m.GetNotifySwitches())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("rollingOrder", m.GetRollingOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useBeaconProbing", m.GetUseBeaconProbing())
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
// SetActiveNics sets the activeNics property value. List of active network adapter information.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) SetActiveNics(value []V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable)() {
    m.activeNics = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNotifySwitches sets the notifySwitches property value. Flag to specify whether or not to notify the switch if a link fails.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) SetNotifySwitches(value *bool)() {
    m.notifySwitches = value
}
// SetPolicyName sets the policyName property value. Name of the network adapter teaming policy.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) SetPolicyName(value *string)() {
    m.policyName = value
}
// SetRollingOrder sets the rollingOrder property value. Flag to specify whether or not to use a rolling policy when restoring links.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) SetRollingOrder(value *bool)() {
    m.rollingOrder = value
}
// SetUseBeaconProbing sets the useBeaconProbing property value. Flag to specify whether or not to use beacon probing.
func (m *V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy) SetUseBeaconProbing(value *bool)() {
    m.useBeaconProbing = value
}
type V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveNics()([]V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable)
    GetNotifySwitches()(*bool)
    GetPolicyName()(*string)
    GetRollingOrder()(*bool)
    GetUseBeaconProbing()(*bool)
    SetActiveNics(value []V1beta1HypervisorHostsItemHostGetResponse_hostNetworkSystem_portGroups_nicTeamingPolicy_activeNicsable)()
    SetNotifySwitches(value *bool)()
    SetPolicyName(value *string)()
    SetRollingOrder(value *bool)()
    SetUseBeaconProbing(value *bool)()
}
