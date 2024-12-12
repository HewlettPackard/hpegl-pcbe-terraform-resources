package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1VirtualMachinesPostRequestBody_storageConfig specifies the storage configurations for a virtual machine
type V1beta1VirtualMachinesPostRequestBody_storageConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The UUID of the hypervisor datastore where the virtual machine is to be deployed
    defaultDatastoreId *string
    // Attributes of the volume.
    volumeInfo V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable
}
// NewV1beta1VirtualMachinesPostRequestBody_storageConfig instantiates a new V1beta1VirtualMachinesPostRequestBody_storageConfig and sets the default values.
func NewV1beta1VirtualMachinesPostRequestBody_storageConfig()(*V1beta1VirtualMachinesPostRequestBody_storageConfig) {
    m := &V1beta1VirtualMachinesPostRequestBody_storageConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1VirtualMachinesPostRequestBody_storageConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1VirtualMachinesPostRequestBody_storageConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1VirtualMachinesPostRequestBody_storageConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultDatastoreId gets the defaultDatastoreId property value. The UUID of the hypervisor datastore where the virtual machine is to be deployed
// returns a *string when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) GetDefaultDatastoreId()(*string) {
    return m.defaultDatastoreId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultDatastoreId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDatastoreId(val)
        }
        return nil
    }
    res["volumeInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeInfo(val.(V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable))
        }
        return nil
    }
    return res
}
// GetVolumeInfo gets the volumeInfo property value. Attributes of the volume.
// returns a V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable when successful
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) GetVolumeInfo()(V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable) {
    return m.volumeInfo
}
// Serialize serializes information the current object
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultDatastoreId", m.GetDefaultDatastoreId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("volumeInfo", m.GetVolumeInfo())
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
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultDatastoreId sets the defaultDatastoreId property value. The UUID of the hypervisor datastore where the virtual machine is to be deployed
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) SetDefaultDatastoreId(value *string)() {
    m.defaultDatastoreId = value
}
// SetVolumeInfo sets the volumeInfo property value. Attributes of the volume.
func (m *V1beta1VirtualMachinesPostRequestBody_storageConfig) SetVolumeInfo(value V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable)() {
    m.volumeInfo = value
}
type V1beta1VirtualMachinesPostRequestBody_storageConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultDatastoreId()(*string)
    GetVolumeInfo()(V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable)
    SetDefaultDatastoreId(value *string)()
    SetVolumeInfo(value V1beta1VirtualMachinesPostRequestBody_storageConfig_volumeInfoable)()
}
