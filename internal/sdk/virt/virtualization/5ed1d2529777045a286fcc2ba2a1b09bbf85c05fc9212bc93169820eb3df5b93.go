package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings block device mapping
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The device name (for example, /dev/sdh or xvdh ).
    deviceName *string
    // Parameters used to automatically set up EBS volumes when the instance is launched.
    ebs V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable
}
// NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings instantiates a new V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings and sets the default values.
func NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings()(*V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) {
    m := &V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceName gets the deviceName property value. The device name (for example, /dev/sdh or xvdh ).
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) GetDeviceName()(*string) {
    return m.deviceName
}
// GetEbs gets the ebs property value. Parameters used to automatically set up EBS volumes when the instance is launched.
// returns a V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) GetEbs()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable) {
    return m.ebs
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["ebs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEbs(val.(V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ebs", m.GetEbs())
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
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceName sets the deviceName property value. The device name (for example, /dev/sdh or xvdh ).
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetEbs sets the ebs property value. Parameters used to automatically set up EBS volumes when the instance is launched.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings) SetEbs(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable)() {
    m.ebs = value
}
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceName()(*string)
    GetEbs()(V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable)
    SetDeviceName(value *string)()
    SetEbs(value V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable)()
}
