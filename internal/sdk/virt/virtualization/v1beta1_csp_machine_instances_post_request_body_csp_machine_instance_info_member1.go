package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the csp machine image resource representing the cloud provider machine image
    imageId *string
    // Instance Type chosen by user
    instanceType *string
    // Region on which CSP machine instance will be created
    region *string
}
// NewV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 instantiates a new V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1 and sets the default values.
func NewV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1()(*V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) {
    m := &V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["imageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["instanceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceType(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    return res
}
// GetImageId gets the imageId property value. ID of the csp machine image resource representing the cloud provider machine image
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) GetImageId()(*string) {
    return m.imageId
}
// GetInstanceType gets the instanceType property value. Instance Type chosen by user
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) GetInstanceType()(*string) {
    return m.instanceType
}
// GetRegion gets the region property value. Region on which CSP machine instance will be created
// returns a *string when successful
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) GetRegion()(*string) {
    return m.region
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("imageId", m.GetImageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceType", m.GetInstanceType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
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
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImageId sets the imageId property value. ID of the csp machine image resource representing the cloud provider machine image
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) SetImageId(value *string)() {
    m.imageId = value
}
// SetInstanceType sets the instanceType property value. Instance Type chosen by user
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) SetInstanceType(value *string)() {
    m.instanceType = value
}
// SetRegion sets the region property value. Region on which CSP machine instance will be created
func (m *V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1) SetRegion(value *string)() {
    m.region = value
}
type V1beta1CspMachineInstancesPostRequestBody_cspMachineInstanceInfoMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImageId()(*string)
    GetInstanceType()(*string)
    GetRegion()(*string)
    SetImageId(value *string)()
    SetInstanceType(value *string)()
    SetRegion(value *string)()
}
