package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1WebhooksTestConnectionPostRequestBody_destination an object containing the publishing destination attributes
type V1beta1WebhooksTestConnectionPostRequestBody_destination struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cloudEventEndpoint property
    cloudEventEndpoint V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable
}
// NewV1beta1WebhooksTestConnectionPostRequestBody_destination instantiates a new V1beta1WebhooksTestConnectionPostRequestBody_destination and sets the default values.
func NewV1beta1WebhooksTestConnectionPostRequestBody_destination()(*V1beta1WebhooksTestConnectionPostRequestBody_destination) {
    m := &V1beta1WebhooksTestConnectionPostRequestBody_destination{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksTestConnectionPostRequestBody_destinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksTestConnectionPostRequestBody_destinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksTestConnectionPostRequestBody_destination(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksTestConnectionPostRequestBody_destination) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudEventEndpoint gets the cloudEventEndpoint property value. The cloudEventEndpoint property
// returns a V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable when successful
func (m *V1beta1WebhooksTestConnectionPostRequestBody_destination) GetCloudEventEndpoint()(V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable) {
    return m.cloudEventEndpoint
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksTestConnectionPostRequestBody_destination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudEventEndpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudEventEndpoint(val.(V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksTestConnectionPostRequestBody_destination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cloudEventEndpoint", m.GetCloudEventEndpoint())
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
func (m *V1beta1WebhooksTestConnectionPostRequestBody_destination) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudEventEndpoint sets the cloudEventEndpoint property value. The cloudEventEndpoint property
func (m *V1beta1WebhooksTestConnectionPostRequestBody_destination) SetCloudEventEndpoint(value V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable)() {
    m.cloudEventEndpoint = value
}
type V1beta1WebhooksTestConnectionPostRequestBody_destinationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudEventEndpoint()(V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable)
    SetCloudEventEndpoint(value V1beta1WebhooksTestConnectionPostRequestBody_destination_cloudEventEndpointable)()
}
