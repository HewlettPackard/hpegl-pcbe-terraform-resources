package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspRdsInstancesGetResponse_items_metadata metadata information related to possible backups supported for given instance
type V1beta1CspRdsInstancesGetResponse_items_metadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates if HPE Cloud backup enabled for this asset
    cloudBackupEnabled *bool
    // Indicates if the cloud native backup enabled for this asset
    cloudNativeBackupEnabled *bool
    // Message indicates the reason if cloudNativeBackupEnabled or hpeCloudBackupEnabled or both are set to false
    notSupportedReason *string
}
// NewV1beta1CspRdsInstancesGetResponse_items_metadata instantiates a new V1beta1CspRdsInstancesGetResponse_items_metadata and sets the default values.
func NewV1beta1CspRdsInstancesGetResponse_items_metadata()(*V1beta1CspRdsInstancesGetResponse_items_metadata) {
    m := &V1beta1CspRdsInstancesGetResponse_items_metadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesGetResponse_items_metadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesGetResponse_items_metadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesGetResponse_items_metadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudBackupEnabled gets the cloudBackupEnabled property value. Indicates if HPE Cloud backup enabled for this asset
// returns a *bool when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) GetCloudBackupEnabled()(*bool) {
    return m.cloudBackupEnabled
}
// GetCloudNativeBackupEnabled gets the cloudNativeBackupEnabled property value. Indicates if the cloud native backup enabled for this asset
// returns a *bool when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) GetCloudNativeBackupEnabled()(*bool) {
    return m.cloudNativeBackupEnabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudBackupEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudBackupEnabled(val)
        }
        return nil
    }
    res["cloudNativeBackupEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudNativeBackupEnabled(val)
        }
        return nil
    }
    res["notSupportedReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotSupportedReason(val)
        }
        return nil
    }
    return res
}
// GetNotSupportedReason gets the notSupportedReason property value. Message indicates the reason if cloudNativeBackupEnabled or hpeCloudBackupEnabled or both are set to false
// returns a *string when successful
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) GetNotSupportedReason()(*string) {
    return m.notSupportedReason
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudBackupEnabled sets the cloudBackupEnabled property value. Indicates if HPE Cloud backup enabled for this asset
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) SetCloudBackupEnabled(value *bool)() {
    m.cloudBackupEnabled = value
}
// SetCloudNativeBackupEnabled sets the cloudNativeBackupEnabled property value. Indicates if the cloud native backup enabled for this asset
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) SetCloudNativeBackupEnabled(value *bool)() {
    m.cloudNativeBackupEnabled = value
}
// SetNotSupportedReason sets the notSupportedReason property value. Message indicates the reason if cloudNativeBackupEnabled or hpeCloudBackupEnabled or both are set to false
func (m *V1beta1CspRdsInstancesGetResponse_items_metadata) SetNotSupportedReason(value *string)() {
    m.notSupportedReason = value
}
type V1beta1CspRdsInstancesGetResponse_items_metadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudBackupEnabled()(*bool)
    GetCloudNativeBackupEnabled()(*bool)
    GetNotSupportedReason()(*string)
    SetCloudBackupEnabled(value *bool)()
    SetCloudNativeBackupEnabled(value *bool)()
    SetNotSupportedReason(value *string)()
}
