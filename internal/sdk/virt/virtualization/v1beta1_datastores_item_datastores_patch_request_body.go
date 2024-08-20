package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DatastoresItemDatastoresPatchRequestBody request body for Edit Datastore API
type V1beta1DatastoresItemDatastoresPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies if data reduction capability is enabled. Data reduction is a combination of both deduplication and compression.
    dataReduction *bool
    // Specifies if dedupe is enabled for volumes created. Deduplication is a process that eliminates excessive copies of data and significantly decreases storage capacity requirements. 
    deduplication *bool
    // Quality of service.
    qos V1beta1DatastoresItemDatastoresPatchRequestBody_qosable
    // Size of the datastore in bytes
    sizeInBytes *int64
}
// NewV1beta1DatastoresItemDatastoresPatchRequestBody instantiates a new V1beta1DatastoresItemDatastoresPatchRequestBody and sets the default values.
func NewV1beta1DatastoresItemDatastoresPatchRequestBody()(*V1beta1DatastoresItemDatastoresPatchRequestBody) {
    m := &V1beta1DatastoresItemDatastoresPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DatastoresItemDatastoresPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DatastoresItemDatastoresPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DatastoresItemDatastoresPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDataReduction gets the dataReduction property value. Specifies if data reduction capability is enabled. Data reduction is a combination of both deduplication and compression.
// returns a *bool when successful
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) GetDataReduction()(*bool) {
    return m.dataReduction
}
// GetDeduplication gets the deduplication property value. Specifies if dedupe is enabled for volumes created. Deduplication is a process that eliminates excessive copies of data and significantly decreases storage capacity requirements. 
// returns a *bool when successful
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) GetDeduplication()(*bool) {
    return m.deduplication
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["qos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DatastoresItemDatastoresPatchRequestBody_qosFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQos(val.(V1beta1DatastoresItemDatastoresPatchRequestBody_qosable))
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
    return res
}
// GetQos gets the qos property value. Quality of service.
// returns a V1beta1DatastoresItemDatastoresPatchRequestBody_qosable when successful
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) GetQos()(V1beta1DatastoresItemDatastoresPatchRequestBody_qosable) {
    return m.qos
}
// GetSizeInBytes gets the sizeInBytes property value. Size of the datastore in bytes
// returns a *int64 when successful
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) GetSizeInBytes()(*int64) {
    return m.sizeInBytes
}
// Serialize serializes information the current object
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("qos", m.GetQos())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDataReduction sets the dataReduction property value. Specifies if data reduction capability is enabled. Data reduction is a combination of both deduplication and compression.
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) SetDataReduction(value *bool)() {
    m.dataReduction = value
}
// SetDeduplication sets the deduplication property value. Specifies if dedupe is enabled for volumes created. Deduplication is a process that eliminates excessive copies of data and significantly decreases storage capacity requirements. 
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) SetDeduplication(value *bool)() {
    m.deduplication = value
}
// SetQos sets the qos property value. Quality of service.
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) SetQos(value V1beta1DatastoresItemDatastoresPatchRequestBody_qosable)() {
    m.qos = value
}
// SetSizeInBytes sets the sizeInBytes property value. Size of the datastore in bytes
func (m *V1beta1DatastoresItemDatastoresPatchRequestBody) SetSizeInBytes(value *int64)() {
    m.sizeInBytes = value
}
type V1beta1DatastoresItemDatastoresPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataReduction()(*bool)
    GetDeduplication()(*bool)
    GetQos()(V1beta1DatastoresItemDatastoresPatchRequestBody_qosable)
    GetSizeInBytes()(*int64)
    SetDataReduction(value *bool)()
    SetDeduplication(value *bool)()
    SetQos(value V1beta1DatastoresItemDatastoresPatchRequestBody_qosable)()
    SetSizeInBytes(value *int64)()
}
