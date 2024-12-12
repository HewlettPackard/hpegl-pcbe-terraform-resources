package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifier assigned to the security group in the cloud provider
    cspId *string
    // Name assigned to the security group in the cloud provider
    cspName *string
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups()(*V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCspId gets the cspId property value. Identifier assigned to the security group in the cloud provider
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) GetCspId()(*string) {
    return m.cspId
}
// GetCspName gets the cspName property value. Name assigned to the security group in the cloud provider
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) GetCspName()(*string) {
    return m.cspName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cspId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspId(val)
        }
        return nil
    }
    res["cspName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cspId", m.GetCspId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cspName", m.GetCspName())
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
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCspId sets the cspId property value. Identifier assigned to the security group in the cloud provider
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) SetCspId(value *string)() {
    m.cspId = value
}
// SetCspName sets the cspName property value. Name assigned to the security group in the cloud provider
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroups) SetCspName(value *string)() {
    m.cspName = value
}
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_cspInfoMember2_networkInfo_securityGroupsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCspId()(*string)
    GetCspName()(*string)
    SetCspId(value *string)()
    SetCspName(value *string)()
}
