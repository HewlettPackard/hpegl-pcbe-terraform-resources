package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemStorageArraysGetResponse_items_details details regarding the Storage Array
type V1beta1SystemsItemStorageArraysGetResponse_items_details struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Extended Model of the Array
    extendedModel *string
    // Model of the Array
    model *string
    // Model subtype of the Array
    modelSubType *string
    // Serial of the Array
    serial *string
    // Version of the Array
    version *string
}
// NewV1beta1SystemsItemStorageArraysGetResponse_items_details instantiates a new V1beta1SystemsItemStorageArraysGetResponse_items_details and sets the default values.
func NewV1beta1SystemsItemStorageArraysGetResponse_items_details()(*V1beta1SystemsItemStorageArraysGetResponse_items_details) {
    m := &V1beta1SystemsItemStorageArraysGetResponse_items_details{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemStorageArraysGetResponse_items_detailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemStorageArraysGetResponse_items_detailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemStorageArraysGetResponse_items_details(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExtendedModel gets the extendedModel property value. Extended Model of the Array
// returns a *string when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetExtendedModel()(*string) {
    return m.extendedModel
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["extendedModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtendedModel(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["modelSubType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelSubType(val)
        }
        return nil
    }
    res["serial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerial(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetModel gets the model property value. Model of the Array
// returns a *string when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetModel()(*string) {
    return m.model
}
// GetModelSubType gets the modelSubType property value. Model subtype of the Array
// returns a *string when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetModelSubType()(*string) {
    return m.modelSubType
}
// GetSerial gets the serial property value. Serial of the Array
// returns a *string when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetSerial()(*string) {
    return m.serial
}
// GetVersion gets the version property value. Version of the Array
// returns a *string when successful
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("extendedModel", m.GetExtendedModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("modelSubType", m.GetModelSubType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serial", m.GetSerial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExtendedModel sets the extendedModel property value. Extended Model of the Array
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) SetExtendedModel(value *string)() {
    m.extendedModel = value
}
// SetModel sets the model property value. Model of the Array
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) SetModel(value *string)() {
    m.model = value
}
// SetModelSubType sets the modelSubType property value. Model subtype of the Array
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) SetModelSubType(value *string)() {
    m.modelSubType = value
}
// SetSerial sets the serial property value. Serial of the Array
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) SetSerial(value *string)() {
    m.serial = value
}
// SetVersion sets the version property value. Version of the Array
func (m *V1beta1SystemsItemStorageArraysGetResponse_items_details) SetVersion(value *string)() {
    m.version = value
}
type V1beta1SystemsItemStorageArraysGetResponse_items_detailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExtendedModel()(*string)
    GetModel()(*string)
    GetModelSubType()(*string)
    GetSerial()(*string)
    GetVersion()(*string)
    SetExtendedModel(value *string)()
    SetModel(value *string)()
    SetModelSubType(value *string)()
    SetSerial(value *string)()
    SetVersion(value *string)()
}
