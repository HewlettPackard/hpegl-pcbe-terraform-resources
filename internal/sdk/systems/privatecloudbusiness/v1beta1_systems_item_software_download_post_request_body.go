package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemSoftwareDownloadPostRequestBody request parameters for the software download operation.
type V1beta1SystemsItemSoftwareDownloadPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Version of the target software catalog for the software download request.
    catalogVersion *string
}
// NewV1beta1SystemsItemSoftwareDownloadPostRequestBody instantiates a new V1beta1SystemsItemSoftwareDownloadPostRequestBody and sets the default values.
func NewV1beta1SystemsItemSoftwareDownloadPostRequestBody()(*V1beta1SystemsItemSoftwareDownloadPostRequestBody) {
    m := &V1beta1SystemsItemSoftwareDownloadPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemSoftwareDownloadPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemSoftwareDownloadPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemSoftwareDownloadPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemSoftwareDownloadPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCatalogVersion gets the catalogVersion property value. Version of the target software catalog for the software download request.
// returns a *string when successful
func (m *V1beta1SystemsItemSoftwareDownloadPostRequestBody) GetCatalogVersion()(*string) {
    return m.catalogVersion
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemSoftwareDownloadPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogVersion(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemSoftwareDownloadPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("catalogVersion", m.GetCatalogVersion())
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
func (m *V1beta1SystemsItemSoftwareDownloadPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCatalogVersion sets the catalogVersion property value. Version of the target software catalog for the software download request.
func (m *V1beta1SystemsItemSoftwareDownloadPostRequestBody) SetCatalogVersion(value *string)() {
    m.catalogVersion = value
}
type V1beta1SystemsItemSoftwareDownloadPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalogVersion()(*string)
    SetCatalogVersion(value *string)()
}
