package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo describes the attachment of a CSP volume to a CSP machine instance.
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A CSP machine instance to which the volume is attached (when given in the context of aCSP volume) or an attached CSP volume (when given in the context of a CSP machine instance.
    attachedTo V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable
    // The device name, for example /dev/sdh or xvdh. Supported by AWS cloud service provider.
    device *string
    // The logical unit number of the disk. Supported by Azure cloud service provider.
    lun *int32
}
// NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo instantiates a new V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo and sets the default values.
func NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo()(*V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) {
    m := &V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachedTo gets the attachedTo property value. A CSP machine instance to which the volume is attached (when given in the context of aCSP volume) or an attached CSP volume (when given in the context of a CSP machine instance.
// returns a V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) GetAttachedTo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable) {
    return m.attachedTo
}
// GetDevice gets the device property value. The device name, for example /dev/sdh or xvdh. Supported by AWS cloud service provider.
// returns a *string when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) GetDevice()(*string) {
    return m.device
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachedTo(val.(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable))
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val)
        }
        return nil
    }
    res["lun"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLun(val)
        }
        return nil
    }
    return res
}
// GetLun gets the lun property value. The logical unit number of the disk. Supported by Azure cloud service provider.
// returns a *int32 when successful
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) GetLun()(*int32) {
    return m.lun
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachedTo", m.GetAttachedTo())
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
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachedTo sets the attachedTo property value. A CSP machine instance to which the volume is attached (when given in the context of aCSP volume) or an attached CSP volume (when given in the context of a CSP machine instance.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) SetAttachedTo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable)() {
    m.attachedTo = value
}
// SetDevice sets the device property value. The device name, for example /dev/sdh or xvdh. Supported by AWS cloud service provider.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) SetDevice(value *string)() {
    m.device = value
}
// SetLun sets the lun property value. The logical unit number of the disk. Supported by Azure cloud service provider.
func (m *V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo) SetLun(value *int32)() {
    m.lun = value
}
type V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachedTo()(V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable)
    GetDevice()(*string)
    GetLun()(*int32)
    SetAttachedTo(value V1beta1CspMachineInstancesItemCspMachineInstancesGetResponse_volumeAttachmentInfo_attachedToable)()
    SetDevice(value *string)()
    SetLun(value *int32)()
}
