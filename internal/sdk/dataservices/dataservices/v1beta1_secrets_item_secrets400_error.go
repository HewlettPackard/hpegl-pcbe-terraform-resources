package dataservices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SecretsItemSecrets400Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Debug ID
    debugId *string
    // Error code
    errorCode *string
    // HTTP status code
    httpStatusCode *int32
    // Error message
    message *string
}
// NewV1beta1SecretsItemSecrets400Error instantiates a new V1beta1SecretsItemSecrets400Error and sets the default values.
func NewV1beta1SecretsItemSecrets400Error()(*V1beta1SecretsItemSecrets400Error) {
    m := &V1beta1SecretsItemSecrets400Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SecretsItemSecrets400ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SecretsItemSecrets400Error(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *V1beta1SecretsItemSecrets400Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SecretsItemSecrets400Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDebugId gets the debugId property value. Debug ID
// returns a *string when successful
func (m *V1beta1SecretsItemSecrets400Error) GetDebugId()(*string) {
    return m.debugId
}
// GetErrorCode gets the errorCode property value. Error code
// returns a *string when successful
func (m *V1beta1SecretsItemSecrets400Error) GetErrorCode()(*string) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SecretsItemSecrets400Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["debugId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebugId(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["httpStatusCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpStatusCode(val)
        }
        return nil
    }
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
// GetHttpStatusCode gets the httpStatusCode property value. HTTP status code
// returns a *int32 when successful
func (m *V1beta1SecretsItemSecrets400Error) GetHttpStatusCode()(*int32) {
    return m.httpStatusCode
}
// GetMessage gets the message property value. Error message
// returns a *string when successful
func (m *V1beta1SecretsItemSecrets400Error) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *V1beta1SecretsItemSecrets400Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("debugId", m.GetDebugId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("httpStatusCode", m.GetHttpStatusCode())
        if err != nil {
            return err
        }
    }
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
func (m *V1beta1SecretsItemSecrets400Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDebugId sets the debugId property value. Debug ID
func (m *V1beta1SecretsItemSecrets400Error) SetDebugId(value *string)() {
    m.debugId = value
}
// SetErrorCode sets the errorCode property value. Error code
func (m *V1beta1SecretsItemSecrets400Error) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetHttpStatusCode sets the httpStatusCode property value. HTTP status code
func (m *V1beta1SecretsItemSecrets400Error) SetHttpStatusCode(value *int32)() {
    m.httpStatusCode = value
}
// SetMessage sets the message property value. Error message
func (m *V1beta1SecretsItemSecrets400Error) SetMessage(value *string)() {
    m.message = value
}
type V1beta1SecretsItemSecrets400Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDebugId()(*string)
    GetErrorCode()(*string)
    GetHttpStatusCode()(*int32)
    GetMessage()(*string)
    SetDebugId(value *string)()
    SetErrorCode(value *string)()
    SetHttpStatusCode(value *int32)()
    SetMessage(value *string)()
}
