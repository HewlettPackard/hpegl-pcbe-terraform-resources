package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs parameters used to automatically set up EBS volumes when the instance is launched.
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the EBS volume is deleted on instance termination.
    deleteOnTermination *bool
    // Indicates whether the encryption state of an EBS volume is changed while being restored from a backing snapshot.
    encrypted *bool
    // The identifier of the snapshot.
    snapshotId *string
    // The size of the volume, in GiBs.
    volumeSize *int32
    // The volume type.
    volumeType *string
}
// NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs instantiates a new V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs and sets the default values.
func NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs()(*V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) {
    m := &V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeleteOnTermination gets the deleteOnTermination property value. Indicates whether the EBS volume is deleted on instance termination.
// returns a *bool when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetDeleteOnTermination()(*bool) {
    return m.deleteOnTermination
}
// GetEncrypted gets the encrypted property value. Indicates whether the encryption state of an EBS volume is changed while being restored from a backing snapshot.
// returns a *bool when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetEncrypted()(*bool) {
    return m.encrypted
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deleteOnTermination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteOnTermination(val)
        }
        return nil
    }
    res["encrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncrypted(val)
        }
        return nil
    }
    res["snapshotId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnapshotId(val)
        }
        return nil
    }
    res["volumeSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeSize(val)
        }
        return nil
    }
    res["volumeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeType(val)
        }
        return nil
    }
    return res
}
// GetSnapshotId gets the snapshotId property value. The identifier of the snapshot.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetSnapshotId()(*string) {
    return m.snapshotId
}
// GetVolumeSize gets the volumeSize property value. The size of the volume, in GiBs.
// returns a *int32 when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetVolumeSize()(*int32) {
    return m.volumeSize
}
// GetVolumeType gets the volumeType property value. The volume type.
// returns a *string when successful
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) GetVolumeType()(*string) {
    return m.volumeType
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("deleteOnTermination", m.GetDeleteOnTermination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("encrypted", m.GetEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("snapshotId", m.GetSnapshotId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("volumeSize", m.GetVolumeSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("volumeType", m.GetVolumeType())
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
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeleteOnTermination sets the deleteOnTermination property value. Indicates whether the EBS volume is deleted on instance termination.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) SetDeleteOnTermination(value *bool)() {
    m.deleteOnTermination = value
}
// SetEncrypted sets the encrypted property value. Indicates whether the encryption state of an EBS volume is changed while being restored from a backing snapshot.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) SetEncrypted(value *bool)() {
    m.encrypted = value
}
// SetSnapshotId sets the snapshotId property value. The identifier of the snapshot.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) SetSnapshotId(value *string)() {
    m.snapshotId = value
}
// SetVolumeSize sets the volumeSize property value. The size of the volume, in GiBs.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) SetVolumeSize(value *int32)() {
    m.volumeSize = value
}
// SetVolumeType sets the volumeType property value. The volume type.
func (m *V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebs) SetVolumeType(value *string)() {
    m.volumeType = value
}
type V1beta1CspMachineImagesGetResponse_items_cspInfoMember1_blockDeviceMappings_ebsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeleteOnTermination()(*bool)
    GetEncrypted()(*bool)
    GetSnapshotId()(*string)
    GetVolumeSize()(*int32)
    GetVolumeType()(*string)
    SetDeleteOnTermination(value *bool)()
    SetEncrypted(value *bool)()
    SetSnapshotId(value *string)()
    SetVolumeSize(value *int32)()
    SetVolumeType(value *string)()
}
