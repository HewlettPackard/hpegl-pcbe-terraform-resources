package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresPostRequestBody request body for Create Datastore API
type V1beta1DatastoresPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of datastore to be created
    name *string
    // Unique identifier of a protection policy
    protectionPolicyId *string
    // Size of the datastore in bytes
    sizeInBytes *int64
    // Storage System ID
    storageSystemId *string
    // Target Hypervisor Cluster ID
    targetHypervisorClusterId *string
    // Attributes of the volume.
    volumeInfo V1beta1DatastoresPostRequestBody_volumeInfoable
}
// NewV1beta1DatastoresPostRequestBody instantiates a new V1beta1DatastoresPostRequestBody and sets the default values.
func NewV1beta1DatastoresPostRequestBody()(*V1beta1DatastoresPostRequestBody) {
    m := &V1beta1DatastoresPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["protectionPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionPolicyId(val)
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    res["storageSystemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageSystemId(val)
        }
        return nil
    }
    res["targetHypervisorClusterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetHypervisorClusterId(val)
        }
        return nil
    }
    res["volumeInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresPostRequestBody_volumeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeInfo(val.(V1beta1DatastoresPostRequestBody_volumeInfoable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of datastore to be created
// returns a *string when successful
func (m *V1beta1DatastoresPostRequestBody) GetName()(*string) {
    return m.name
}
// GetProtectionPolicyId gets the protectionPolicyId property value. Unique identifier of a protection policy
// returns a *string when successful
func (m *V1beta1DatastoresPostRequestBody) GetProtectionPolicyId()(*string) {
    return m.protectionPolicyId
}
// GetSizeInBytes gets the sizeInBytes property value. Size of the datastore in bytes
// returns a *int64 when successful
func (m *V1beta1DatastoresPostRequestBody) GetSizeInBytes()(*int64) {
    return m.sizeInBytes
}
// GetStorageSystemId gets the storageSystemId property value. Storage System ID
// returns a *string when successful
func (m *V1beta1DatastoresPostRequestBody) GetStorageSystemId()(*string) {
    return m.storageSystemId
}
// GetTargetHypervisorClusterId gets the targetHypervisorClusterId property value. Target Hypervisor Cluster ID
// returns a *string when successful
func (m *V1beta1DatastoresPostRequestBody) GetTargetHypervisorClusterId()(*string) {
    return m.targetHypervisorClusterId
}
// GetVolumeInfo gets the volumeInfo property value. Attributes of the volume.
// returns a V1beta1DatastoresPostRequestBody_volumeInfoable when successful
func (m *V1beta1DatastoresPostRequestBody) GetVolumeInfo()(V1beta1DatastoresPostRequestBody_volumeInfoable) {
    return m.volumeInfo
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("protectionPolicyId", m.GetProtectionPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("storageSystemId", m.GetStorageSystemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetHypervisorClusterId", m.GetTargetHypervisorClusterId())
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
func (m *V1beta1DatastoresPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of datastore to be created
func (m *V1beta1DatastoresPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetProtectionPolicyId sets the protectionPolicyId property value. Unique identifier of a protection policy
func (m *V1beta1DatastoresPostRequestBody) SetProtectionPolicyId(value *string)() {
    m.protectionPolicyId = value
}
// SetSizeInBytes sets the sizeInBytes property value. Size of the datastore in bytes
func (m *V1beta1DatastoresPostRequestBody) SetSizeInBytes(value *int64)() {
    m.sizeInBytes = value
}
// SetStorageSystemId sets the storageSystemId property value. Storage System ID
func (m *V1beta1DatastoresPostRequestBody) SetStorageSystemId(value *string)() {
    m.storageSystemId = value
}
// SetTargetHypervisorClusterId sets the targetHypervisorClusterId property value. Target Hypervisor Cluster ID
func (m *V1beta1DatastoresPostRequestBody) SetTargetHypervisorClusterId(value *string)() {
    m.targetHypervisorClusterId = value
}
// SetVolumeInfo sets the volumeInfo property value. Attributes of the volume.
func (m *V1beta1DatastoresPostRequestBody) SetVolumeInfo(value V1beta1DatastoresPostRequestBody_volumeInfoable)() {
    m.volumeInfo = value
}
type V1beta1DatastoresPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetProtectionPolicyId()(*string)
    GetSizeInBytes()(*int64)
    GetStorageSystemId()(*string)
    GetTargetHypervisorClusterId()(*string)
    GetVolumeInfo()(V1beta1DatastoresPostRequestBody_volumeInfoable)
    SetName(value *string)()
    SetProtectionPolicyId(value *string)()
    SetSizeInBytes(value *int64)()
    SetStorageSystemId(value *string)()
    SetTargetHypervisorClusterId(value *string)()
    SetVolumeInfo(value V1beta1DatastoresPostRequestBody_volumeInfoable)()
}
