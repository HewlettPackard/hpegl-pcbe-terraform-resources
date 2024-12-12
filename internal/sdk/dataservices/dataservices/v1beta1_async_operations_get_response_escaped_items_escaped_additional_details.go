package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AsyncOperationsGetResponse_items_additionalDetails references to other UI link include the consoleUri
type V1beta1AsyncOperationsGetResponse_items_additionalDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The consoleUri property
    consoleUri *string
}
// NewV1beta1AsyncOperationsGetResponse_items_additionalDetails instantiates a new V1beta1AsyncOperationsGetResponse_items_additionalDetails and sets the default values.
func NewV1beta1AsyncOperationsGetResponse_items_additionalDetails()(*V1beta1AsyncOperationsGetResponse_items_additionalDetails) {
    m := &V1beta1AsyncOperationsGetResponse_items_additionalDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1AsyncOperationsGetResponse_items_additionalDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AsyncOperationsGetResponse_items_additionalDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AsyncOperationsGetResponse_items_additionalDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1AsyncOperationsGetResponse_items_additionalDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsoleUri gets the consoleUri property value. The consoleUri property
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items_additionalDetails) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AsyncOperationsGetResponse_items_additionalDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["consoleUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsoleUri(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1AsyncOperationsGetResponse_items_additionalDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("consoleUri", m.GetConsoleUri())
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
func (m *V1beta1AsyncOperationsGetResponse_items_additionalDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsoleUri sets the consoleUri property value. The consoleUri property
func (m *V1beta1AsyncOperationsGetResponse_items_additionalDetails) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
type V1beta1AsyncOperationsGetResponse_items_additionalDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsoleUri()(*string)
    SetConsoleUri(value *string)()
}
