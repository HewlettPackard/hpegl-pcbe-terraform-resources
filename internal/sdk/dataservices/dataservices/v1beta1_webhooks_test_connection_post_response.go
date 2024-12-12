package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1WebhooksTestConnectionPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error property
    error V1beta1WebhooksTestConnectionPostResponse_errorable
    // Indicates whether or not connection to the remote endpoint succeeded.
    succeeded *bool
}
// NewV1beta1WebhooksTestConnectionPostResponse instantiates a new V1beta1WebhooksTestConnectionPostResponse and sets the default values.
func NewV1beta1WebhooksTestConnectionPostResponse()(*V1beta1WebhooksTestConnectionPostResponse) {
    m := &V1beta1WebhooksTestConnectionPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1WebhooksTestConnectionPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1WebhooksTestConnectionPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1WebhooksTestConnectionPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1WebhooksTestConnectionPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. The error property
// returns a V1beta1WebhooksTestConnectionPostResponse_errorable when successful
func (m *V1beta1WebhooksTestConnectionPostResponse) GetError()(V1beta1WebhooksTestConnectionPostResponse_errorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1WebhooksTestConnectionPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1WebhooksTestConnectionPostResponse_errorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(V1beta1WebhooksTestConnectionPostResponse_errorable))
        }
        return nil
    }
    res["succeeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSucceeded(val)
        }
        return nil
    }
    return res
}
// GetSucceeded gets the succeeded property value. Indicates whether or not connection to the remote endpoint succeeded.
// returns a *bool when successful
func (m *V1beta1WebhooksTestConnectionPostResponse) GetSucceeded()(*bool) {
    return m.succeeded
}
// Serialize serializes information the current object
func (m *V1beta1WebhooksTestConnectionPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("succeeded", m.GetSucceeded())
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
func (m *V1beta1WebhooksTestConnectionPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. The error property
func (m *V1beta1WebhooksTestConnectionPostResponse) SetError(value V1beta1WebhooksTestConnectionPostResponse_errorable)() {
    m.error = value
}
// SetSucceeded sets the succeeded property value. Indicates whether or not connection to the remote endpoint succeeded.
func (m *V1beta1WebhooksTestConnectionPostResponse) SetSucceeded(value *bool)() {
    m.succeeded = value
}
type V1beta1WebhooksTestConnectionPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(V1beta1WebhooksTestConnectionPostResponse_errorable)
    GetSucceeded()(*bool)
    SetError(value V1beta1WebhooksTestConnectionPostResponse_errorable)()
    SetSucceeded(value *bool)()
}
