package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The message property
    message *string
}
// NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations instantiates a new V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations and sets the default values.
func NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations()(*V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) {
    m := &V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. The message property
func (m *V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendations) SetMessage(value *string)() {
    m.message = value
}
type V1beta1AsyncOperationsItemAsyncOperationsGetResponse_recommendationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    SetMessage(value *string)()
}
