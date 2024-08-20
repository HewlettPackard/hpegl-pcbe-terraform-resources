package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cloud provider security groups this RDS database instance is attached to
    securityGroups []string
    // The name of the subnet group this RDS Instance is associated with
    subnetGroup V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable
    // Information about the VPC to which the RDS instance is associated.
    vpcInfo V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable
    // The cloud provider VPC security groups this RDS database instance is attached to
    vpcSecurityGroupIds []string
}
// NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo instantiates a new V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo and sets the default values.
func NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo()(*V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) {
    m := &V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["securityGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSecurityGroups(res)
        }
        return nil
    }
    res["subnetGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetGroup(val.(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable))
        }
        return nil
    }
    res["vpcInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpcInfo(val.(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable))
        }
        return nil
    }
    res["vpcSecurityGroupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetVpcSecurityGroupIds(res)
        }
        return nil
    }
    return res
}
// GetSecurityGroups gets the securityGroups property value. The cloud provider security groups this RDS database instance is attached to
// returns a []string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) GetSecurityGroups()([]string) {
    return m.securityGroups
}
// GetSubnetGroup gets the subnetGroup property value. The name of the subnet group this RDS Instance is associated with
// returns a V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) GetSubnetGroup()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable) {
    return m.subnetGroup
}
// GetVpcInfo gets the vpcInfo property value. Information about the VPC to which the RDS instance is associated.
// returns a V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) GetVpcInfo()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable) {
    return m.vpcInfo
}
// GetVpcSecurityGroupIds gets the vpcSecurityGroupIds property value. The cloud provider VPC security groups this RDS database instance is attached to
// returns a []string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) GetVpcSecurityGroupIds()([]string) {
    return m.vpcSecurityGroupIds
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecurityGroups sets the securityGroups property value. The cloud provider security groups this RDS database instance is attached to
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) SetSecurityGroups(value []string)() {
    m.securityGroups = value
}
// SetSubnetGroup sets the subnetGroup property value. The name of the subnet group this RDS Instance is associated with
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) SetSubnetGroup(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable)() {
    m.subnetGroup = value
}
// SetVpcInfo sets the vpcInfo property value. Information about the VPC to which the RDS instance is associated.
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) SetVpcInfo(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable)() {
    m.vpcInfo = value
}
// SetVpcSecurityGroupIds sets the vpcSecurityGroupIds property value. The cloud provider VPC security groups this RDS database instance is attached to
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo) SetVpcSecurityGroupIds(value []string)() {
    m.vpcSecurityGroupIds = value
}
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecurityGroups()([]string)
    GetSubnetGroup()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable)
    GetVpcInfo()(V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable)
    GetVpcSecurityGroupIds()([]string)
    SetSecurityGroups(value []string)()
    SetSubnetGroup(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable)()
    SetVpcInfo(value V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_vpcInfoable)()
    SetVpcSecurityGroupIds(value []string)()
}
