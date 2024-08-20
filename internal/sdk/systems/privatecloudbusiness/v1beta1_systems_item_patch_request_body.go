package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemPatchRequestBody request parameter for patch operation on the system.
type V1beta1SystemsItemPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Location details to be updated on the system
    location V1beta1SystemsItemPatchRequestBody_locationable
}
// NewV1beta1SystemsItemPatchRequestBody instantiates a new V1beta1SystemsItemPatchRequestBody and sets the default values.
func NewV1beta1SystemsItemPatchRequestBody()(*V1beta1SystemsItemPatchRequestBody) {
    m := &V1beta1SystemsItemPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemPatchRequestBody_locationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(V1beta1SystemsItemPatchRequestBody_locationable))
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. Location details to be updated on the system
// returns a V1beta1SystemsItemPatchRequestBody_locationable when successful
func (m *V1beta1SystemsItemPatchRequestBody) GetLocation()(V1beta1SystemsItemPatchRequestBody_locationable) {
    return m.location
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("location", m.GetLocation())
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
func (m *V1beta1SystemsItemPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocation sets the location property value. Location details to be updated on the system
func (m *V1beta1SystemsItemPatchRequestBody) SetLocation(value V1beta1SystemsItemPatchRequestBody_locationable)() {
    m.location = value
}
type V1beta1SystemsItemPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(V1beta1SystemsItemPatchRequestBody_locationable)
    SetLocation(value V1beta1SystemsItemPatchRequestBody_locationable)()
}
