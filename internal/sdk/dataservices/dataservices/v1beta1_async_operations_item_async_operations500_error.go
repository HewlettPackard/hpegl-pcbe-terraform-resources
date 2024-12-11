package dataservices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1AsyncOperationsItemAsyncOperations500Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Possible values: INTERNAL_ERROR
    errorCode *string
}
// NewV1beta1AsyncOperationsItemAsyncOperations500Error instantiates a new V1beta1AsyncOperationsItemAsyncOperations500Error and sets the default values.
func NewV1beta1AsyncOperationsItemAsyncOperations500Error()(*V1beta1AsyncOperationsItemAsyncOperations500Error) {
    m := &V1beta1AsyncOperationsItemAsyncOperations500Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1AsyncOperationsItemAsyncOperations500ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AsyncOperationsItemAsyncOperations500ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AsyncOperationsItemAsyncOperations500Error(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrorCode gets the errorCode property value. Possible values: INTERNAL_ERROR
// returns a *string when successful
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) GetErrorCode()(*string) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// Serialize serializes information the current object
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("errorCode", m.GetErrorCode())
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
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrorCode sets the errorCode property value. Possible values: INTERNAL_ERROR
func (m *V1beta1AsyncOperationsItemAsyncOperations500Error) SetErrorCode(value *string)() {
    m.errorCode = value
}
type V1beta1AsyncOperationsItemAsyncOperations500Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorCode()(*string)
    SetErrorCode(value *string)()
}