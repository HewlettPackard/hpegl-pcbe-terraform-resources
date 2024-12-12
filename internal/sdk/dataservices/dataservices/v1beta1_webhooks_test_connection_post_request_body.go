package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksTestConnectionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An object containing the publishing destination attributes
    destination V1beta1WebhooksTestConnectionPostRequestBody_destinationable
}
// NewV1beta1WebhooksTestConnectionPostRequestBody instantiates a new V1beta1WebhooksTestConnectionPostRequestBody and sets the default values.
func NewV1beta1WebhooksTestConnectionPostRequestBody()(*V1beta1WebhooksTestConnectionPostRequestBody) {
    m := &V1beta1WebhooksTestConnectionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksTestConnectionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksTestConnectionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksTestConnectionPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksTestConnectionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestination gets the destination property value. An object containing the publishing destination attributes
// returns a V1beta1WebhooksTestConnectionPostRequestBody_destinationable when successful
func (m *V1beta1WebhooksTestConnectionPostRequestBody) GetDestination()(V1beta1WebhooksTestConnectionPostRequestBody_destinationable) {
    return m.destination
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksTestConnectionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksTestConnectionPostRequestBody_destinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestination(val.(V1beta1WebhooksTestConnectionPostRequestBody_destinationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksTestConnectionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("destination", m.GetDestination())
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
func (m *V1beta1WebhooksTestConnectionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestination sets the destination property value. An object containing the publishing destination attributes
func (m *V1beta1WebhooksTestConnectionPostRequestBody) SetDestination(value V1beta1WebhooksTestConnectionPostRequestBody_destinationable)() {
    m.destination = value
}
type V1beta1WebhooksTestConnectionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestination()(V1beta1WebhooksTestConnectionPostRequestBody_destinationable)
    SetDestination(value V1beta1WebhooksTestConnectionPostRequestBody_destinationable)()
}
