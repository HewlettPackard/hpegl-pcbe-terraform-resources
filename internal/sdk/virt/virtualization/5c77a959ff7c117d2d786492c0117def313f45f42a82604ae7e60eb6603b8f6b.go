package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Network connections of the machine instance.
    connections []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable
    // The cloud provider security groups for the machine instance.
    securityGroups []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo()(*V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConnections gets the connections property value. Network connections of the machine instance.
// returns a []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) GetConnections()([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable) {
    return m.connections
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable)
                }
            }
            m.SetConnections(res)
        }
        return nil
    }
    res["securityGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable)
                }
            }
            m.SetSecurityGroups(res)
        }
        return nil
    }
    return res
}
// GetSecurityGroups gets the securityGroups property value. The cloud provider security groups for the machine instance.
// returns a []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) GetSecurityGroups()([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable) {
    return m.securityGroups
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnections()))
        for i, v := range m.GetConnections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("connections", cast)
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
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConnections sets the connections property value. Network connections of the machine instance.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) SetConnections(value []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable)() {
    m.connections = value
}
// SetSecurityGroups sets the securityGroups property value. The cloud provider security groups for the machine instance.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo) SetSecurityGroups(value []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable)() {
    m.securityGroups = value
}
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnections()([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable)
    GetSecurityGroups()([]V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable)
    SetConnections(value []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_connectionsable)()
    SetSecurityGroups(value []V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable)()
}
