package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager details of the HPE Storage Connection Manager software.
type V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewV1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager instantiates a new V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager and sets the default values.
func NewV1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager()(*V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager) {
    m := &V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManagerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManagerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManager) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type V1beta1SystemSoftwareCatalogsGetResponse_items_storageConnectionManagerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
