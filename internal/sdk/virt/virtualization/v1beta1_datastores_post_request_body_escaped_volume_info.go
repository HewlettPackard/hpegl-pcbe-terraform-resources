package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresPostRequestBody_volumeInfo attributes of the volume.
type V1beta1DatastoresPostRequestBody_volumeInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether all flash is enabled or not.
    allFlash *bool
    // Specifies if data reduction capability is enabled. Data reduction is a combination of both deduplication and compression.
    dataReduction *bool
    // Specifies if dedupe is enabled for volumes created. Deduplication is a process that eliminates excessive copies of data and significantly decreases storage capacity requirements. 
    deduplication *bool
    // Encryption for the volumes.
    encryption V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable
    // Quality of service.
    qos V1beta1DatastoresPostRequestBody_volumeInfo_qosable
}
// NewV1beta1DatastoresPostRequestBody_volumeInfo instantiates a new V1beta1DatastoresPostRequestBody_volumeInfo and sets the default values.
func NewV1beta1DatastoresPostRequestBody_volumeInfo()(*V1beta1DatastoresPostRequestBody_volumeInfo) {
    m := &V1beta1DatastoresPostRequestBody_volumeInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresPostRequestBody_volumeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresPostRequestBody_volumeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresPostRequestBody_volumeInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllFlash gets the allFlash property value. Indicates whether all flash is enabled or not.
// returns a *bool when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetAllFlash()(*bool) {
    return m.allFlash
}
// GetDataReduction gets the dataReduction property value. Specifies if data reduction capability is enabled. Data reduction is a combination of both deduplication and compression.
// returns a *bool when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetDataReduction()(*bool) {
    return m.dataReduction
}
// GetDeduplication gets the deduplication property value. Specifies if dedupe is enabled for volumes created. Deduplication is a process that eliminates excessive copies of data and significantly decreases storage capacity requirements. 
// returns a *bool when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetDeduplication()(*bool) {
    return m.deduplication
}
// GetEncryption gets the encryption property value. Encryption for the volumes.
// returns a V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetEncryption()(V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable) {
    return m.encryption
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allFlash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllFlash(val)
        }
        return nil
    }
    res["dataReduction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataReduction(val)
        }
        return nil
    }
    res["deduplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeduplication(val)
        }
        return nil
    }
    res["encryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresPostRequestBody_volumeInfo_encryptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryption(val.(V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable))
        }
        return nil
    }
    res["qos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresPostRequestBody_volumeInfo_qosFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQos(val.(V1beta1DatastoresPostRequestBody_volumeInfo_qosable))
        }
        return nil
    }
    return res
}
// GetQos gets the qos property value. Quality of service.
// returns a V1beta1DatastoresPostRequestBody_volumeInfo_qosable when successful
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) GetQos()(V1beta1DatastoresPostRequestBody_volumeInfo_qosable) {
    return m.qos
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allFlash", m.GetAllFlash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dataReduction", m.GetDataReduction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deduplication", m.GetDeduplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("encryption", m.GetEncryption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("qos", m.GetQos())
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
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllFlash sets the allFlash property value. Indicates whether all flash is enabled or not.
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) SetAllFlash(value *bool)() {
    m.allFlash = value
}
// SetDataReduction sets the dataReduction property value. Specifies if data reduction capability is enabled. Data reduction is a combination of both deduplication and compression.
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) SetDataReduction(value *bool)() {
    m.dataReduction = value
}
// SetDeduplication sets the deduplication property value. Specifies if dedupe is enabled for volumes created. Deduplication is a process that eliminates excessive copies of data and significantly decreases storage capacity requirements. 
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) SetDeduplication(value *bool)() {
    m.deduplication = value
}
// SetEncryption sets the encryption property value. Encryption for the volumes.
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) SetEncryption(value V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable)() {
    m.encryption = value
}
// SetQos sets the qos property value. Quality of service.
func (m *V1beta1DatastoresPostRequestBody_volumeInfo) SetQos(value V1beta1DatastoresPostRequestBody_volumeInfo_qosable)() {
    m.qos = value
}
type V1beta1DatastoresPostRequestBody_volumeInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllFlash()(*bool)
    GetDataReduction()(*bool)
    GetDeduplication()(*bool)
    GetEncryption()(V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable)
    GetQos()(V1beta1DatastoresPostRequestBody_volumeInfo_qosable)
    SetAllFlash(value *bool)()
    SetDataReduction(value *bool)()
    SetDeduplication(value *bool)()
    SetEncryption(value V1beta1DatastoresPostRequestBody_volumeInfo_encryptionable)()
    SetQos(value V1beta1DatastoresPostRequestBody_volumeInfo_qosable)()
}
