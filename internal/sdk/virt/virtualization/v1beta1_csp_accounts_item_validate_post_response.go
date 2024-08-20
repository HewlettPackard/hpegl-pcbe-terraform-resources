package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1CspAccountsItemValidatePostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An action that must be performed manually to enable the requested operation
    requiredAction *string
}
// NewV1beta1CspAccountsItemValidatePostResponse instantiates a new V1beta1CspAccountsItemValidatePostResponse and sets the default values.
func NewV1beta1CspAccountsItemValidatePostResponse()(*V1beta1CspAccountsItemValidatePostResponse) {
    m := &V1beta1CspAccountsItemValidatePostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspAccountsItemValidatePostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspAccountsItemValidatePostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspAccountsItemValidatePostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspAccountsItemValidatePostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspAccountsItemValidatePostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["requiredAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAction(val)
        }
        return nil
    }
    return res
}
// GetRequiredAction gets the requiredAction property value. An action that must be performed manually to enable the requested operation
// returns a *string when successful
func (m *V1beta1CspAccountsItemValidatePostResponse) GetRequiredAction()(*string) {
    return m.requiredAction
}
// Serialize serializes information the current object
func (m *V1beta1CspAccountsItemValidatePostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("requiredAction", m.GetRequiredAction())
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
func (m *V1beta1CspAccountsItemValidatePostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRequiredAction sets the requiredAction property value. An action that must be performed manually to enable the requested operation
func (m *V1beta1CspAccountsItemValidatePostResponse) SetRequiredAction(value *string)() {
    m.requiredAction = value
}
type V1beta1CspAccountsItemValidatePostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequiredAction()(*string)
    SetRequiredAction(value *string)()
}
