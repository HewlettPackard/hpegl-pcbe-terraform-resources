package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo information about the protection policy that was used to create the protection job
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User defined name of the protection policy
    name *string
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo instantiates a new V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo and sets the default values.
func NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo()(*V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) {
    m := &V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. User defined name of the protection policy
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. User defined name of the protection policy
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfo) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_protectionJobInfo_protectionPolicyInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
}
